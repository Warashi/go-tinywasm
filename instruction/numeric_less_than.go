package instruction

import (
	"fmt"
	"io"

	"github.com/Warashi/go-tinywasm/opcode"
	"github.com/Warashi/go-tinywasm/types/runtime"
)

type I32LtS struct{}

func (i *I32LtS) Opcode() opcode.Opcode {
	return opcode.OpcodeI32LtS
}

func (i *I32LtS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

func (i *I32LtS) Execute(r runtime.Runtime, f runtime.Frame) error {
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

	if v1 < v2 {
		r.PushStack(runtime.ValueI32(1))
	} else {
		r.PushStack(runtime.ValueI32(0))
	}

	return nil
}

type I32LtU struct{}

func (i *I32LtU) Opcode() opcode.Opcode {
	return opcode.OpcodeI32LtU
}

func (i *I32LtU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64LtS struct{}

func (i *I64LtS) Opcode() opcode.Opcode {
	return opcode.OpcodeI64LtS
}

func (i *I64LtS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64LtU struct{}

func (i *I64LtU) Opcode() opcode.Opcode {
	return opcode.OpcodeI64LtU
}

func (i *I64LtU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32Lt struct{}

func (f *F32Lt) Opcode() opcode.Opcode {
	return opcode.OpcodeF32Lt
}

func (f *F32Lt) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64Lt struct{}

func (f *F64Lt) Opcode() opcode.Opcode {
	return opcode.OpcodeF64Lt
}

func (f *F64Lt) ReadOperandsFrom(r io.Reader) error {
	return nil
}
