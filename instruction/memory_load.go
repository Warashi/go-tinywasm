package instruction

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/Warashi/wasmium/leb128"
	"github.com/Warashi/wasmium/opcode"
	"github.com/Warashi/wasmium/types/runtime"
)

type I32Load struct {
	Align  uint32
	Offset uint32
}

func (i *I32Load) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Load
}

func (i *I32Load) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read align: %w", err)
	}

	i.Offset, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read offset: %w", err)
	}

	return nil
}

func (i *I32Load) Execute(r runtime.Runtime, f *runtime.Frame) error {
	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [4]byte
	if n, err := r.ReadMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to read memory(%d): %w", n, err)
	}

	var result int32
	if _, err := binary.Decode(buf[:], binary.LittleEndian, &result); err != nil {
		return fmt.Errorf("failed to decode value: %w", err)
	}

	r.PushStack(runtime.ValueI32(result))

	return nil
}

type I64Load struct {
	Align  uint32
	Offset uint32
}

func (i *I64Load) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Load
}

func (i *I64Load) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read align: %w", err)
	}

	i.Offset, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read offset: %w", err)
	}

	return nil
}

func (i *I64Load) Execute(r runtime.Runtime, f *runtime.Frame) error {
	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [8]byte
	if n, err := r.ReadMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to read memory(%d): %w", n, err)
	}

	var result int64
	if _, err := binary.Decode(buf[:], binary.LittleEndian, &result); err != nil {
		return fmt.Errorf("failed to decode value: %w", err)
	}

	r.PushStack(runtime.ValueI64(result))

	return nil
}

type I32Load8S struct {
	Align  uint32
	Offset uint32
}

func (i *I32Load8S) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Load8S
}

func (i *I32Load8S) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read align: %w", err)
	}

	i.Offset, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read offset: %w", err)
	}

	return nil
}

func (i *I32Load8S) Execute(r runtime.Runtime, f *runtime.Frame) error {
	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [1]byte
	if n, err := r.ReadMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to read memory(%d): %w", n, err)
	}

	var result int8
	if _, err := binary.Decode(buf[:], binary.LittleEndian, &result); err != nil {
		return fmt.Errorf("failed to decode value: %w", err)
	}

	r.PushStack(runtime.ValueI32(result))

	return nil

}

type I32Load8U struct {
	Align  uint32
	Offset uint32
}

func (i *I32Load8U) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Load8U
}

func (i *I32Load8U) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read align: %w", err)
	}

	i.Offset, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read offset: %w", err)
	}

	return nil
}

func (i *I32Load8U) Execute(r runtime.Runtime, f *runtime.Frame) error {
	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [1]byte
	if n, err := r.ReadMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to read memory(%d): %w", n, err)
	}

	var result uint8
	if _, err := binary.Decode(buf[:], binary.LittleEndian, &result); err != nil {
		return fmt.Errorf("failed to decode value: %w", err)
	}

	r.PushStack(runtime.ValueI32(result))

	return nil
}

type I32Load16S struct {
	Align  uint32
	Offset uint32
}

func (i *I32Load16S) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Load16S
}

func (i *I32Load16S) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read align: %w", err)
	}

	i.Offset, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read offset: %w", err)
	}

	return nil
}

func (i *I32Load16S) Execute(r runtime.Runtime, f *runtime.Frame) error {
	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [2]byte
	if n, err := r.ReadMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to read memory(%d): %w", n, err)
	}

	var result int16
	if _, err := binary.Decode(buf[:], binary.LittleEndian, &result); err != nil {
		return fmt.Errorf("failed to decode value: %w", err)
	}

	r.PushStack(runtime.ValueI32(result))

	return nil

}

type I32Load16U struct {
	Align  uint32
	Offset uint32
}

func (i *I32Load16U) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Load16U
}

func (i *I32Load16U) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read align: %w", err)
	}

	i.Offset, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read offset: %w", err)
	}

	return nil
}

func (i *I32Load16U) Execute(r runtime.Runtime, f *runtime.Frame) error {
	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [2]byte
	if n, err := r.ReadMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to read memory(%d): %w", n, err)
	}

	var result uint16
	if _, err := binary.Decode(buf[:], binary.LittleEndian, &result); err != nil {
		return fmt.Errorf("failed to decode value: %w", err)
	}

	r.PushStack(runtime.ValueI32(result))

	return nil
}

type I64Load8S struct {
	Align  uint32
	Offset uint32
}

func (i *I64Load8S) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Load8S
}

func (i *I64Load8S) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read align: %w", err)
	}

	i.Offset, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read offset: %w", err)
	}

	return nil
}

func (i *I64Load8S) Execute(r runtime.Runtime, f *runtime.Frame) error {
	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [1]byte
	if n, err := r.ReadMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to read memory(%d): %w", n, err)
	}

	var result int8
	if _, err := binary.Decode(buf[:], binary.LittleEndian, &result); err != nil {
		return fmt.Errorf("failed to decode value: %w", err)
	}

	r.PushStack(runtime.ValueI64(result))

	return nil
}

type I64Load8U struct {
	Align  uint32
	Offset uint32
}

func (i *I64Load8U) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Load8U
}

func (i *I64Load8U) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read align: %w", err)
	}

	i.Offset, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read offset: %w", err)
	}

	return nil
}

func (i *I64Load8U) Execute(r runtime.Runtime, f *runtime.Frame) error {
	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [1]byte
	if n, err := r.ReadMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to read memory(%d): %w", n, err)
	}

	var result uint8
	if _, err := binary.Decode(buf[:], binary.LittleEndian, &result); err != nil {
		return fmt.Errorf("failed to decode value: %w", err)
	}

	r.PushStack(runtime.ValueI64(result))

	return nil
}

type I64Load16S struct {
	Align  uint32
	Offset uint32
}

func (i *I64Load16S) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Load16S
}

func (i *I64Load16S) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read align: %w", err)
	}

	i.Offset, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read offset: %w", err)
	}

	return nil
}

func (i *I64Load16S) Execute(r runtime.Runtime, f *runtime.Frame) error {
	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [2]byte
	if n, err := r.ReadMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to read memory(%d): %w", n, err)
	}

	var result int16
	if _, err := binary.Decode(buf[:], binary.LittleEndian, &result); err != nil {
		return fmt.Errorf("failed to decode value: %w", err)
	}

	r.PushStack(runtime.ValueI64(result))

	return nil
}

type I64Load16U struct {
	Align  uint32
	Offset uint32
}

func (i *I64Load16U) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Load16U
}

func (i *I64Load16U) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read align: %w", err)
	}

	i.Offset, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read offset: %w", err)
	}

	return nil
}

func (i *I64Load16U) Execute(r runtime.Runtime, f *runtime.Frame) error {
	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [2]byte
	if n, err := r.ReadMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to read memory(%d): %w", n, err)
	}

	var result uint16
	if _, err := binary.Decode(buf[:], binary.LittleEndian, &result); err != nil {
		return fmt.Errorf("failed to decode value: %w", err)
	}

	r.PushStack(runtime.ValueI64(result))

	return nil
}

type I64Load32U struct {
	Align  uint32
	Offset uint32
}

func (i *I64Load32U) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Load32U
}

func (i *I64Load32U) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read align: %w", err)
	}

	i.Offset, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read offset: %w", err)
	}

	return nil
}

func (i *I64Load32U) Execute(r runtime.Runtime, f *runtime.Frame) error {
	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [4]byte
	if n, err := r.ReadMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to read memory(%d): %w", n, err)
	}

	var result uint32
	if _, err := binary.Decode(buf[:], binary.LittleEndian, &result); err != nil {
		return fmt.Errorf("failed to decode value: %w", err)
	}

	r.PushStack(runtime.ValueI64(result))

	return nil
}

type I64Load32S struct {
	Align  uint32
	Offset uint32
}

func (i *I64Load32S) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Load32S
}

func (i *I64Load32S) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read align: %w", err)
	}

	i.Offset, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read offset: %w", err)
	}

	return nil
}

func (i *I64Load32S) Execute(r runtime.Runtime, f *runtime.Frame) error {
	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [4]byte
	if n, err := r.ReadMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to read memory(%d): %w", n, err)
	}

	var result int32
	if _, err := binary.Decode(buf[:], binary.LittleEndian, &result); err != nil {
		return fmt.Errorf("failed to decode value: %w", err)
	}

	r.PushStack(runtime.ValueI64(result))

	return nil
}

type F32Load struct {
	Align  uint32
	Offset uint32
}

func (i *F32Load) Opcode() opcode.Opcode {
	return opcode.OpcodeF32Load
}

func (i *F32Load) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read align: %w", err)
	}

	i.Offset, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read offset: %w", err)
	}

	return nil
}

func (i *F32Load) Execute(r runtime.Runtime, f *runtime.Frame) error {
	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [4]byte
	if n, err := r.ReadMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to read memory(%d): %w", n, err)
	}

	r.PushStack(runtime.ValueF32(buf))

	return nil
}

type F64Load struct {
	Align  uint32
	Offset uint32
}

func (i *F64Load) Opcode() opcode.Opcode {
	return opcode.OpcodeF64Load
}

func (i *F64Load) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.Align, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read align: %w", err)
	}

	i.Offset, err = leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read offset: %w", err)
	}

	return nil
}

func (i *F64Load) Execute(r runtime.Runtime, f *runtime.Frame) error {
	addr, err := r.PopStack()
	if err != nil {
		return err
	}

	a, ok := addr.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid addr(%T): %w", addr, runtime.ErrInvalidValue)
	}

	var buf [8]byte
	if n, err := r.ReadMemoryAt(0, buf[:], int64(uint32(a)+i.Offset)); err != nil || n != len(buf) {
		return fmt.Errorf("failed to read memory(%d): %w", n, err)
	}

	r.PushStack(runtime.ValueF64(buf))

	return nil
}
