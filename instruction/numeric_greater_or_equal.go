package instruction

import (
	"io"

	"github.com/Warashi/wasmium/opcode"
)

type I32GeS struct{}

func (i *I32GeS) Opcode() opcode.Opcode {
	return opcode.OpcodeI32GeS
}

func (i *I32GeS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I32GeU struct{}

func (i *I32GeU) Opcode() opcode.Opcode {
	return opcode.OpcodeI32GeU
}

func (i *I32GeU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64GeS struct{}

func (i *I64GeS) Opcode() opcode.Opcode {
	return opcode.OpcodeI64GeS
}

func (i *I64GeS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64GeU struct{}

func (i *I64GeU) Opcode() opcode.Opcode {
	return opcode.OpcodeI64GeU
}

func (i *I64GeU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32Ge struct{}

func (f *F32Ge) Opcode() opcode.Opcode {
	return opcode.OpcodeF32Ge
}

func (f *F32Ge) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64Ge struct{}

func (f *F64Ge) Opcode() opcode.Opcode {
	return opcode.OpcodeF64Ge
}

func (f *F64Ge) ReadOperandsFrom(r io.Reader) error {
	return nil
}
