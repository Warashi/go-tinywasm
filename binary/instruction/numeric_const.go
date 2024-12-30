package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/binary"
	"github.com/Warashi/go-tinywasm/leb128"
)

type I32Const struct {
	value int32
}

func (i *I32Const) Value() int32 {
	return i.value
}

func (i *I32Const) Opcode() binary.Opcode {
	return binary.OpcodeI32Const
}

func (i *I32Const) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.value, err = leb128.Int32(r)
	return err
}

type I64Const struct {
	value int64
}

func (i *I64Const) Value() int64 {
	return i.value
}

func (i *I64Const) Opcode() binary.Opcode {
	return binary.OpcodeI64Const
}

func (i *I64Const) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.value, err = leb128.Int64(r)
	return err
}

type F32Const struct {
	value float32
}

func (f *F32Const) Value() float32 {
	return f.value
}

func (f *F32Const) Opcode() binary.Opcode {
	return binary.OpcodeF32Const
}

func (f *F32Const) ReadOperandsFrom(r io.Reader) error {
	var err error
	f.value, err = readF32(r)
	return err
}

type F64Const struct {
	value float64
}

func (f *F64Const) Value() float64 {
	return f.value
}

func (f *F64Const) Opcode() binary.Opcode {
	return binary.OpcodeF64Const
}

func (f *F64Const) ReadOperandsFrom(r io.Reader) error {
	var err error
	f.value, err = readF64(r)
	return err
}
