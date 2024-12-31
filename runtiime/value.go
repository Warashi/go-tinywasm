package runtime

import (
	"fmt"
)

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
		return ValueI32(v), nil
	case int64:
		return ValueI64(v), nil
	case float32:
		return ValueF32(v), nil
	case float64:
		return ValueF64(v), nil
	case bool:
		if v {
			return ValueI32(1), nil
		}
		return ValueI32(0), nil
	}
	return nil, fmt.Errorf("unsupported type %T", v)
}

type Value interface {
	Type() ValueType
}

type ValueI32 int32

func (ValueI32) Type() ValueType { return ValueTypeI32 }

type ValueI64 int64

func (ValueI64) Type() ValueType { return ValueTypeI64 }

type ValueF32 float32

func (ValueF32) Type() ValueType { return ValueTypeF32 }

type ValueF64 float64

func (ValueF64) Type() ValueType { return ValueTypeF64 }
