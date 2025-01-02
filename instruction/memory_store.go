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
	Align  uint32
	Offset uint32
}

func (i *I32Store) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Store
}

func (i *I32Store) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return err
	}
	i.Offset, err = leb128.Uint32(r)
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
		return fmt.Errorf("invalid value(%T): %w", value, runtime.ErrInvalidValue)
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [4]byte
	if _, err := binary.Encode(buf[:], endian, int32(v)); err != nil {
		return fmt.Errorf("failed to encode value: %w", err)
	}

	if n, err := r.WriteMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to write memory(%d): %w", n, err)
	}

	return nil
}

type I64Store struct {
	Align  uint32
	Offset uint32
}

func (i *I64Store) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Store
}

func (i *I64Store) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return err
	}
	i.Offset, err = leb128.Uint32(r)
	return err
}

func (i *I64Store) Execute(r runtime.Runtime, f *runtime.Frame) error {
	value, err := r.PopStack()
	if err != nil {
		return err
	}

	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	v, ok := value.(runtime.ValueI64)
	if !ok {
		return fmt.Errorf("invalid value(%T): %w", value, runtime.ErrInvalidValue)
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [8]byte
	if _, err := binary.Encode(buf[:], endian, int64(v)); err != nil {
		return fmt.Errorf("failed to encode value: %w", err)
	}

	if n, err := r.WriteMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to write memory(%d): %w", n, err)
	}

	return nil
}

type I32Store8 struct {
	Align  uint32
	Offset uint32
}

func (i *I32Store8) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Store8
}

func (i *I32Store8) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return err
	}
	i.Offset, err = leb128.Uint32(r)
	return err
}

func (i *I32Store8) Execute(r runtime.Runtime, f *runtime.Frame) error {
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
		return fmt.Errorf("invalid value(%T): %w", value, runtime.ErrInvalidValue)
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [1]byte
	if _, err := binary.Encode(buf[:], endian, int8(v)); err != nil {
		return fmt.Errorf("failed to encode value: %w", err)
	}

	if n, err := r.WriteMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to write memory(%d): %w", n, err)
	}

	return nil
}

type I32Store16 struct {
	Align  uint32
	Offset uint32
}

func (i *I32Store16) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Store16
}

func (i *I32Store16) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return err
	}
	i.Offset, err = leb128.Uint32(r)
	return err
}

func (i *I32Store16) Execute(r runtime.Runtime, f *runtime.Frame) error {
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
		return fmt.Errorf("invalid value(%T): %w", value, runtime.ErrInvalidValue)
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [2]byte
	if _, err := binary.Encode(buf[:], endian, int16(v)); err != nil {
		return fmt.Errorf("failed to encode value: %w", err)
	}

	if n, err := r.WriteMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to write memory(%d): %w", n, err)
	}

	return nil
}

type I64Store8 struct {
	Align  uint32
	Offset uint32
}

func (i *I64Store8) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Store8
}

func (i *I64Store8) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return err
	}
	i.Offset, err = leb128.Uint32(r)
	return err
}

func (i *I64Store8) Execute(r runtime.Runtime, f *runtime.Frame) error {
	value, err := r.PopStack()
	if err != nil {
		return err
	}

	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	v, ok := value.(runtime.ValueI64)
	if !ok {
		return fmt.Errorf("invalid value(%T): %w", value, runtime.ErrInvalidValue)
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [1]byte
	if _, err := binary.Encode(buf[:], endian, int8(v)); err != nil {
		return fmt.Errorf("failed to encode value: %w", err)
	}

	if n, err := r.WriteMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to write memory(%d): %w", n, err)
	}

	return nil
}

type I64Store16 struct {
	Align  uint32
	Offset uint32
}

func (i *I64Store16) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Store16
}

func (i *I64Store16) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return err
	}
	i.Offset, err = leb128.Uint32(r)
	return err
}

func (i *I64Store16) Execute(r runtime.Runtime, f *runtime.Frame) error {
	value, err := r.PopStack()
	if err != nil {
		return err
	}

	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	v, ok := value.(runtime.ValueI64)
	if !ok {
		return fmt.Errorf("invalid value(%T): %w", value, runtime.ErrInvalidValue)
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [2]byte
	if _, err := binary.Encode(buf[:], endian, int16(v)); err != nil {
		return fmt.Errorf("failed to encode value: %w", err)
	}

	if n, err := r.WriteMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to write memory(%d): %w", n, err)
	}

	return nil
}

type I64Store32 struct {
	Align  uint32
	Offset uint32
}

func (i *I64Store32) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Store32
}

func (i *I64Store32) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return err
	}
	i.Offset, err = leb128.Uint32(r)
	return err
}

func (i *I64Store32) Execute(r runtime.Runtime, f *runtime.Frame) error {
	value, err := r.PopStack()
	if err != nil {
		return err
	}

	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	v, ok := value.(runtime.ValueI64)
	if !ok {
		return fmt.Errorf("invalid value(%T): %w", value, runtime.ErrInvalidValue)
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [4]byte
	if _, err := binary.Encode(buf[:], endian, int32(v)); err != nil {
		return fmt.Errorf("failed to encode value: %w", err)
	}

	if n, err := r.WriteMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to write memory(%d): %w", n, err)
	}

	return nil
}

type F32Store struct {
	Align  uint32
	Offset uint32
}

func (i *F32Store) Opcode() opcode.Opcode {
	return opcode.OpcodeF32Store
}

func (i *F32Store) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return err
	}
	i.Offset, err = leb128.Uint32(r)
	return err
}

func (i *F32Store) Execute(r runtime.Runtime, f *runtime.Frame) error {
	value, err := r.PopStack()
	if err != nil {
		return err
	}

	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	v, ok := value.(runtime.ValueF32)
	if !ok {
		return fmt.Errorf("invalid value(%T): %w", value, runtime.ErrInvalidValue)
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [4]byte
	if _, err := binary.Encode(buf[:], endian, v); err != nil {
		return fmt.Errorf("failed to encode value: %w", err)
	}

	if n, err := r.WriteMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to write memory(%d): %w", n, err)
	}

	return nil
}

type F64Store struct {
	Align  uint32
	Offset uint32
}

func (i *F64Store) Opcode() opcode.Opcode {
	return opcode.OpcodeF64Store
}

func (i *F64Store) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return err
	}
	i.Offset, err = leb128.Uint32(r)
	return err
}

func (i *F64Store) Execute(r runtime.Runtime, f *runtime.Frame) error {
	value, err := r.PopStack()
	if err != nil {
		return err
	}

	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	v, ok := value.(runtime.ValueF64)
	if !ok {
		return fmt.Errorf("invalid value(%T): %w", value, runtime.ErrInvalidValue)
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [8]byte
	if _, err := binary.Encode(buf[:], endian, v); err != nil {
		return fmt.Errorf("failed to encode value: %w", err)
	}

	if n, err := r.WriteMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to write memory(%d): %w", n, err)
	}

	return nil
}
