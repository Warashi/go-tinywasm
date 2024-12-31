package runtime

import (
	"fmt"
	"slices"
)

type Instruction interface {
	Execute(*Runtime, *Frame) error
}

type Frame struct {
	ProgramCounter int
	StackPointer   int
	Instructions   []Instruction
	Arity          int
	Labels         Stack[Label]
	Locals         []Value
}

type Stack[T any] []T

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() T {
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}

func (s *Stack[T]) Drain(n int) {
	*s = (*s)[:n]
}

func (s *Stack[T]) SplitOff(n int) Stack[T] {
	r := (*s)[n:]
	*s = (*s)[:n]
	return slices.Clone(r)
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

type Runtime struct {
	store     *Store
	stack     Stack[Value]
	callStack Stack[*Frame]
	imports   Import
}

func (r *Runtime) execute() error {
	for len(r.callStack) > 0 {
		frame := r.callStack[len(r.callStack)-1]

		frame.ProgramCounter++

		if len(frame.Instructions) <= frame.ProgramCounter {
			break
		}

		instruction := frame.Instructions[frame.ProgramCounter]
		if err := instruction.Execute(r, frame); err != nil {
			return fmt.Errorf("failed to execute instruction: %w", err)
		}
	}

	return nil
}
