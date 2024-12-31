package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/leb128"
	"github.com/Warashi/go-tinywasm/opcode"
	"github.com/Warashi/go-tinywasm/types/runtime"
)

type I32Const struct {
	value int32
}

func (i *I32Const) Value() int32 {
	return i.value
}

func (i *I32Const) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Const
}

func (i *I32Const) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.value, err = leb128.Int32(r)
	return err
}

func (i *I32Const) Execute(r runtime.Runtime, f runtime.Frame) error {
	r.PushStack(runtime.ValueI32(i.value))
	return nil
}

type I64Const struct {
	value int64
}

func (i *I64Const) Value() int64 {
	return i.value
}

func (i *I64Const) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Const
}

func (i *I64Const) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.value, err = leb128.Int64(r)
	return err
}

func (i *I64Const) Execute(r runtime.Runtime, f runtime.Frame) error {
	r.PushStack(runtime.ValueI64(i.value))
	return nil
}

type F32Const struct {
	value float32
}

func (i *F32Const) Value() float32 {
	return i.value
}

func (i *F32Const) Opcode() opcode.Opcode {
	return opcode.OpcodeF32Const
}

func (i *F32Const) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.value, err = readF32(r)
	return err
}

func (i *F32Const) Execute(r runtime.Runtime, f runtime.Frame) error {
	r.PushStack(runtime.ValueF32(i.value))
	return nil
}

type F64Const struct {
	value float64
}

func (i *F64Const) Value() float64 {
	return i.value
}

func (i *F64Const) Opcode() opcode.Opcode {
	return opcode.OpcodeF64Const
}

func (i *F64Const) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.value, err = readF64(r)
	return err
}

func (i *F64Const) Execute(r runtime.Runtime, f runtime.Frame) error {
	r.PushStack(runtime.ValueF64(i.value))
	return nil
}
