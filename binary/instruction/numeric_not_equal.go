package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/binary"
)

type I32Ne struct{}

func (i *I32Ne) Opcode() binary.Opcode {
	return binary.OpcodeI32Ne
}

func (i *I32Ne) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64Ne struct{}

func (i *I64Ne) Opcode() binary.Opcode {
	return binary.OpcodeI64Ne
}

func (i *I64Ne) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32Ne struct{}

func (f *F32Ne) Opcode() binary.Opcode {
	return binary.OpcodeF32Ne
}

func (f *F32Ne) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64Ne struct{}

func (f *F64Ne) Opcode() binary.Opcode {
	return binary.OpcodeF64Ne
}

func (f *F64Ne) ReadOperandsFrom(r io.Reader) error {
	return nil
}
