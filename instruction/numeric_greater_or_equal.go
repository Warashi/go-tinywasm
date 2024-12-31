package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/binary"
)

type I32GeS struct{}

func (i *I32GeS) Opcode() binary.Opcode {
	return binary.OpcodeI32GeS
}

func (i *I32GeS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I32GeU struct{}

func (i *I32GeU) Opcode() binary.Opcode {
	return binary.OpcodeI32GeU
}

func (i *I32GeU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64GeS struct{}

func (i *I64GeS) Opcode() binary.Opcode {
	return binary.OpcodeI64GeS
}

func (i *I64GeS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64GeU struct{}

func (i *I64GeU) Opcode() binary.Opcode {
	return binary.OpcodeI64GeU
}

func (i *I64GeU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32Ge struct{}

func (f *F32Ge) Opcode() binary.Opcode {
	return binary.OpcodeF32Ge
}

func (f *F32Ge) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64Ge struct{}

func (f *F64Ge) Opcode() binary.Opcode {
	return binary.OpcodeF64Ge
}

func (f *F64Ge) ReadOperandsFrom(r io.Reader) error {
	return nil
}
