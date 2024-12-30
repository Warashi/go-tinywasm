package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/binary"
)

type I32DivS struct{}

func (i *I32DivS) Opcode() binary.Opcode {
	return binary.OpcodeI32DivS
}

func (i *I32DivS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I32DivU struct{}

func (i *I32DivU) Opcode() binary.Opcode {
	return binary.OpcodeI32DivU
}

func (i *I32DivU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64DivS struct{}

func (i *I64DivS) Opcode() binary.Opcode {
	return binary.OpcodeI64DivS
}

func (i *I64DivS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64DivU struct{}

func (i *I64DivU) Opcode() binary.Opcode {
	return binary.OpcodeI64DivU
}

func (i *I64DivU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32Div struct{}

func (f *F32Div) Opcode() binary.Opcode {
	return binary.OpcodeF32Div
}

func (f *F32Div) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64Div struct{}

func (f *F64Div) Opcode() binary.Opcode {
	return binary.OpcodeF64Div
}

func (f *F64Div) ReadOperandsFrom(r io.Reader) error {
	return nil
}
