package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/opcode"
)

type I32Add struct{}

func (i *I32Add) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Add
}

func (i *I32Add) ReadOperandsFrom(r io.Reader) error {
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
