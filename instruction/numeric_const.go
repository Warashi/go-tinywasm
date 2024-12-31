package instruction

import (
	"io"

	"github.com/Warashi/wasmium/leb128"
	"github.com/Warashi/wasmium/opcode"
	"github.com/Warashi/wasmium/types/runtime"
)

type I32Const struct {
	Value int32
}

func (i *I32Const) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Const
}

func (i *I32Const) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Value, err = leb128.Int32(r)
	return err
}

func (i *I32Const) Execute(r runtime.Runtime, f *runtime.Frame) error {
	r.PushStack(runtime.ValueI32(i.Value))
	return nil
}

type I64Const struct {
	Value int64
}

func (i *I64Const) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Const
}

func (i *I64Const) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Value, err = leb128.Int64(r)
	return err
}

func (i *I64Const) Execute(r runtime.Runtime, f *runtime.Frame) error {
	r.PushStack(runtime.ValueI64(i.Value))
	return nil
}

type F32Const struct {
	Value float32
}

func (i *F32Const) Opcode() opcode.Opcode {
	return opcode.OpcodeF32Const
}

func (i *F32Const) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Value, err = readF32(r)
	return err
}

func (i *F32Const) Execute(r runtime.Runtime, f *runtime.Frame) error {
	r.PushStack(runtime.ValueF32(i.Value))
	return nil
}

type F64Const struct {
	Value float64
}

func (i *F64Const) Opcode() opcode.Opcode {
	return opcode.OpcodeF64Const
}

func (i *F64Const) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Value, err = readF64(r)
	return err
}

func (i *F64Const) Execute(r runtime.Runtime, f *runtime.Frame) error {
	r.PushStack(runtime.ValueF64(i.Value))
	return nil
}
