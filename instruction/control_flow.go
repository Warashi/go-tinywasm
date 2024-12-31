package instruction

import (
	"fmt"
	"io"

	"github.com/Warashi/go-tinywasm/leb128"
	"github.com/Warashi/go-tinywasm/opcode"
	"github.com/Warashi/go-tinywasm/types/runtime"
)

type If struct {
	block Block
}

func (*If) Opcode() opcode.Opcode { return opcode.OpcodeIf }
func (i *If) ReadOperandsFrom(r io.Reader) error {
	return i.block.decode(r)
}
func (i *If) Block() Block { return i.block }
func (i *If) Execute(r runtime.Runtime, f *runtime.Frame) error {
	cond, err := r.PopStack()
	if err != nil {
		return fmt.Errorf("failed to pop stack: %w", err)
	}

	nextProgramCounter, err := i.getEndAddress(f.Instructions, f.ProgramCounter)
	if err != nil {
		return fmt.Errorf("failed to get end address: %w", err)
	}

	if runtime.Falsy(cond) {
		f.ProgramCounter = nextProgramCounter
	}

	f.Labels.Push(runtime.NewLabel(runtime.LabelKindIf, nextProgramCounter, r.StackLen(), i.Block().BlockType().ResultCount()))

	return nil
}

func (*If) getEndAddress(insts []runtime.Instruction, programCounter int) (int, error) {
	depth := 0
	for {
		programCounter++
		if programCounter < 0 || len(insts) <= programCounter {
			return 0, fmt.Errorf("unexpected end of instructions")
		}

		switch insts[programCounter].(type) {
		case *If:
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
	index uint32
}

func (c *Call) Opcode() opcode.Opcode { return opcode.OpcodeCall }

func (c *Call) ReadOperandsFrom(r io.Reader) error {
	var err error
	c.index, err = leb128.Uint32(r)
	return err
}

func (c *Call) Index() uint32 { return c.index }

func (c *Call) Execute(r runtime.Runtime, f *runtime.Frame) error {
	funcInst, err := r.Func(int(c.index))
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
