package runtime

import (
	"fmt"

	"github.com/Warashi/go-tinywasm/stack"
)

var (
	ErrEmptyStack      = fmt.Errorf("stack is empty")
	ErrIndexOufOfRange = fmt.Errorf("index out of range")
)

type Runtime interface {
	PopStack() (Value, error)
	PopLabel() (Label, error)
	PopCallStack() (*Frame, error)

	PushStack(Value)
	PushCallStack(*Frame)

	SplitOffStack(n int) (stack.Stack[Value], error)

	StackLen() int
	StackUnwind(stackPointer, arity int) error

	Func(i int) (FuncInst, error)
	InvokeInternal(InternalFuncInst) ([]Value, error)
	InvokeExternal(ExternalFuncInst) ([]Value, error)

	WriteMemoryAt(n int, data []byte, offset int64) (int, error)
}
