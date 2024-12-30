package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/binary"
)

type I32WrapI64 struct{}

func (i *I32WrapI64) Opcode() binary.Opcode {
	return binary.OpcodeI32WrapI64
}

func (i *I32WrapI64) ReadOperandsFrom(r io.Reader) error {
	return nil
}
