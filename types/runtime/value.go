package runtime

import (
	"encoding/binary"
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
	case bool:
		if v {
			return ValueI32(1), nil
		}
		return ValueI32(0), nil
	}
	return nil, fmt.Errorf("unsupported type %T", v)
}

type Value interface {
	isValue()
	Type() ValueType
	Int() int
	Bool() bool
}

type ValueI32 int32

func (ValueI32) isValue()        {}
func (ValueI32) Type() ValueType { return ValueTypeI32 }
func (v ValueI32) Int() int      { return int(v) }
func (v ValueI32) Bool() bool    { return v != 0 }

type ValueI64 int64

func (ValueI64) isValue()        {}
func (ValueI64) Type() ValueType { return ValueTypeI64 }
func (v ValueI64) Int() int      { return int(v) }
func (v ValueI64) Bool() bool    { return v != 0 }

type ValueF32 [4]byte

func (ValueF32) isValue()        {}
func (ValueF32) Type() ValueType { return ValueTypeF32 }
func (ValueF32) Int() int        { panic("int for f32 is not allowed") }
func (ValueF32) Bool() bool      { panic("bool for f32 is not allowed") }
func (v ValueF32) Float32() float32 {
	var f float32
	binary.Decode(v[:], binary.LittleEndian, &f)
	return f
}

type ValueF64 [8]byte

func (ValueF64) isValue()        {}
func (ValueF64) Type() ValueType { return ValueTypeF64 }
func (ValueF64) Int() int        { panic("int for f64 is not allowed") }
func (ValueF64) Bool() bool      { panic("bool for f64 is not allowed") }
func (v ValueF64) Float64() float64 {
	var f float64
	binary.Decode(v[:], binary.LittleEndian, &f)
	return f
}

type LabelKind int

const (
	LabelKindBlock LabelKind = iota
	LabelKindLoop
	LabelKindIf
)

type Label struct {
	kind           LabelKind
	start          int
	programCounter int
	stackPointer   int
	arity          int
}

func NewLabel(kind LabelKind, start, pc, sp, arity int) Label {
	return Label{
		kind:           kind,
		start:          start,
		programCounter: pc,
		stackPointer:   sp,
		arity:          arity,
	}
}

func (l Label) Kind() LabelKind {
	return l.kind
}

func (l Label) Start() int {
	return l.start
}

func (l Label) ProgramCounter() int {
	return l.programCounter
}

func (l Label) StackPointer() int {
	return l.stackPointer
}

func (l Label) Arity() int {
	return l.arity
}

func writeValue(buf []byte, v Value) (int, error) {
	switch v := v.(type) {
	case ValueI32:
		return binary.Encode(buf, binary.LittleEndian, int32(v))
	case ValueI64:
		return binary.Encode(buf, binary.LittleEndian, int64(v))
	}
	return 0, fmt.Errorf("unsupported type %T", v)
}

func readValue[T Value](buf []byte) (int, T, error) {
	var v T
	n, err := binary.Decode(buf, binary.LittleEndian, &v)
	return n, v, err
}
