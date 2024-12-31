package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/leb128"
	"github.com/Warashi/go-tinywasm/opcode"
	"github.com/Warashi/go-tinywasm/types/runtime"
)

type LocalGet struct {
	index uint32
}

func (i *LocalGet) Index() uint32 {
	return i.index
}

func (i *LocalGet) Opcode() opcode.Opcode {
	return opcode.OpcodeLocalGet
}

func (i *LocalGet) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.index, err = leb128.Uint32(r)
	return err
}

func (i *LocalGet) Execute(r runtime.Runtime, f *runtime.Frame) error {
	if i.index < 0 || len(f.Locals) <= int(i.index) {
		return runtime.ErrOutOfBounds
	}

	r.PushStack(f.Locals[i.index])

	return nil
}

type LocalSet struct {
	index uint32
}

func (i *LocalSet) Index() uint32 {
	return i.index
}

func (i *LocalSet) Opcode() opcode.Opcode {
	return opcode.OpcodeLocalSet
}

func (i *LocalSet) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.index, err = leb128.Uint32(r)
	return err
}

func (i *LocalSet) Execute(r runtime.Runtime, f *runtime.Frame) error {
	if i.index < 0 || len(f.Locals) <= int(i.index) {
		return runtime.ErrOutOfBounds
	}

	value, err := r.PopStack()
	if err != nil {
		return err
	}

	f.Locals[i.index] = value

	return nil
}
