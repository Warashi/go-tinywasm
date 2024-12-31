package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/binary"
)

type I32GtS struct{}

func (i *I32GtS) Opcode() binary.Opcode {
	return binary.OpcodeI32GtS
}

func (i *I32GtS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I32GtU struct{}

func (i *I32GtU) Opcode() binary.Opcode {
	return binary.OpcodeI32GtU
}

func (i *I32GtU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64GtS struct{}

func (i *I64GtS) Opcode() binary.Opcode {
	return binary.OpcodeI64GtS
}

func (i *I64GtS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64GtU struct{}

func (i *I64GtU) Opcode() binary.Opcode {
	return binary.OpcodeI64GtU
}

func (i *I64GtU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32Gt struct{}

func (f *F32Gt) Opcode() binary.Opcode {
	return binary.OpcodeF32Gt
}

func (f *F32Gt) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64Gt struct{}

func (f *F64Gt) Opcode() binary.Opcode {
	return binary.OpcodeF64Gt
}

func (f *F64Gt) ReadOperandsFrom(r io.Reader) error {
	return nil
}
