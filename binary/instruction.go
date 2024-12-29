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

type InstructionLocalGet struct{ idx uint32 }

func (InstructionLocalGet) Opcode() Opcode { return OpcodeLocalGet }
func (i *InstructionLocalGet) ReadFrom(r io.Reader) error {
	var err error
	i.idx, err = leb128.Uint32(r)
	return err
}

type InstructionI32Add struct{}

func (InstructionI32Add) Opcode() Opcode           { return OpcodeI32Add }
func (InstructionI32Add) ReadFrom(io.Reader) error { return nil }
