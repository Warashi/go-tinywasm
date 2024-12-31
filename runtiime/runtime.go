package runtime

import (
	"fmt"

	"github.com/Warashi/go-tinywasm/stack"
	"github.com/Warashi/go-tinywasm/types/runtime"
)

type Runtime struct {
	store     *Store
	stack     stack.Stack[Value]
	callStack stack.Stack[*runtime.Frame]
	imports   Import
}

func (r *Runtime) WriteMemoryAt(n int, data []byte, offset int64) (int, error) {
	if n < 0 || len(r.store.memories) <= n {
		return 0, fmt.Errorf("invalid memory index: %d", n)
	}

	return r.store.memories[n].WriteAt(data, offset)
}

// Func implements types.Runtime.
func (r *Runtime) Func(i int) (runtime.FuncInst, error) {
	if i < 0 || len(r.store.funcs) <= i {
		return nil, fmt.Errorf("invalid function index: %d", i)
	}
	return r.store.funcs[i], nil
}

// InvokeExternal implements types.Runtime.
func (r *Runtime) InvokeExternal(runtime.ExternalFuncInst) ([]runtime.Value, error) {
	panic("unimplemented")
}

// PopCallStack implements types.Runtime.
func (r *Runtime) PopCallStack() (*runtime.Frame, error) {
	panic("unimplemented")
}

// PopLabel implements types.Runtime.
func (r *Runtime) PopLabel() (runtime.Label, error) {
	panic("unimplemented")
}

// PopStack implements types.Runtime.
func (r *Runtime) PopStack() (runtime.Value, error) {
	panic("unimplemented")
}

// PushCallStack implements types.Runtime.
func (r *Runtime) PushCallStack(*runtime.Frame) {
	panic("unimplemented")
}

// PushStack implements types.Runtime.
func (r *Runtime) PushStack(runtime.Value) {
	panic("unimplemented")
}

// SplitOffStack implements types.Runtime.
func (r *Runtime) SplitOffStack(n int) (stack.Stack[runtime.Value], error) {
	panic("unimplemented")
}

// StackLen implements types.Runtime.
func (r *Runtime) StackLen() int {
	panic("unimplemented")
}

// StackUnwind implements types.Runtime.
func (r *Runtime) StackUnwind(stackPointer int, arity int) error {
	panic("unimplemented")
}

// invokeInternal implements types.Runtime.
func (r *Runtime) InvokeInternal(runtime.InternalFuncInst) ([]runtime.Value, error) {
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
