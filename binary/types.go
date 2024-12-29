package binary

type FuncType struct {
	params  []ValueType
	results []ValueType
}

type ValueType byte

const (
	ValueTypeI32 ValueType = 0x7f
	ValueTypeI64 ValueType = 0x7e
)
