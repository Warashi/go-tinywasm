package instruction

import (
	"io"

	"github.com/Warashi/wasmium/leb128"
	"github.com/Warashi/wasmium/opcode"
	"github.com/Warashi/wasmium/types/runtime"
)

type LocalGet struct {
	Index uint32
}

func (i *LocalGet) Opcode() opcode.Opcode {
	return opcode.OpcodeLocalGet
}

func (i *LocalGet) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Index, err = leb128.Uint32(r)
	return err
}

func (i *LocalGet) Execute(r runtime.Runtime, f *runtime.Frame) error {
	if i.Index < 0 || len(f.Locals) <= int(i.Index) {
		return runtime.ErrOutOfBounds
	}

	r.PushStack(f.Locals[i.Index])

	return nil
}

type LocalSet struct {
	Index uint32
}

func (i *LocalSet) Opcode() opcode.Opcode {
	return opcode.OpcodeLocalSet
}

func (i *LocalSet) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Index, err = leb128.Uint32(r)
	return err
}

func (i *LocalSet) Execute(r runtime.Runtime, f *runtime.Frame) error {
	if i.Index < 0 || len(f.Locals) <= int(i.Index) {
		return runtime.ErrOutOfBounds
	}

	value, err := r.PopStack()
	if err != nil {
		return err
	}

	f.Locals[i.Index] = value

	return nil
}
