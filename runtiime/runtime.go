package runtime

import (
	"fmt"

	"github.com/Warashi/go-tinywasm/interfaces"
	"github.com/Warashi/go-tinywasm/stack"
)

type Runtime struct {
	store     *Store
	stack     stack.Stack[Value]
	callStack stack.Stack[*interfaces.Frame]
	imports   Import
}

// Func implements interfaces.Runtime.
func (r *Runtime) Func(i int) (interfaces.FuncInst, error) {
	panic("unimplemented")
}

// InvokeExternal implements interfaces.Runtime.
func (r *Runtime) InvokeExternal(interfaces.ExternalFuncInst) ([]interfaces.Value, error) {
	panic("unimplemented")
}

// PopCallStack implements interfaces.Runtime.
func (r *Runtime) PopCallStack() (*interfaces.Frame, error) {
	panic("unimplemented")
}

// PopLabel implements interfaces.Runtime.
func (r *Runtime) PopLabel() (interfaces.Label, error) {
	panic("unimplemented")
}

// PopStack implements interfaces.Runtime.
func (r *Runtime) PopStack() (interfaces.Value, error) {
	panic("unimplemented")
}

// PushCallStack implements interfaces.Runtime.
func (r *Runtime) PushCallStack(*interfaces.Frame) {
	panic("unimplemented")
}

// PushStack implements interfaces.Runtime.
func (r *Runtime) PushStack(interfaces.Value) {
	panic("unimplemented")
}

// SplitOffStack implements interfaces.Runtime.
func (r *Runtime) SplitOffStack(n int) (stack.Stack[interfaces.Value], error) {
	panic("unimplemented")
}

// StackLen implements interfaces.Runtime.
func (r *Runtime) StackLen() int {
	panic("unimplemented")
}

// StackUnwind implements interfaces.Runtime.
func (r *Runtime) StackUnwind(stackPointer int, arity int) error {
	panic("unimplemented")
}

// invokeInternal implements interfaces.Runtime.
func (r *Runtime) InvokeInternal(interfaces.InternalFuncInst) ([]interfaces.Value, error) {
	panic("unimplemented")
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
