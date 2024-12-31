package instruction

import (
	"fmt"
	"io"

	"github.com/Warashi/go-tinywasm/opcode"
	"github.com/Warashi/go-tinywasm/types/runtime"
)

type I32Add struct{}

func (i *I32Add) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Add
}

func (i *I32Add) ReadOperandsFrom(r io.Reader) error {
	return nil
}

func (i *I32Add) Execute(r runtime.Runtime, f runtime.Frame) error {
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

	r.PushStack(runtime.ValueI32(v1 + v2))

	return nil
}

type I64Add struct{}

func (i *I64Add) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Add
}

func (i *I64Add) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32Add struct{}

func (f *F32Add) Opcode() opcode.Opcode {
	return opcode.OpcodeF32Add
}

func (f *F32Add) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64Add struct{}

func (f *F64Add) Opcode() opcode.Opcode {
	return opcode.OpcodeF64Add
}

func (f *F64Add) ReadOperandsFrom(r io.Reader) error {
	return nil
}
