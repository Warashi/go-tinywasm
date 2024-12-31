package runtime

import (
	"github.com/Warashi/go-tinywasm/stack"
)

type Instruction interface {
	Execute(Runtime, *Frame) error
}

type Frame struct {
	ProgramCounter int
	StackPointer   int
	Instructions   []Instruction
	Arity          int
	Labels         stack.Stack[Label]
	Locals         []Value
}
