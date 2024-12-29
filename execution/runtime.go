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
