package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/binary"
)

type I32TruncF32S struct{}

func (i *I32TruncF32S) Opcode() binary.Opcode {
	return binary.OpcodeI32TruncF32S
}

func (i *I32TruncF32S) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I32TruncF32U struct{}

func (i *I32TruncF32U) Opcode() binary.Opcode {
	return binary.OpcodeI32TruncF32U
}

func (i *I32TruncF32U) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I32TruncF64S struct{}

func (i *I32TruncF64S) Opcode() binary.Opcode {
	return binary.OpcodeI32TruncF64S
}

func (i *I32TruncF64S) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I32TruncF64U struct{}

func (i *I32TruncF64U) Opcode() binary.Opcode {
	return binary.OpcodeI32TruncF64U
}

func (i *I32TruncF64U) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64TruncF32S struct{}

func (i *I64TruncF32S) Opcode() binary.Opcode {
	return binary.OpcodeI64TruncF32S
}

func (i *I64TruncF32S) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64TruncF32U struct{}

func (i *I64TruncF32U) Opcode() binary.Opcode {
	return binary.OpcodeI64TruncF32U
}

func (i *I64TruncF32U) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64TruncF64S struct{}

func (i *I64TruncF64S) Opcode() binary.Opcode {
	return binary.OpcodeI64TruncF64S
}

func (i *I64TruncF64S) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64TruncF64U struct{}

func (i *I64TruncF64U) Opcode() binary.Opcode {
	return binary.OpcodeI64TruncF64U
}

func (i *I64TruncF64U) ReadOperandsFrom(r io.Reader) error {
	return nil
}
