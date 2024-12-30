package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/binary"
)

type I32LeS struct{}

func (i *I32LeS) Opcode() binary.Opcode {
	return binary.OpcodeI32LeS
}

func (i *I32LeS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I32LeU struct{}

func (i *I32LeU) Opcode() binary.Opcode {
	return binary.OpcodeI32LeU
}

func (i *I32LeU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64LeS struct{}

func (i *I64LeS) Opcode() binary.Opcode {
	return binary.OpcodeI64LeS
}

func (i *I64LeS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64LeU struct{}

func (i *I64LeU) Opcode() binary.Opcode {
	return binary.OpcodeI64LeU
}

func (i *I64LeU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32Le struct{}

func (f *F32Le) Opcode() binary.Opcode {
	return binary.OpcodeF32Le
}

func (f *F32Le) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64Le struct{}

func (f *F64Le) Opcode() binary.Opcode {
	return binary.OpcodeF64Le
}

func (f *F64Le) ReadOperandsFrom(r io.Reader) error {
	return nil
}
