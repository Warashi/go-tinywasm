package instruction

import (
	"fmt"
	"io"
)

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

type ImportDesc interface {
	isImportDesc()
}

type ImportDescFunc struct {
	index uint32
}

func (i ImportDescFunc) isImportDesc() {}
func (i ImportDescFunc) Index() uint32 { return i.index }

type Import struct {
	module string
	field  string
	desc   ImportDesc
}

func (i Import) Module() string   { return i.module }
func (i Import) Field() string    { return i.field }
func (i Import) Desc() ImportDesc { return i.desc }

type Limits struct {
	min uint32
	max uint32
}

func (l Limits) Min() uint32 { return l.min }
func (l Limits) Max() uint32 { return l.max }

type Memory struct {
	limits Limits
}

func (m Memory) Limits() Limits { return m.limits }

type Data struct {
	memoryIndex uint32
	offset      uint32
	init        []byte
}

func (d Data) MemoryIndex() uint32 { return d.memoryIndex }
func (d Data) Offset() uint32      { return d.offset }
func (d Data) Init() []byte        { return d.init }

type Block struct {
	blockType BlockType
}

func (b *Block) decode(r io.Reader) error {
	buf, err := readByte(r)
	if err != nil {
		return fmt.Errorf("failed to read block type: %w", err)
	}

	switch buf {
	case 0x40:
		*b = Block{blockType: BlockTypeVoid{}}
	default:
		*b = Block{blockType: BlockTypeValue{valueTypes: []ValueType{ValueType(buf)}}}
	}

	return nil
}

func (b Block) BlockType() BlockType { return b.blockType }

type BlockType interface {
	isBlockType()
	ResultCount() int
}

type BlockTypeVoid struct{}

func (b BlockTypeVoid) isBlockType()     {}
func (b BlockTypeVoid) ResultCount() int { return 0 }

type BlockTypeValue struct {
	valueTypes []ValueType
}

func (b BlockTypeValue) isBlockType()     {}
func (b BlockTypeValue) ResultCount() int { return len(b.valueTypes) }
