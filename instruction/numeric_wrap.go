package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/opcode"
)

type I32WrapI64 struct{}

func (i *I32WrapI64) Opcode() opcode.Opcode {
	return opcode.OpcodeI32WrapI64
}

func (i *I32WrapI64) ReadOperandsFrom(r io.Reader) error {
	return nil
}
