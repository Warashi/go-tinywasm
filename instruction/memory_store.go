package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/leb128"
	"github.com/Warashi/go-tinywasm/opcode"
)

type I32Store struct {
	align  uint32
	offset uint32
}

func (i *I32Store) Align() uint32 {
	return i.align
}

func (i *I32Store) Offset() uint32 {
	return i.offset
}

func (i *I32Store) Opcode() opcode.Opcode {
	return opcode.OpcodeI32Store
}

func (i *I32Store) ReadOperandsFrom(r io.Reader) error {
	var err error
	i.align, err = leb128.Uint32(r)
	if err != nil {
		return err
	}
	i.offset, err = leb128.Uint32(r)
	return err
}

// TODO: implement the rest of the store instructions
