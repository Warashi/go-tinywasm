package instruction

import (
	"fmt"
	"io"

	"github.com/Warashi/go-tinywasm/opcode"
	"github.com/Warashi/go-tinywasm/types/runtime"
)

type I32Sub struct{}

func (i *I32Sub) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Sub
}

func (i *I32Sub) ReadOperandsFrom(r io.Reader) error {
	return nil
}

func (i *I32Sub) Execute(r runtime.Runtime, f *runtime.Frame) error {
	right, err := r.PopStack()
	if err != nil {
		return fmt.Errorf("failed to pop stack: %w", err)
	}

	left, err := r.PopStack()
	if err != nil {
		return fmt.Errorf("failed to pop stack: %w", err)
	}

	v1, ok := left.(runtime.ValueI32)
	if !ok {
		return runtime.ErrInvalidValue
	}

	v2, ok := right.(runtime.ValueI32)
	if !ok {
		return runtime.ErrInvalidValue
	}

	r.PushStack(runtime.ValueI32(v1 - v2))

	return nil
}

type I64Sub struct{}

func (i *I64Sub) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Sub
}

func (i *I64Sub) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32Sub struct{}

func (f *F32Sub) Opcode() opcode.Opcode {
	return opcode.OpcodeF32Sub
}

func (f *F32Sub) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64Sub struct{}

func (f *F64Sub) Opcode() opcode.Opcode {
	return opcode.OpcodeF64Sub
}

func (f *F64Sub) ReadOperandsFrom(r io.Reader) error {
	return nil
}
