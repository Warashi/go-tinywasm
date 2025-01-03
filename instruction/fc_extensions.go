package instruction

import (
	"fmt"
	"io"
	"math"

	"github.com/Warashi/wasmium/leb128"
	"github.com/Warashi/wasmium/opcode"
	"github.com/Warashi/wasmium/types/runtime"
)

type FC interface {
	Opcode() opcode.OpcodeFC
	ReadOperandsFrom(r io.Reader) error
	Execute(r runtime.Runtime, f *runtime.Frame) error
}

type FCPrefix struct {
	FC
}

func (*FCPrefix) Opcode() opcode.Opcode { return opcode.OpcodeFCPrefix }

func (i *FCPrefix) ReadOperandsFrom(r io.Reader) error {
	b, err := leb128.Uint32(r)
	if err != nil {
		return fmt.Errorf("failed to read byte: %w", err)
	}

	switch b {
	case 0x00:
		i.FC = new(FCI32TruncSatF32S)
	case 0x01:
		i.FC = new(FCI32TruncSatF32U)
	case 0x02:
		i.FC = new(FCI32TruncSatF64S)
	case 0x03:
		i.FC = new(FCI32TruncSatF64U)
	case 0x04:
		i.FC = new(FCI64TruncSatF32S)
	case 0x05:
		i.FC = new(FCI64TruncSatF32U)
	case 0x06:
		i.FC = new(FCI64TruncSatF64S)
	case 0x07:
		i.FC = new(FCI64TruncSatF64U)
	default:
		return fmt.Errorf("invalid FC opcode: %v", b)
	}

	return i.FC.ReadOperandsFrom(r)
}

type FCI32TruncSatF32S struct{}

func (*FCI32TruncSatF32S) Opcode() opcode.OpcodeFC { return opcode.OpcodeFCI32TruncSatF32S }

func (*FCI32TruncSatF32S) ReadOperandsFrom(r io.Reader) error { return nil }

func (*FCI32TruncSatF32S) Execute(r runtime.Runtime, f *runtime.Frame) error {
	v, err := r.PopStack()
	if err != nil {
		return fmt.Errorf("failed to pop stack: %w", err)
	}

	f32, ok := v.(runtime.ValueF32)
	if !ok {
		return runtime.ErrInvalidValue
	}

	if math.IsNaN(float64(f32)) {
		r.PushStack(runtime.ValueI32(0))
	} else if f32 < math.MinInt32 {
		r.PushStack(runtime.ValueI32(math.MinInt32))
	} else if f32 > math.MaxInt32 {
		r.PushStack(runtime.ValueI32(math.MaxInt32))
	} else {
		r.PushStack(runtime.ValueI32(int32(f32)))
	}

	return nil
}

type FCI32TruncSatF32U struct{}

func (*FCI32TruncSatF32U) Opcode() opcode.OpcodeFC { return opcode.OpcodeFCI32TruncSatF32U }

func (*FCI32TruncSatF32U) ReadOperandsFrom(r io.Reader) error { return nil }

func (*FCI32TruncSatF32U) Execute(r runtime.Runtime, f *runtime.Frame) error {
	v, err := r.PopStack()
	if err != nil {
		return fmt.Errorf("failed to pop stack: %w", err)
	}

	f32, ok := v.(runtime.ValueF32)
	if !ok {
		return runtime.ErrInvalidValue
	}

	if math.IsNaN(float64(f32)) {
		r.PushStack(runtime.ValueI32(0))
	} else if f32 < 0 {
		r.PushStack(runtime.ValueI32(0))
	} else if f32 > math.MaxUint32 {
		r.PushStack(runtime.ValueI32(math.MaxInt32))
	} else {
		r.PushStack(runtime.ValueI32(uint32(f32)))
	}

	return nil
}

type FCI32TruncSatF64S struct{}

func (*FCI32TruncSatF64S) Opcode() opcode.OpcodeFC { return opcode.OpcodeFCI32TruncSatF64S }

func (*FCI32TruncSatF64S) ReadOperandsFrom(r io.Reader) error { return nil }

func (*FCI32TruncSatF64S) Execute(r runtime.Runtime, f *runtime.Frame) error {
	v, err := r.PopStack()
	if err != nil {
		return fmt.Errorf("failed to pop stack: %w", err)
	}

	f64, ok := v.(runtime.ValueF64)
	if !ok {
		return runtime.ErrInvalidValue
	}

	if math.IsNaN(float64(f64)) {
		r.PushStack(runtime.ValueI32(0))
	} else if f64 < math.MinInt32 {
		r.PushStack(runtime.ValueI32(math.MinInt32))
	} else if f64 > math.MaxInt32 {
		r.PushStack(runtime.ValueI32(math.MaxInt32))
	} else {
		r.PushStack(runtime.ValueI32(int32(f64)))
	}

	return nil
}

type FCI32TruncSatF64U struct{}

func (*FCI32TruncSatF64U) Opcode() opcode.OpcodeFC { return opcode.OpcodeFCI32TruncSatF64U }

func (*FCI32TruncSatF64U) ReadOperandsFrom(r io.Reader) error { return nil }

func (*FCI32TruncSatF64U) Execute(r runtime.Runtime, f *runtime.Frame) error {
	v, err := r.PopStack()
	if err != nil {
		return fmt.Errorf("failed to pop stack: %w", err)
	}

	f64, ok := v.(runtime.ValueF64)
	if !ok {
		return runtime.ErrInvalidValue
	}

	if math.IsNaN(float64(f64)) {
		r.PushStack(runtime.ValueI32(0))
	} else if f64 < 0 {
		r.PushStack(runtime.ValueI32(0))
	} else if f64 > math.MaxUint32 {
		r.PushStack(runtime.ValueI32(math.MaxInt32))
	} else {
		r.PushStack(runtime.ValueI32(uint32(f64)))
	}

	return nil
}

type FCI64TruncSatF32S struct{}

func (*FCI64TruncSatF32S) Opcode() opcode.OpcodeFC { return opcode.OpcodeFCI64TruncSatF32S }

func (*FCI64TruncSatF32S) ReadOperandsFrom(r io.Reader) error { return nil }

func (*FCI64TruncSatF32S) Execute(r runtime.Runtime, f *runtime.Frame) error {
	v, err := r.PopStack()
	if err != nil {
		return fmt.Errorf("failed to pop stack: %w", err)
	}

	f32, ok := v.(runtime.ValueF32)
	if !ok {
		return runtime.ErrInvalidValue
	}

	if math.IsNaN(float64(f32)) {
		r.PushStack(runtime.ValueI64(0))
	} else if f32 < math.MinInt64 {
		r.PushStack(runtime.ValueI64(math.MinInt64))
	} else if f32 > math.MaxInt64 {
		r.PushStack(runtime.ValueI64(math.MaxInt64))
	} else {
		r.PushStack(runtime.ValueI64(int64(f32)))
	}

	return nil
}

type FCI64TruncSatF32U struct{}

func (*FCI64TruncSatF32U) Opcode() opcode.OpcodeFC { return opcode.OpcodeFCI64TruncSatF32U }

func (*FCI64TruncSatF32U) ReadOperandsFrom(r io.Reader) error { return nil }

func (*FCI64TruncSatF32U) Execute(r runtime.Runtime, f *runtime.Frame) error {
	v, err := r.PopStack()
	if err != nil {
		return fmt.Errorf("failed to pop stack: %w", err)
	}

	f32, ok := v.(runtime.ValueF32)
	if !ok {
		return runtime.ErrInvalidValue
	}

	if math.IsNaN(float64(f32)) {
		r.PushStack(runtime.ValueI64(0))
	} else if f32 < 0 {
		r.PushStack(runtime.ValueI64(0))
	} else if f32 > math.MaxUint64 {
		r.PushStack(runtime.ValueI64(math.MaxInt64))
	} else {
		r.PushStack(runtime.ValueI64(uint64(f32)))
	}

	return nil
}

type FCI64TruncSatF64S struct{}

func (*FCI64TruncSatF64S) Opcode() opcode.OpcodeFC { return opcode.OpcodeFCI64TruncSatF64S }

func (*FCI64TruncSatF64S) ReadOperandsFrom(r io.Reader) error { return nil }

func (*FCI64TruncSatF64S) Execute(r runtime.Runtime, f *runtime.Frame) error {
	v, err := r.PopStack()
	if err != nil {
		return fmt.Errorf("failed to pop stack: %w", err)
	}

	f64, ok := v.(runtime.ValueF64)
	if !ok {
		return runtime.ErrInvalidValue
	}

	if math.IsNaN(float64(f64)) {
		r.PushStack(runtime.ValueI64(0))
	} else if f64 < math.MinInt64 {
		r.PushStack(runtime.ValueI64(math.MinInt64))
	} else if f64 > math.MaxInt64 {
		r.PushStack(runtime.ValueI64(math.MaxInt64))
	} else {
		r.PushStack(runtime.ValueI64(int64(f64)))
	}

	return nil
}

type FCI64TruncSatF64U struct{}

func (*FCI64TruncSatF64U) Opcode() opcode.OpcodeFC { return opcode.OpcodeFCI64TruncSatF64U }

func (*FCI64TruncSatF64U) ReadOperandsFrom(r io.Reader) error { return nil }

func (*FCI64TruncSatF64U) Execute(r runtime.Runtime, f *runtime.Frame) error {
	v, err := r.PopStack()
	if err != nil {
		return fmt.Errorf("failed to pop stack: %w", err)
	}

	f64, ok := v.(runtime.ValueF64)
	if !ok {
		return runtime.ErrInvalidValue
	}

	if math.IsNaN(float64(f64)) {
		r.PushStack(runtime.ValueI64(0))
	} else if f64 < 0 {
		r.PushStack(runtime.ValueI64(0))
	} else if f64 > math.MaxUint64 {
		r.PushStack(runtime.ValueI64(math.MaxInt64))
	} else {
		r.PushStack(runtime.ValueI64(uint64(f64)))
	}

	return nil
}
