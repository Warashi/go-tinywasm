package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/binary"
)

type F32ConvertI32S struct{}

func (f *F32ConvertI32S) Opcode() binary.Opcode {
	return binary.OpcodeF32ConvertI32S
}

func (f *F32ConvertI32S) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32ConvertI32U struct{}

func (f *F32ConvertI32U) Opcode() binary.Opcode {
	return binary.OpcodeF32ConvertI32U
}

func (f *F32ConvertI32U) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32ConvertI64S struct{}

func (f *F32ConvertI64S) Opcode() binary.Opcode {
	return binary.OpcodeF32ConvertI64S
}

func (f *F32ConvertI64S) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F32ConvertI64U struct{}

func (f *F32ConvertI64U) Opcode() binary.Opcode {
	return binary.OpcodeF32ConvertI64U
}

func (f *F32ConvertI64U) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64ConvertI32S struct{}

func (f *F64ConvertI32S) Opcode() binary.Opcode {
	return binary.OpcodeF64ConvertI32S
}

func (f *F64ConvertI32S) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64ConvertI32U struct{}

func (f *F64ConvertI32U) Opcode() binary.Opcode {
	return binary.OpcodeF64ConvertI32U
}

func (f *F64ConvertI32U) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64ConvertI64S struct{}

func (f *F64ConvertI64S) Opcode() binary.Opcode {
	return binary.OpcodeF64ConvertI64S
}

func (f *F64ConvertI64S) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type F64ConvertI64U struct{}

func (f *F64ConvertI64U) Opcode() binary.Opcode {
	return binary.OpcodeF64ConvertI64U
}

func (f *F64ConvertI64U) ReadOperandsFrom(r io.Reader) error {
	return nil
}
