package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/binary"
)

type F32DemoteF64 struct{}

func (f *F32DemoteF64) Opcode() binary.Opcode {
	return binary.OpcodeF32DemoteF64
}

func (f *F32DemoteF64) ReadOperandsFrom(r io.Reader) error {
	return nil
}
