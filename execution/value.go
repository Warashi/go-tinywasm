package execution

import "fmt"

type ValueType byte

const (
	ValueTypeI32 ValueType = 0x7F
	ValueTypeI64 ValueType = 0x7E
	ValueTypeF32 ValueType = 0x7D
	ValueTypeF64 ValueType = 0x7C
)

func (t ValueType) String() string {
	switch t {
	case ValueTypeI32:
		return "i32"
	case ValueTypeI64:
		return "i64"
	case ValueTypeF32:
		return "f32"
	case ValueTypeF64:
		return "f64"
	default:
		return "unknown"
	}
}

func ValueFrom(v any) (Value, error) {
	switch v := v.(type) {
	case int32:
		return I32(v), nil
	case int64:
		return I64(v), nil
	}
	return nil, fmt.Errorf("unsupported type %T", v)
}

type Value interface {
	Type() ValueType
	Add(Value) (Value, error)
}

type I32 int32

func (I32) Type() ValueType { return ValueTypeI32 }

func (i I32) Add(v Value) (Value, error) {
	if v.Type() != ValueTypeI32 {
		return nil, fmt.Errorf("type mismatch: expected I32, got %v", v.Type())
	}
	return i + v.(I32), nil
}

type I64 int64

func (I64) Type() ValueType { return ValueTypeI64 }

func (i I64) Add(v Value) (Value, error) {
	if v.Type() != ValueTypeI64 {
		return nil, fmt.Errorf("type mismatch: expected I64, got %v", v.Type())
	}
	return i + v.(I64), nil
}
