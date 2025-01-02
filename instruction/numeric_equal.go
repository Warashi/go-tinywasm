package instruction

import (
	"fmt"
	"io"

	"github.com/Warashi/wasmium/opcode"
	"github.com/Warashi/wasmium/types/runtime"
)

type I32Eqz struct{}

func (i *I32Eqz) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Eqz
}

func (i *I32Eqz) ReadOperandsFrom(r io.Reader) error {
	return nil
}

func (i *I32Eqz) Execute(r runtime.Runtime, f *runtime.Frame) error {
	a, err := r.PopStack()
	if err != nil {
		return err
	}

	v, ok := a.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid value(%T): %w", a, runtime.ErrInvalidValue)
	}

	if v == 0 {
		r.PushStack(runtime.ValueI32(1))
	} else {
		r.PushStack(runtime.ValueI32(0))
	}

	return nil
}

type I32Eq struct{}

func (i *I32Eq) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Eq
}

func (i *I32Eq) ReadOperandsFrom(r io.Reader) error {
	return nil
}

func (i *I32Eq) Execute(r runtime.Runtime, f *runtime.Frame) error {
	b, err := r.PopStack()
	if err != nil {
		return err
	}

	a, err := r.PopStack()
	if err != nil {
		return err
	}

	va, ok := a.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid value(%T): %w", a, runtime.ErrInvalidValue)
	}

	vb, ok := b.(runtime.ValueI32)
	if !ok {
		return fmt.Errorf("invalid value(%T): %w", b, runtime.ErrInvalidValue)
	}

	if va == vb {
		r.PushStack(runtime.ValueI32(1))
	} else {
		r.PushStack(runtime.ValueI32(0))
	}

	return nil
}

type I64Eqz struct{}

func (i *I64Eqz) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Eqz
}

func (i *I64Eqz) ReadOperandsFrom(r io.Reader) error {
	return nil
}

func (i *I64Eqz) Execute(r runtime.Runtime, f *runtime.Frame) error {
	a, err := r.PopStack()
	if err != nil {
		return err
	}

	v, ok := a.(runtime.ValueI64)
	if !ok {
		return fmt.Errorf("invalid value(%T): %w", a, runtime.ErrInvalidValue)
	}

	if v == 0 {
		r.PushStack(runtime.ValueI32(1))
	} else {
		r.PushStack(runtime.ValueI32(0))
	}

	return nil
}

type I64Eq struct{}

func (i *I64Eq) Opcode() opcode.Opcode {
	return opcode.OpcodeI64Eq
}

func (i *I64Eq) ReadOperandsFrom(r io.Reader) error {
	return nil
}

func (i *I64Eq) Execute(r runtime.Runtime, f *runtime.Frame) error {
	b, err := r.PopStack()
	if err != nil {
		return err
	}

	a, err := r.PopStack()
	if err != nil {
		return err
	}

	va, ok := a.(runtime.ValueI64)
	if !ok {
		return fmt.Errorf("invalid value(%T): %w", a, runtime.ErrInvalidValue)
	}

	vb, ok := b.(runtime.ValueI64)
	if !ok {
		return fmt.Errorf("invalid value(%T): %w", b, runtime.ErrInvalidValue)
	}

	if va == vb {
		r.PushStack(runtime.ValueI32(1))
	} else {
		r.PushStack(runtime.ValueI32(0))
	}

	return nil
}

type F32Eq struct{}

func (i *F32Eq) Opcode() opcode.Opcode {
	return opcode.OpcodeF32Eq
}

func (i *F32Eq) ReadOperandsFrom(r io.Reader) error {
	return nil
}

func (i *F32Eq) Execute(r runtime.Runtime, f *runtime.Frame) error {
	b, err := r.PopStack()
	if err != nil {
		return err
	}

	a, err := r.PopStack()
	if err != nil {
		return err
	}

	va, ok := a.(runtime.ValueF32)
	if !ok {
		return fmt.Errorf("invalid value(%T): %w", a, runtime.ErrInvalidValue)
	}

	vb, ok := b.(runtime.ValueF32)
	if !ok {
		return fmt.Errorf("invalid value(%T): %w", b, runtime.ErrInvalidValue)
	}

	if va == vb {
		r.PushStack(runtime.ValueI32(1))
	} else {
		r.PushStack(runtime.ValueI32(0))
	}

	return nil
}

type F64Eq struct{}

func (i *F64Eq) Opcode() opcode.Opcode {
	return opcode.OpcodeF64Eq
}

func (i *F64Eq) ReadOperandsFrom(r io.Reader) error {
	return nil
}

func (i *F64Eq) Execute(r runtime.Runtime, f *runtime.Frame) error {
	b, err := r.PopStack()
	if err != nil {
		return err
	}

	a, err := r.PopStack()
	if err != nil {
		return err
	}

	va, ok := a.(runtime.ValueF64)
	if !ok {
		return fmt.Errorf("invalid value(%T): %w", a, runtime.ErrInvalidValue)
	}

	vb, ok := b.(runtime.ValueF64)
	if !ok {
		return fmt.Errorf("invalid value(%T): %w", b, runtime.ErrInvalidValue)
	}

	if va == vb {
		r.PushStack(runtime.ValueI32(1))
	} else {
		r.PushStack(runtime.ValueI32(0))
	}

	return nil
}
