package instruction

import (
	"io"

	"github.com/Warashi/wasmium/opcode"
)

type I32Ne struct{}

func (i *I32Ne) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Ne
}

func (i *I32Ne) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64Ne struct{}

func (i *I64Ne) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Ne
}

func (i *I64Ne) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32Ne struct{}

func (f *F32Ne) Opcode() opcode.Opcode {
	return opcode.OpcodeF32Ne
}

func (f *F32Ne) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64Ne struct{}

func (f *F64Ne) Opcode() opcode.Opcode {
	return opcode.OpcodeF64Ne
}

func (f *F64Ne) ReadOperandsFrom(r io.Reader) error {
	return nil
}
