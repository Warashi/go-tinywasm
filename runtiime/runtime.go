package runtime

import (
	"fmt"

	"github.com/Warashi/go-tinywasm/stack"
	"github.com/Warashi/go-tinywasm/types/binary"
	"github.com/Warashi/go-tinywasm/types/runtime"
)

type Runtime struct {
	store     *Store
	stack     stack.Stack[runtime.Value]
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

// PopCallStack implements types.Runtime.
func (r *Runtime) PopCallStack() (*runtime.Frame, error) {
	if len(r.callStack) == 0 {
		return nil, runtime.ErrEmptyStack
	}

	return r.callStack.Pop(), nil
}

// PopStack implements types.Runtime.
func (r *Runtime) PopStack() (runtime.Value, error) {
	if len(r.stack) == 0 {
		return nil, runtime.ErrEmptyStack
	}

	return r.stack.Pop(), nil
}

// PushCallStack implements types.Runtime.
func (r *Runtime) PushCallStack(frame *runtime.Frame) {
	r.callStack.Push(frame)
}

// PushStack implements types.Runtime.
func (r *Runtime) PushStack(value runtime.Value) {
	r.stack.Push(value)
}

// SplitOffStack implements types.Runtime.
func (r *Runtime) SplitOffStack(n int) (stack.Stack[runtime.Value], error) {
	if len(r.stack) < n {
		return nil, runtime.ErrIndexOufOfRange
	}
	return r.stack.SplitOff(n), nil
}

// StackLen implements types.Runtime.
func (r *Runtime) StackLen() int {
	return r.stack.Len()
}

// StackUnwind implements types.Runtime.
func (r *Runtime) StackUnwind(stackPointer int, arity int) error {
	if arity == 0 {
		if r.stack.Len() < stackPointer {
			return fmt.Errorf("stack underflow")
		}
		r.stack.Drain(stackPointer)
		return nil
	}
	if r.stack.Len() < stackPointer+arity {
		return fmt.Errorf("stack underflow")
	}

	returns := make([]runtime.Value, 0, arity)
	for range arity {
		returns = append(returns, r.stack.Pop())
	}

	r.stack.Drain(stackPointer)

	for _, v := range returns {
		r.stack.Push(v)
	}

	return nil
}

// invokeInternal implements types.Runtime.
func (r *Runtime) InvokeInternal(f runtime.InternalFuncInst) ([]runtime.Value, error) {
	arity := len(f.FuncType.Results)

	r.PushFrame(f)

	if err := r.execute(); err != nil {
		r.Cleanup()
		return nil, fmt.Errorf("failed to execute: %w", err)
	}

	if arity < 1 {
		return nil, nil
	}

	if r.stack.Len() < arity {
		r.Cleanup()
		return nil, fmt.Errorf("stack underflow")
	}

	returns := make([]runtime.Value, 0, arity)
	for range arity {
		returns = append(returns, r.stack.Pop())
	}

	return returns, nil
}

// InvokeExternal implements types.Runtime.
func (r *Runtime) InvokeExternal(f runtime.ExternalFuncInst) ([]runtime.Value, error) {
	bottom := r.stack.Len() - len(f.FuncType.Params)
	args := r.stack.SplitOff(bottom)

	module, ok := r.imports[f.Module]
	if !ok {
		return nil, fmt.Errorf("module not found: %s", f.Module)
	}
	fn, ok := module[f.Func]
	if !ok {
		return nil, fmt.Errorf("function not found: %s", f.Func)
	}
	return fn(r.store, args...)
}

func (r *Runtime) PushFrame(f runtime.InternalFuncInst) error {
	bottom := r.StackLen() - len(f.FuncType.Params)
	locals, err := r.SplitOffStack(bottom)
	if err != nil {
		return fmt.Errorf("failed to split off stack: %w", err)
	}

	for _, local := range f.Code.Locals {
		switch local {
		case binary.ValueTypeI32:
			locals.Push(runtime.ValueI32(0))
		case binary.ValueTypeI64:
			locals.Push(runtime.ValueI64(0))
		}
	}

	arity := len(f.FuncType.Results)

	frame := runtime.Frame{
		ProgramCounter: -1,
		StackPointer:   r.StackLen(),
		Instructions:   f.Code.Body,
		Arity:          arity,
		Locals:         locals,
	}

	r.PushCallStack(&frame)

	return nil
}

func (r *Runtime) Cleanup() {
	r.stack = nil
	r.callStack = nil
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
