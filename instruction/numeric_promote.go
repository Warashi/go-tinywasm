package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/opcode"
)

type F64PromoteF32 struct{}

func (f *F64PromoteF32) Opcode() opcode.Opcode {
	return opcode.OpcodeF64PromoteF32
}

func (f *F64PromoteF32) ReadOperandsFrom(r io.Reader) error {
	return nil
}
