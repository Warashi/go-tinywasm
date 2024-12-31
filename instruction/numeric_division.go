package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/opcode"
)

type I32DivS struct{}

func (i *I32DivS) Opcode() opcode.Opcode {
	return opcode.OpcodeI32DivS
}

func (i *I32DivS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I32DivU struct{}

func (i *I32DivU) Opcode() opcode.Opcode {
	return opcode.OpcodeI32DivU
}

func (i *I32DivU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64DivS struct{}

func (i *I64DivS) Opcode() opcode.Opcode {
	return opcode.OpcodeI64DivS
}

func (i *I64DivS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64DivU struct{}

func (i *I64DivU) Opcode() opcode.Opcode {
	return opcode.OpcodeI64DivU
}

func (i *I64DivU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32Div struct{}

func (f *F32Div) Opcode() opcode.Opcode {
	return opcode.OpcodeF32Div
}

func (f *F32Div) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64Div struct{}

func (f *F64Div) Opcode() opcode.Opcode {
	return opcode.OpcodeF64Div
}

func (f *F64Div) ReadOperandsFrom(r io.Reader) error {
	return nil
}
