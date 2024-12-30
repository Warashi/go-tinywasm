package binary

import (
	"io"

	"github.com/Warashi/go-tinywasm/leb128"
)

type Instruction interface {
	Opcode() Opcode
	ReadFrom(r io.Reader) error
}

type InstructionEnd struct{}

func (InstructionEnd) Opcode() Opcode           { return OpcodeEnd }
func (InstructionEnd) ReadFrom(io.Reader) error { return nil }

type InstructionCall struct{ index uint32 }

func (InstructionCall) Opcode() Opcode { return OpcodeCall }
func (i *InstructionCall) ReadFrom(r io.Reader) error {
	var err error
	i.index, err = leb128.Uint32(r)
	return err
}
func (i InstructionCall) Index() uint32 { return i.index }

type InstructionLocalGet struct{ index uint32 }

func (InstructionLocalGet) Opcode() Opcode { return OpcodeLocalGet }
func (i *InstructionLocalGet) ReadFrom(r io.Reader) error {
	var err error
	i.index, err = leb128.Uint32(r)
	return err
}
func (i InstructionLocalGet) Index() uint32 { return i.index }

type InstructionLocalSet struct{ index uint32 }

func (InstructionLocalSet) Opcode() Opcode { return OpcodeLocalSet }
func (i *InstructionLocalSet) ReadFrom(r io.Reader) error {
	var err error
	i.index, err = leb128.Uint32(r)
	return err
}
func (i InstructionLocalSet) Index() uint32 { return i.index }

type InstructionI32Store struct {
	align  uint32
	offset uint32
}

func (InstructionI32Store) Opcode() Opcode { return OpcodeI32Store }
func (i *InstructionI32Store) ReadFrom(r io.Reader) error {
	var err error
	i.align, err = leb128.Uint32(r)
	if err != nil {
		return err
	}
	i.offset, err = leb128.Uint32(r)
	return err
}
func (i InstructionI32Store) Align() uint32  { return i.align }
func (i InstructionI32Store) Offset() uint32 { return i.offset }

type InstructionI32Const struct{ value int32 }

func (InstructionI32Const) Opcode() Opcode { return OpcodeI32Const }
func (i *InstructionI32Const) ReadFrom(r io.Reader) error {
	var err error
	i.value, err = leb128.Int32(r)
	return err
}
func (i InstructionI32Const) Value() int32 { return i.value }

type InstructionI32Add struct{}

func (InstructionI32Add) Opcode() Opcode           { return OpcodeI32Add }
func (InstructionI32Add) ReadFrom(io.Reader) error { return nil }
