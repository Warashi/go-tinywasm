package binary

type FuncType struct {
	params  []ValueType
	results []ValueType
}

func (f FuncType) Params() []ValueType {
	return f.params
}

func (f FuncType) Results() []ValueType {
	return f.results
}

type ValueType byte

const (
	ValueTypeI32 ValueType = 0x7f
	ValueTypeI64 ValueType = 0x7e
)

type FunctionLocal struct {
	typeCount uint32
	valueType ValueType
}

func (f FunctionLocal) TypeCount() uint32    { return f.typeCount }
func (f FunctionLocal) ValueType() ValueType { return f.valueType }

type ExportDesc interface {
	isExportDesc()
}

type ExportDescFunc struct {
	index uint32
}

func (e ExportDescFunc) isExportDesc() {}
func (e ExportDescFunc) Index() uint32 { return e.index }

type Export struct {
	name string
	desc ExportDesc
}

func (e Export) Name() string     { return e.name }
func (e Export) Desc() ExportDesc { return e.desc }
