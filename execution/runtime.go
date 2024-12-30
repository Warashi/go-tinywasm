package execution

import (
	"fmt"
	"io"
	"slices"

	"github.com/Warashi/go-tinywasm/binary"
)

type Frame struct {
	programCounter int
	stackPointer   int
	instructions   []binary.Instruction
	arity          int
	labels         stack[Label]
	locals         []Value
}

type stack[T any] []T

func (s *stack[T]) push(v T) {
	*s = append(*s, v)
}

func (s *stack[T]) pop() T {
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}

func (s *stack[T]) drain(n int) {
	*s = (*s)[:n]
}

func (s *stack[T]) splitOff(n int) stack[T] {
	r := (*s)[n:]
	*s = (*s)[:n]
	return slices.Clone(r)
}

func (s *stack[T]) len() int {
	return len(*s)
}

type Runtime struct {
	store     *Store
	stack     stack[Value]
	callStack stack[*Frame]
	imports   Import
	wasi      *WasiSnapshotPreview1
}

func NewRuntime(r io.Reader) (*Runtime, error) {
	module, err := binary.NewModule(r)
	if err != nil {
		return nil, fmt.Errorf("failed to create module: %w", err)
	}

	store, err := NewStore(module)
	if err != nil {
		return nil, fmt.Errorf("failed to create store: %w", err)
	}

	return &Runtime{
		store: store,
	}, nil
}

func NewRuntimeWithWasi(r io.Reader) (*Runtime, error) {
	runtime, err := NewRuntime(r)
	if err != nil {
		return nil, fmt.Errorf("failed to create runtime: %w", err)
	}
	runtime.wasi = NewWasiSnapshotPreview1()

	return runtime, nil
}

func (r *Runtime) Call(name string, args []Value) ([]Value, error) {
	inst, ok := r.store.module.exports[name]
	if !ok {
		return nil, fmt.Errorf("export not found: %s", name)
	}

	var index int
	switch desc := inst.desc.(type) {
	case binary.ExportDescFunc:
		index = int(desc.Index())
	}

	if index < 0 || len(r.store.funcs) <= index {
		return nil, fmt.Errorf("invalid function index: %d", index)
	}

	f := r.store.funcs[index]

	for _, arg := range args {
		r.stack.push(arg)
	}

	switch f := f.(type) {
	case InternalFuncInst:
		return r.invokeInternal(f)
	case ExternalFuncInst:
		return r.invokeExternal(f)
	default:
		return nil, fmt.Errorf("unsupported function type: %T", f)
	}
}

func (r *Runtime) AddImport(moduleName, funcName string, fn ImportFunc) {
	if r.imports == nil {
		r.imports = make(Import)
	}
	if _, ok := r.imports[moduleName]; !ok {
		r.imports[moduleName] = make(map[string]ImportFunc)
	}
	r.imports[moduleName][funcName] = fn

	return
}

func (r *Runtime) execute() error {
	for len(r.callStack) > 0 {
		frame := r.callStack[len(r.callStack)-1]

		frame.programCounter++

		if len(frame.instructions) <= frame.programCounter {
			break
		}

		switch inst := frame.instructions[frame.programCounter].(type) {
		case *binary.InstructionIf:
			if r.stack.len() < 1 {
				return fmt.Errorf("stack underflow")
			}

			condition := r.stack.pop()

			nextProgramCounter, err := r.getEndAddress(frame.instructions, frame.programCounter)
			if err != nil {
				return fmt.Errorf("failed to get end address: %w", err)
			}

			if condition == ValueI32(0) {
				frame.programCounter = nextProgramCounter
			}

			frame.labels.push(Label{
				kind:           LabelKindIf,
				programCounter: nextProgramCounter,
				stackPointer:   r.stack.len(),
				arity:          inst.Block().BlockType().ResultCount(),
			})
		case *binary.InstructionEnd:
			if frame.labels.len() > 0 {
				label := frame.labels.pop()
				frame.programCounter = label.programCounter
				if err := r.stackUnwind(label.stackPointer, label.arity); err != nil {
					return fmt.Errorf("failed to unwind stack: %w", err)
				}
			} else {
				if r.callStack.len() < 1 {
					return fmt.Errorf("call stack underflow")
				}
				frame := r.callStack.pop()
				if err := r.stackUnwind(frame.stackPointer, frame.arity); err != nil {
					return fmt.Errorf("failed to unwind stack: %w", err)
				}
			}
		case *binary.InstructionReturn:
			if r.callStack.len() < 1 {
				return fmt.Errorf("call stack underflow")
			}
			frame := r.callStack.pop()
			if err := r.stackUnwind(frame.stackPointer, frame.arity); err != nil {
				return fmt.Errorf("failed to unwind stack: %w", err)
			}
		case *binary.InstructionCall:
			if int(inst.Index()) < 0 || len(r.store.funcs) <= int(inst.Index()) {
				return fmt.Errorf("invalid function index: %d", inst.Index())
			}
			switch f := r.store.funcs[inst.Index()].(type) {
			case InternalFuncInst:
				r.pushFrame(f)
			case ExternalFuncInst:
				v, err := r.invokeExternal(f)
				if err != nil {
					return fmt.Errorf("failed to invoke external function: %w", err)
				}
				for _, v := range v {
					r.stack.push(v)
				}
			}
		case *binary.InstructionLocalGet:
			if len(frame.locals) <= int(inst.Index()) {
				return fmt.Errorf("invalid local index: %d", inst.Index())
			}
			r.stack.push(frame.locals[inst.Index()])
		case *binary.InstructionLocalSet:
			if len(frame.locals) <= int(inst.Index()) {
				return fmt.Errorf("invalid local index: %d", inst.Index())
			}
			if r.stack.len() < 1 {
				return fmt.Errorf("stack underflow")
			}
			frame.locals[inst.Index()] = r.stack.pop()
		case *binary.InstructionI32Store:
			if r.stack.len() < 2 {
				return fmt.Errorf("stack underflow")
			}
			value, addr := r.stack.pop(), r.stack.pop()
			at := int(addr.(ValueI32)) + int(inst.Offset())
			end := at + 4
			if _, err := writeValue(r.store.memories[0].data[at:end], value); err != nil {
				return fmt.Errorf("failed to store: %w", err)
			}
		case *binary.InstructionI32Const:
			r.stack.push(ValueI32(inst.Value()))
		case *binary.InstructionI32Add:
			if r.stack.len() < 2 {
				return fmt.Errorf("stack underflow")
			}
			right, left := r.stack.pop(), r.stack.pop()

			result, err := Add(left, right)
			if err != nil {
				return fmt.Errorf("failed to add: %w", err)
			}
			r.stack.push(result)
		case *binary.InstructionI32Sub:
			if r.stack.len() < 2 {
				return fmt.Errorf("stack underflow")
			}
			right, left := r.stack.pop(), r.stack.pop()

			result, err := Sub(left, right)
			if err != nil {
				return fmt.Errorf("failed to sub: %w", err)
			}
			r.stack.push(result)
		case *binary.InstructionI32LtS:
			if r.stack.len() < 2 {
				return fmt.Errorf("stack underflow")
			}
			right, left := r.stack.pop(), r.stack.pop()

			result, err := LessThan(left, right)
			if err != nil {
				return fmt.Errorf("failed to compare: %w", err)
			}
			r.stack.push(result)
		default:
			return fmt.Errorf("unsupported instruction: %T", inst)
		}
	}

	return nil
}

func (r *Runtime) stackUnwind(stackPointer, arity int) error {
	if arity == 0 {
		if r.stack.len() < stackPointer {
			return fmt.Errorf("stack underflow")
		}
		r.stack.drain(stackPointer)
		return nil
	}
	if r.stack.len() < stackPointer+arity {
		return fmt.Errorf("stack underflow")
	}

	returns := make([]Value, 0, arity)
	for range arity {
		returns = append(returns, r.stack.pop())
	}

	r.stack.drain(stackPointer)

	for _, v := range returns {
		r.stack.push(v)
	}

	return nil
}
func (r *Runtime) pushFrame(f InternalFuncInst) {
	bottom := r.stack.len() - len(f.funcType.Params())
	locals := r.stack.splitOff(bottom)

	for _, local := range f.code.locals {
		switch local {
		case binary.ValueTypeI32:
			locals.push(ValueI32(0))
		case binary.ValueTypeI64:
			locals.push(ValueI64(0))
		}
	}

	arity := len(f.funcType.Results())

	frame := Frame{
		programCounter: -1,
		stackPointer:   r.stack.len(),
		instructions:   f.code.body,
		arity:          arity,
		locals:         locals,
	}

	r.callStack.push(&frame)
}

func (r *Runtime) invokeInternal(f InternalFuncInst) ([]Value, error) {
	arity := len(f.funcType.Results())

	r.pushFrame(f)

	if err := r.execute(); err != nil {
		r.cleanup()
		return nil, fmt.Errorf("failed to execute: %w", err)
	}

	if arity < 1 {
		return nil, nil
	}

	if r.stack.len() < arity {
		r.cleanup()
		return nil, fmt.Errorf("stack underflow")
	}

	returns := make([]Value, 0, arity)
	for range arity {
		returns = append(returns, r.stack.pop())
	}

	return returns, nil
}

func (r *Runtime) invokeExternal(f ExternalFuncInst) ([]Value, error) {
	bottom := r.stack.len() - len(f.funcType.Params())
	args := r.stack.splitOff(bottom)

	if f.module == "wasi_snapshot_preview1" && r.wasi != nil {
		return r.wasi.invoke(r.store, f.fn, args...)
	}

	module, ok := r.imports[f.module]
	if !ok {
		return nil, fmt.Errorf("module not found: %s", f.module)
	}
	fn, ok := module[f.fn]
	if !ok {
		return nil, fmt.Errorf("function not found: %s", f.fn)
	}
	return fn(r.store, args...)
}

func (r *Runtime) getEndAddress(insts []binary.Instruction, programCounter int) (int, error) {
	depth := 0
	for {
		programCounter++
		if programCounter < 0 || len(insts) <= programCounter {
			return 0, fmt.Errorf("unexpected end of instructions")
		}

		switch insts[programCounter].(type) {
		case *binary.InstructionIf:
			depth++
		case *binary.InstructionEnd:
			if depth == 0 {
				return programCounter, nil
			}
			depth--
		default:
			// do nothing
		}
	}
}

func (r *Runtime) cleanup() {
	r.callStack = nil
	r.stack = nil
}
