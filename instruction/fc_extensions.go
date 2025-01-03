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

	switch b := opcode.OpcodeFC(b); b {
	case opcode.OpcodeFCI32TruncSatF32S:
		i.FC = new(FCI32TruncSatF32S)
	case opcode.OpcodeFCI32TruncSatF32U:
		i.FC = new(FCI32TruncSatF32U)
	case opcode.OpcodeFCI32TruncSatF64S:
		i.FC = new(FCI32TruncSatF64S)
	case opcode.OpcodeFCI32TruncSatF64U:
		i.FC = new(FCI32TruncSatF64U)
	case opcode.OpcodeFCI64TruncSatF32S:
		i.FC = new(FCI64TruncSatF32S)
	case opcode.OpcodeFCI64TruncSatF32U:
		i.FC = new(FCI64TruncSatF32U)
	case opcode.OpcodeFCI64TruncSatF64S:
		i.FC = new(FCI64TruncSatF64S)
	case opcode.OpcodeFCI64TruncSatF64U:
		i.FC = new(FCI64TruncSatF64U)
	default:
		return fmt.Errorf("unknown FC opcode: %v", b)
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

	fv := float64(f32.Float32())

	if math.IsNaN(fv) {
		r.PushStack(runtime.ValueI32(0))
	} else if fv < math.MinInt32 {
		r.PushStack(runtime.ValueI32(math.MinInt32))
	} else if fv > math.MaxInt32 {
		r.PushStack(runtime.ValueI32(math.MaxInt32))
	} else {
		r.PushStack(runtime.ValueI32(int32(fv)))
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

	fv := float64(f32.Float32())

	if math.IsNaN(float64(fv)) {
		r.PushStack(runtime.ValueI32(0))
	} else if fv < 0 {
		r.PushStack(runtime.ValueI32(0))
	} else if fv > math.MaxUint32 {
		r.PushStack(runtime.ValueI32(math.MaxInt32))
	} else {
		r.PushStack(runtime.ValueI32(uint32(fv)))
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

	fv := f64.Float64()

	if math.IsNaN(fv) {
		r.PushStack(runtime.ValueI32(0))
	} else if fv < math.MinInt32 {
		r.PushStack(runtime.ValueI32(math.MinInt32))
	} else if fv > math.MaxInt32 {
		r.PushStack(runtime.ValueI32(math.MaxInt32))
	} else {
		r.PushStack(runtime.ValueI32(int32(fv)))
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

	fv := f64.Float64()

	if math.IsNaN(fv) {
		r.PushStack(runtime.ValueI32(0))
	} else if fv < 0 {
		r.PushStack(runtime.ValueI32(0))
	} else if fv > math.MaxUint32 {
		r.PushStack(runtime.ValueI32(math.MaxInt32))
	} else {
		r.PushStack(runtime.ValueI32(uint32(fv)))
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

	fv := float64(f32.Float32())

	if math.IsNaN(fv) {
		r.PushStack(runtime.ValueI64(0))
	} else if fv < math.MinInt64 {
		r.PushStack(runtime.ValueI64(math.MinInt64))
	} else if fv > math.MaxInt64 {
		r.PushStack(runtime.ValueI64(math.MaxInt64))
	} else {
		r.PushStack(runtime.ValueI64(int64(fv)))
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

	fv := float64(f32.Float32())

	if math.IsNaN(fv) {
		r.PushStack(runtime.ValueI64(0))
	} else if fv < 0 {
		r.PushStack(runtime.ValueI64(0))
	} else if fv > math.MaxUint64 {
		r.PushStack(runtime.ValueI64(math.MaxInt64))
	} else {
		r.PushStack(runtime.ValueI64(uint64(fv)))
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

	fv := f64.Float64()

	if math.IsNaN(fv) {
		r.PushStack(runtime.ValueI64(0))
	} else if fv < math.MinInt64 {
		r.PushStack(runtime.ValueI64(math.MinInt64))
	} else if fv > math.MaxInt64 {
		r.PushStack(runtime.ValueI64(math.MaxInt64))
	} else {
		r.PushStack(runtime.ValueI64(int64(fv)))
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

	fv := f64.Float64()

	if math.IsNaN(fv) {
		r.PushStack(runtime.ValueI64(0))
	} else if fv < 0 {
		r.PushStack(runtime.ValueI64(0))
	} else if fv > math.MaxUint64 {
		r.PushStack(runtime.ValueI64(math.MaxInt64))
	} else {
		r.PushStack(runtime.ValueI64(uint64(fv)))
	}

	return nil
}
