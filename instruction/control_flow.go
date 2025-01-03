package instruction

import (
	"fmt"
	"io"

	"github.com/Warashi/wasmium/leb128"
	"github.com/Warashi/wasmium/opcode"
	"github.com/Warashi/wasmium/types/binary"
	"github.com/Warashi/wasmium/types/runtime"
)

func decodeBlock(r io.Reader) (binary.Block, error) {
	var buf [1]byte

	if _, err := io.ReadFull(r, buf[:]); err != nil {
		return binary.Block{}, fmt.Errorf("failed to read block type: %w", err)
	}

	switch buf[0] {
	case 0x40:
		return binary.Block{BlockType: binary.BlockTypeVoid{}}, nil
	default:
		return binary.Block{BlockType: binary.BlockTypeValue{ValueTypes: []binary.ValueType{binary.ValueType(buf[0])}}}, nil
	}
}

func br(r runtime.Runtime, f *runtime.Frame, level uint32) (int, error) {
	index := f.Labels.Len() - 1 - int(level)
	label := f.Labels[index]

	if label.Kind() == runtime.LabelKindLoop {
		// NOTE: we still need loop label to jump to the beginning of the loop.
		f.Labels.Drain(index + 1)

		// NOTE: since it jumps to the beginning of the loop,
		// the stack is unwound without considering the return value.
		if err := r.StackUnwind(label.StackPointer(), 0); err != nil {
			return 0, fmt.Errorf("failed to unwind stack: %w", err)
		}

		return label.Start(), nil
	}
	f.Labels.Drain(index)
	if err := r.StackUnwind(label.StackPointer(), label.Arity()); err != nil {
		return 0, fmt.Errorf("failed to unwind stack: %w", err)
	}
	return label.ProgramCounter(), nil
}

func getEndAddress(insts []runtime.Instruction, programCounter int) (int, error) {
	depth := 0
	for {
		programCounter++
		if programCounter < 0 || len(insts) <= programCounter {
			return 0, fmt.Errorf("unexpected end of instructions")
		}

		switch insts[programCounter].(type) {
		case *If:
			depth++
		case *Block:
			depth++
		case *Loop:
			depth++
		case *End:
			if depth == 0 {
				return programCounter, nil
			}
			depth--
		default:
			// do nothing
		}
	}
}

type Unreachable struct{}

func (*Unreachable) Opcode() opcode.Opcode { return opcode.OpcodeUnreachable }

func (*Unreachable) ReadOperandsFrom(io.Reader) error { return nil }

func (*Unreachable) Execute(runtime.Runtime, *runtime.Frame) error {
	return fmt.Errorf("unreachable")
}

type Nop struct{}

func (*Nop) Opcode() opcode.Opcode { return opcode.OpcodeNop }

func (*Nop) ReadOperandsFrom(io.Reader) error { return nil }

func (*Nop) Execute(runtime.Runtime, *runtime.Frame) error {
	return nil
}

type Block struct {
	Block binary.Block
}

func (*Block) Opcode() opcode.Opcode { return opcode.OpcodeBlock }

func (b *Block) ReadOperandsFrom(r io.Reader) error {
	var err error
	b.Block, err = decodeBlock(r)
	return err
}

func (b *Block) Execute(r runtime.Runtime, f *runtime.Frame) error {
	arity := b.Block.BlockType.ResultCount()
	pc, err := getEndAddress(f.Instructions, f.ProgramCounter)
	if err != nil {
		return fmt.Errorf("failed to get end address: %w", err)
	}
	f.Labels.Push(runtime.NewLabel(runtime.LabelKindBlock, 0, pc, r.StackLen(), arity))
	return nil
}

type Loop struct {
	Block binary.Block
}

func (*Loop) Opcode() opcode.Opcode { return opcode.OpcodeLoop }
func (l *Loop) ReadOperandsFrom(r io.Reader) error {
	var err error
	l.Block, err = decodeBlock(r)
	return err
}

func (l *Loop) Execute(r runtime.Runtime, f *runtime.Frame) error {
	arity := l.Block.BlockType.ResultCount()
	startProgramCounter := f.ProgramCounter
	programCounter, err := getEndAddress(f.Instructions, f.ProgramCounter)
	if err != nil {
		return fmt.Errorf("failed to get end address: %w", err)
	}

	f.Labels.Push(runtime.NewLabel(runtime.LabelKindLoop, startProgramCounter, programCounter, r.StackLen(), arity))

	return nil
}

type If struct {
	Block binary.Block
}

func (*If) Opcode() opcode.Opcode { return opcode.OpcodeIf }
func (i *If) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Block, err = decodeBlock(r)
	return err
}
func (i *If) Execute(r runtime.Runtime, f *runtime.Frame) error {
	cond, err := r.PopStack()
	if err != nil {
		return fmt.Errorf("failed to pop stack: %w", err)
	}

	nextProgramCounter, err := getEndAddress(f.Instructions, f.ProgramCounter)
	if err != nil {
		return fmt.Errorf("failed to get end address: %w", err)
	}

	if !cond.Bool() {
		f.ProgramCounter = nextProgramCounter
	}

	f.Labels.Push(runtime.NewLabel(runtime.LabelKindIf, 0, nextProgramCounter, r.StackLen(), i.Block.BlockType.ResultCount()))

	return nil
}

type End struct{}

func (*End) Opcode() opcode.Opcode { return opcode.OpcodeEnd }

func (*End) ReadOperandsFrom(io.Reader) error { return nil }

func (*End) Execute(r runtime.Runtime, f *runtime.Frame) error {
	if f.Labels.Len() > 0 {
		label := f.Labels.Pop()
		f.ProgramCounter = label.ProgramCounter()
		if err := r.StackUnwind(label.StackPointer(), label.Arity()); err != nil {
			return fmt.Errorf("failed to unwind stack: %w", err)
		}
	} else {
		// If the label stack is empty, it means the end of the function.
		frame, err := r.PopCallStack()
		if err != nil {
			return fmt.Errorf("failed to pop call stack: %w", err)
		}
		if err := r.StackUnwind(frame.StackPointer, frame.Arity); err != nil {
			return fmt.Errorf("failed to unwind stack: %w", err)
		}
	}
	return nil
}

type Br struct {
	Level uint32
}

func (*Br) Opcode() opcode.Opcode { return opcode.OpcodeBr }

func (b *Br) ReadOperandsFrom(r io.Reader) error {
	var err error
	b.Level, err = leb128.Uint32(r)
	return err
}

func (b *Br) Execute(r runtime.Runtime, f *runtime.Frame) error {
	var err error
	f.ProgramCounter, err = br(r, f, b.Level)
	return err
}

type BrIf struct {
	Level uint32
}

func (*BrIf) Opcode() opcode.Opcode { return opcode.OpcodeBrIf }

func (b *BrIf) ReadOperandsFrom(r io.Reader) error {
	var err error
	b.Level, err = leb128.Uint32(r)
	return err
}

func (b *BrIf) Execute(r runtime.Runtime, f *runtime.Frame) error {
	cond, err := r.PopStack()
	if err != nil {
		return fmt.Errorf("failed to pop stack: %w", err)
	}

	if !cond.Bool() {
		return nil
	}

	f.ProgramCounter, err = br(r, f, b.Level)
	return err
}

type BrTable struct {
	Levels  []uint32
	Default uint32
}

func (*BrTable) Opcode() opcode.Opcode { return opcode.OpcodeBrTable }

func (b *BrTable) ReadOperandsFrom(r io.Reader) error {
	var err error
	count, err := leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read count: %w", err)
	}

	b.Levels = make([]uint32, 0, count)
	for range count {
		level, err := leb128.Uint32(r)
		if err != nil {
			return fmt.Errorf("failed to read level: %w", err)
		}
		b.Levels = append(b.Levels, level)
	}

	b.Default, err = leb128.Uint32(r)
	return err
}

func (b *BrTable) Execute(r runtime.Runtime, f *runtime.Frame) error {
	cond, err := r.PopStack()
	if err != nil {
		return fmt.Errorf("failed to pop stack: %w", err)
	}

	var level uint32
	index := cond.Int()
	if index < 0 || len(b.Levels) <= index {
		level = b.Default
	} else {
		level = b.Levels[index]
	}

	f.ProgramCounter, err = br(r, f, level)
	return err
}

type Return struct{}

func (*Return) Opcode() opcode.Opcode { return opcode.OpcodeReturn }

func (*Return) ReadOperandsFrom(io.Reader) error { return nil }

func (*Return) Execute(r runtime.Runtime, f *runtime.Frame) error {
	frame, err := r.PopCallStack()
	if err != nil {
		return fmt.Errorf("failed to pop call stack: %w", err)
	}
	if err := r.StackUnwind(frame.StackPointer, frame.Arity); err != nil {
		return fmt.Errorf("failed to unwind stack: %w", err)
	}

	return nil
}

type Call struct {
	Index uint32
}

func (c *Call) Opcode() opcode.Opcode { return opcode.OpcodeCall }

func (c *Call) ReadOperandsFrom(r io.Reader) error {
	var err error
	c.Index, err = leb128.Uint32(r)
	return err
}

func (c *Call) Execute(r runtime.Runtime, f *runtime.Frame) error {
	funcInst, err := r.Func(int(c.Index))
	if err != nil {
		return fmt.Errorf("failed to get function: %w", err)
	}
	switch f := funcInst.(type) {
	case runtime.InternalFuncInst:
		return r.PushFrame(f)
	case runtime.ExternalFuncInst:
		v, err := r.InvokeExternal(f)
		if err != nil {
			return fmt.Errorf("failed to invoke external function: %w", err)
		}
		for _, v := range v {
			r.PushStack(v)
		}
	}
	return nil
}

type Drop struct{}

func (*Drop) Opcode() opcode.Opcode { return opcode.OpcodeDrop }

func (*Drop) ReadOperandsFrom(io.Reader) error { return nil }

func (*Drop) Execute(r runtime.Runtime, f *runtime.Frame) error {
	_, err := r.PopStack()
	return err
}

type Select struct{}

func (*Select) Opcode() opcode.Opcode { return opcode.OpcodeSelect }

func (*Select) ReadOperandsFrom(io.Reader) error { return nil }

func (*Select) Execute(r runtime.Runtime, f *runtime.Frame) error {
	cond, err := r.PopStack()
	if err != nil {
		return fmt.Errorf("failed to pop stack: %w", err)
	}

	v2, err := r.PopStack()
	if err != nil {
		return fmt.Errorf("failed to pop stack: %w", err)
	}

	v1, err := r.PopStack()
	if err != nil {
		return fmt.Errorf("failed to pop stack: %w", err)
	}

	if cond.Bool() {
		r.PushStack(v1)
	} else {
		r.PushStack(v2)
	}

	return nil
}
