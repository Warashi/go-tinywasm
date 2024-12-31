package binary

import (
	"io"

	"github.com/Warashi/wasmium/opcode"
)

type Instruction interface {
	Opcode() opcode.Opcode
	ReadOperandsFrom(r io.Reader) error
}

type Function struct {
	Locals []FunctionLocal
	Code   []Instruction
}

type FuncType struct {
	Params  []ValueType
	Results []ValueType
}

type ValueType byte

const (
	ValueTypeI32 ValueType = 0x7f
	ValueTypeI64 ValueType = 0x7e
)

type FunctionLocal struct {
	TypeCount uint32
	ValueType ValueType
}

type ImportDesc interface {
	isImportDesc()
}

type ImportDescFunc struct {
	Index uint32
}

func (i ImportDescFunc) isImportDesc() {}

type Import struct {
	Module string
	Field  string
	Desc   ImportDesc
}

type Limits struct {
	Min uint32
	Max uint32
}

type Memory struct {
	Limits Limits
}

type Data struct {
	MemoryIndex uint32
	Offset      uint32
	Init        []byte
}

type Block struct {
	BlockType BlockType
}

type BlockType interface {
	isBlockType()
	ResultCount() int
}

type BlockTypeVoid struct{}

func (b BlockTypeVoid) isBlockType()     {}
func (b BlockTypeVoid) ResultCount() int { return 0 }

type BlockTypeValue struct {
	ValueTypes []ValueType
}

func (b BlockTypeValue) isBlockType()     {}
func (b BlockTypeValue) ResultCount() int { return len(b.ValueTypes) }
