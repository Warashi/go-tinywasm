package binary

import (
	"fmt"
	"io"

	"github.com/Warashi/go-tinywasm/leb128"
)

type Instruction interface {
	Opcode() Opcode
	ReadOperandsFrom(r io.Reader) error
}

func FromOpcode(op Opcode) (Instruction, error) {
	switch op {
	case OpcodeIf:
		return new(InstructionIf), nil
	case OpcodeEnd:
		return new(InstructionEnd), nil
	case OpcodeReturn:
		return new(InstructionReturn), nil
	case OpcodeCall:
		return new(InstructionCall), nil
	case OpcodeLocalGet:
		return new(InstructionLocalGet), nil
	case OpcodeLocalSet:
		return new(InstructionLocalSet), nil
	case OpcodeI32Store:
		return new(InstructionI32Store), nil
	case OpcodeI32Const:
		return new(InstructionI32Const), nil
	case OpcodeI32LtS:
		return new(InstructionI32LtS), nil
	case OpcodeI32Add:
		return new(InstructionI32Add), nil
	case OpcodeI32Sub:
		return new(InstructionI32Sub), nil
	default:
		return nil, fmt.Errorf("unknown opcode: %x", op)
	}
}

type InstructionIf struct {
	block Block
}

func (*InstructionIf) Opcode() Opcode { return OpcodeIf }
func (i *InstructionIf) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.block, err = decodeBlock(r)
	return err
}
func (i *InstructionIf) Block() Block { return i.block }

type InstructionEnd struct{}

func (*InstructionEnd) Opcode() Opcode                   { return OpcodeEnd }
func (*InstructionEnd) ReadOperandsFrom(io.Reader) error { return nil }

type InstructionReturn struct{}

func (*InstructionReturn) Opcode() Opcode                   { return OpcodeReturn }
func (*InstructionReturn) ReadOperandsFrom(io.Reader) error { return nil }

type InstructionCall struct{ index uint32 }

func (*InstructionCall) Opcode() Opcode { return OpcodeCall }
func (i *InstructionCall) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.index, err = leb128.Uint32(r)
	return err
}
func (i *InstructionCall) Index() uint32 { return i.index }

type InstructionLocalGet struct{ index uint32 }

func (*InstructionLocalGet) Opcode() Opcode { return OpcodeLocalGet }
func (i *InstructionLocalGet) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.index, err = leb128.Uint32(r)
	return err
}
func (i *InstructionLocalGet) Index() uint32 { return i.index }

type InstructionLocalSet struct{ index uint32 }

func (*InstructionLocalSet) Opcode() Opcode { return OpcodeLocalSet }
func (i *InstructionLocalSet) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.index, err = leb128.Uint32(r)
	return err
}
func (i *InstructionLocalSet) Index() uint32 { return i.index }

type InstructionI32Store struct {
	align  uint32
	offset uint32
}

func (*InstructionI32Store) Opcode() Opcode { return OpcodeI32Store }
func (i *InstructionI32Store) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.align, err = leb128.Uint32(r)
	if err != nil {
		return err
	}
	i.offset, err = leb128.Uint32(r)
	return err
}
func (i *InstructionI32Store) Align() uint32  { return i.align }
func (i *InstructionI32Store) Offset() uint32 { return i.offset }

type InstructionI32Const struct{ value int32 }

func (*InstructionI32Const) Opcode() Opcode { return OpcodeI32Const }
func (i *InstructionI32Const) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.value, err = leb128.Int32(r)
	return err
}
func (i *InstructionI32Const) Value() int32 { return i.value }

type InstructionI32LtS struct{}

func (*InstructionI32LtS) Opcode() Opcode                   { return OpcodeI32LtS }
func (*InstructionI32LtS) ReadOperandsFrom(io.Reader) error { return nil }

type InstructionI32Add struct{}

func (*InstructionI32Add) Opcode() Opcode                   { return OpcodeI32Add }
func (*InstructionI32Add) ReadOperandsFrom(io.Reader) error { return nil }

type InstructionI32Sub struct{}

func (*InstructionI32Sub) Opcode() Opcode                   { return OpcodeI32Sub }
func (*InstructionI32Sub) ReadOperandsFrom(io.Reader) error { return nil }
