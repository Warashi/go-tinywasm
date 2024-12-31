package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/opcode"
)

type I32ReinterpretF32 struct{}

func (i *I32ReinterpretF32) Opcode() opcode.Opcode {
	return opcode.OpcodeI32ReinterpretF32
}

func (i *I32ReinterpretF32) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64ReinterpretF64 struct{}

func (i *I64ReinterpretF64) Opcode() opcode.Opcode {
	return opcode.OpcodeI64ReinterpretF64
}

func (i *I64ReinterpretF64) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32ReinterpretI32 struct{}

func (f *F32ReinterpretI32) Opcode() opcode.Opcode {
	return opcode.OpcodeF32ReinterpretI32
}

func (f *F32ReinterpretI32) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64ReinterpretI64 struct{}

func (f *F64ReinterpretI64) Opcode() opcode.Opcode {
	return opcode.OpcodeF64ReinterpretI64
}

func (f *F64ReinterpretI64) ReadOperandsFrom(r io.Reader) error {
	return nil
}
