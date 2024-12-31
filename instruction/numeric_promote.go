package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/binary"
)

type F64PromoteF32 struct{}

func (f *F64PromoteF32) Opcode() binary.Opcode {
	return binary.OpcodeF64PromoteF32
}

func (f *F64PromoteF32) ReadOperandsFrom(r io.Reader) error {
	return nil
}
