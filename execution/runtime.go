package execution

import (
	"fmt"
	"io"

	"github.com/Warashi/go-tinywasm/binary"
)

type Frame struct {
	programCounter int
	stackPointer   int
	instructions   []binary.Instruction
	arity          int
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

func (s *stack[T]) drain(n int) []T {
	r := (*s)[n:]
	*s = (*s)[:n]
	return r
}

func (s *stack[T]) splitOff(n int) stack[T] {
	r := (*s)[len(*s)-n:]
	*s = (*s)[:len(*s)-n]
	return r
}

func (s *stack[T]) len() int {
	return len(*s)
}

type Runtime struct {
	store     *Store
	stack     stack[Value]
	callStack stack[Frame]
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

func (r *Runtime) execute() error {
	for {
		if len(r.callStack) == 0 {
			break
		}

		frame := r.callStack[len(r.callStack)-1]

		if len(frame.instructions) <= frame.programCounter {
			break
		}

		inst := frame.instructions[frame.programCounter]

		switch inst := inst.(type) {
		case binary.InstructionEnd:
			if r.callStack.len() < 1 {
				return fmt.Errorf("call stack underflow")
			}

			frame := r.callStack.pop()
			if err := r.stackUnwind(frame.stackPointer, frame.arity); err != nil {
				return fmt.Errorf("failed to unwind stack: %w", err)
			}
		case binary.InstructionI32Add:
			if r.stack.len() < 2 {
				return fmt.Errorf("stack underflow")
			}
			right, left := r.stack.pop(), r.stack.pop()

			result, err := Add(left, right)
			if err != nil {
				return fmt.Errorf("failed to add: %w", err)
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

func (r *Runtime) invokeInternal(f InternalFuncInst) ([]Value, error) {
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

	r.callStack.push(frame)

	if err := r.execute(); err != nil {
		r.cleanup()
		return nil, fmt.Errorf("failed to execute: %w", err)
	}

	if arity < 1 {
		return nil, nil
	}

	return r.stack.drain(r.stack.len() - bottom), nil
}

func (r *Runtime) cleanup() {
	r.callStack = nil
	r.stack = nil
}
