package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/leb128"
	"github.com/Warashi/go-tinywasm/opcode"
)

type LocalGet struct {
	index uint32
}

func (i *LocalGet) Index() uint32 {
	return i.index
}

func (i *LocalGet) Opcode() opcode.Opcode {
	return opcode.OpcodeLocalGet
}

func (i *LocalGet) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.index, err = leb128.Uint32(r)
	return err
}

type LocalSet struct {
	index uint32
}

func (i *LocalSet) Index() uint32 {
	return i.index
}

func (i *LocalSet) Opcode() opcode.Opcode {
	return opcode.OpcodeLocalSet
}

func (i *LocalSet) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.index, err = leb128.Uint32(r)
	return err
}
