package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/binary"
)

type I32ReinterpretF32 struct{}

func (i *I32ReinterpretF32) Opcode() binary.Opcode {
	return binary.OpcodeI32ReinterpretF32
}

func (i *I32ReinterpretF32) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64ReinterpretF64 struct{}

func (i *I64ReinterpretF64) Opcode() binary.Opcode {
	return binary.OpcodeI64ReinterpretF64
}

func (i *I64ReinterpretF64) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32ReinterpretI32 struct{}

func (f *F32ReinterpretI32) Opcode() binary.Opcode {
	return binary.OpcodeF32ReinterpretI32
}

func (f *F32ReinterpretI32) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64ReinterpretI64 struct{}

func (f *F64ReinterpretI64) Opcode() binary.Opcode {
	return binary.OpcodeF64ReinterpretI64
}

func (f *F64ReinterpretI64) ReadOperandsFrom(r io.Reader) error {
	return nil
}
