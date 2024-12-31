package instruction

import (
	"fmt"
	"io"

	"github.com/Warashi/wasmium/leb128"
	"github.com/Warashi/wasmium/opcode"
	"github.com/Warashi/wasmium/types/runtime"
)

type GlobalGet struct {
	Index uint32
}

func (i *GlobalGet) Opcode() opcode.Opcode {
	return opcode.OpcodeGlobalGet
}

func (i *GlobalGet) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Index, err = leb128.Uint32(r)
	return err
}

func (i *GlobalGet) Execute(r runtime.Runtime, f *runtime.Frame) error {
	v, err := r.GlobalGet(int(i.Index))
	if err != nil {
		return fmt.Errorf("failed to get global: %w", err)
	}

	r.PushStack(v)

	return nil
}

type GlobalSet struct {
	Index uint32
}

func (i *GlobalSet) Opcode() opcode.Opcode {
	return opcode.OpcodeGlobalSet
}

func (i *GlobalSet) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Index, err = leb128.Uint32(r)
	return err
}

func (i *GlobalSet) Execute(r runtime.Runtime, f *runtime.Frame) error {
	v, err := r.PopStack()
	if err != nil {
		return fmt.Errorf("failed to pop stack: %w", err)
	}

	if err := r.GlobalSet(int(i.Index), v); err != nil {
		return fmt.Errorf("failed to set global: %w", err)
	}

	return nil
}
