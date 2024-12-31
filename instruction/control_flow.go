package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/leb128"
	"github.com/Warashi/go-tinywasm/opcode"
)

type If struct {
	block Block
}

func (*If) Opcode() opcode.Opcode { return opcode.OpcodeIf }
func (i *If) ReadOperandsFrom(r io.Reader) error {
	return i.block.decode(r)
}
func (i *If) Block() Block { return i.block }

type End struct{}

func (*End) Opcode() opcode.Opcode { return opcode.OpcodeEnd }

func (*End) ReadOperandsFrom(io.Reader) error { return nil }

type Return struct{}

func (*Return) Opcode() opcode.Opcode { return opcode.OpcodeReturn }

func (*Return) ReadOperandsFrom(io.Reader) error { return nil }

type Call struct {
	index uint32
}

func (c *Call) Opcode() opcode.Opcode { return opcode.OpcodeCall }

func (c *Call) ReadOperandsFrom(r io.Reader) error {
	var err error
	c.index, err = leb128.Uint32(r)
	return err
}

func (c *Call) Index() uint32 { return c.index }
