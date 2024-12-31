package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/opcode"
)

type F32DemoteF64 struct{}

func (f *F32DemoteF64) Opcode() opcode.Opcode {
	return opcode.OpcodeF32DemoteF64
}

func (f *F32DemoteF64) ReadOperandsFrom(r io.Reader) error {
	return nil
}
