package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/binary"
)

type I32Eqz struct{}

func (i *I32Eqz) Opcode() binary.Opcode {
	return binary.OpcodeI32Eqz
}

func (i *I32Eqz) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I32Eq struct{}

func (i *I32Eq) Opcode() binary.Opcode {
	return binary.OpcodeI32Eq
}

func (i *I32Eq) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64Eqz struct{}

func (i *I64Eqz) Opcode() binary.Opcode {
	return binary.OpcodeI64Eqz
}

func (i *I64Eqz) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64Eq struct{}

func (i *I64Eq) Opcode() binary.Opcode {
	return binary.OpcodeI64Eq
}

func (i *I64Eq) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32Eq struct{}

func (f *F32Eq) Opcode() binary.Opcode {
	return binary.OpcodeF32Eq
}

func (f *F32Eq) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64Eq struct{}

func (f *F64Eq) Opcode() binary.Opcode {
	return binary.OpcodeF64Eq
}

func (f *F64Eq) ReadOperandsFrom(r io.Reader) error {
	return nil
}
