package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/opcode"
)

type I32GtS struct{}

func (i *I32GtS) Opcode() opcode.Opcode {
	return opcode.OpcodeI32GtS
}

func (i *I32GtS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I32GtU struct{}

func (i *I32GtU) Opcode() opcode.Opcode {
	return opcode.OpcodeI32GtU
}

func (i *I32GtU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64GtS struct{}

func (i *I64GtS) Opcode() opcode.Opcode {
	return opcode.OpcodeI64GtS
}

func (i *I64GtS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64GtU struct{}

func (i *I64GtU) Opcode() opcode.Opcode {
	return opcode.OpcodeI64GtU
}

func (i *I64GtU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32Gt struct{}

func (f *F32Gt) Opcode() opcode.Opcode {
	return opcode.OpcodeF32Gt
}

func (f *F32Gt) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64Gt struct{}

func (f *F64Gt) Opcode() opcode.Opcode {
	return opcode.OpcodeF64Gt
}

func (f *F64Gt) ReadOperandsFrom(r io.Reader) error {
	return nil
}
