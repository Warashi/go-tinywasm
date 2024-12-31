package instruction

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/Warashi/wasmium/leb128"
	"github.com/Warashi/wasmium/opcode"
	"github.com/Warashi/wasmium/types/runtime"
)

type I32Store struct {
	align  uint32
	offset uint32
}

func (i *I32Store) Align() uint32 {
	return i.align
}

func (i *I32Store) Offset() uint32 {
	return i.offset
}

func (i *I32Store) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Store
}

func (i *I32Store) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.align, err = leb128.Uint32(r)
	if err != nil {
		return err
	}
	i.offset, err = leb128.Uint32(r)
	return err
}

func (i *I32Store) Execute(r runtime.Runtime, f *runtime.Frame) error {
	value, err := r.PopStack()
	if err != nil {
		return err
	}

	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	v, ok := value.(runtime.ValueI32)
	if !ok {
		return runtime.ErrInvalidValue
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return runtime.ErrInvalidValue
	}

	var buf [4]byte
	if _, err := binary.Encode(buf[:], endian, int32(v)); err != nil {
		return fmt.Errorf("failed to encode value: %w", err)
	}

	if n, err := r.WriteMemoryAt(0, buf[:], int64(uint32(a)+i.offset)); err != nil || n != 4 {
		return fmt.Errorf("failed to write memory(%d): %w", n, err)
	}

	return nil
}

// TODO: implement the rest of the store instructions
