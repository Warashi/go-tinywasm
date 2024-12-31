package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/binary"
)

type I32Sub struct{}

func (i *I32Sub) Opcode() binary.Opcode {
	return binary.OpcodeI32Sub
}

func (i *I32Sub) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64Sub struct{}

func (i *I64Sub) Opcode() binary.Opcode {
	return binary.OpcodeI64Sub
}

func (i *I64Sub) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32Sub struct{}

func (f *F32Sub) Opcode() binary.Opcode {
	return binary.OpcodeF32Sub
}

func (f *F32Sub) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64Sub struct{}

func (f *F64Sub) Opcode() binary.Opcode {
	return binary.OpcodeF64Sub
}

func (f *F64Sub) ReadOperandsFrom(r io.Reader) error {
	return nil
}
