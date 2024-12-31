package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/binary"
)

type I32LtS struct{}

func (i *I32LtS) Opcode() binary.Opcode {
	return binary.OpcodeI32LtS
}

func (i *I32LtS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I32LtU struct{}

func (i *I32LtU) Opcode() binary.Opcode {
	return binary.OpcodeI32LtU
}

func (i *I32LtU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64LtS struct{}

func (i *I64LtS) Opcode() binary.Opcode {
	return binary.OpcodeI64LtS
}

func (i *I64LtS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64LtU struct{}

func (i *I64LtU) Opcode() binary.Opcode {
	return binary.OpcodeI64LtU
}

func (i *I64LtU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32Lt struct{}

func (f *F32Lt) Opcode() binary.Opcode {
	return binary.OpcodeF32Lt
}

func (f *F32Lt) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64Lt struct{}

func (f *F64Lt) Opcode() binary.Opcode {
	return binary.OpcodeF64Lt
}

func (f *F64Lt) ReadOperandsFrom(r io.Reader) error {
	return nil
}
