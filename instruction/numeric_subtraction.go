package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/opcode"
)

type I32Sub struct{}

func (i *I32Sub) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Sub
}

func (i *I32Sub) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64Sub struct{}

func (i *I64Sub) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Sub
}

func (i *I64Sub) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32Sub struct{}

func (f *F32Sub) Opcode() opcode.Opcode {
	return opcode.OpcodeF32Sub
}

func (f *F32Sub) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64Sub struct{}

func (f *F64Sub) Opcode() opcode.Opcode {
	return opcode.OpcodeF64Sub
}

func (f *F64Sub) ReadOperandsFrom(r io.Reader) error {
	return nil
}
