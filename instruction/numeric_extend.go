package instruction

import (
	"io"

	"github.com/Warashi/wasmium/opcode"
)

type I64ExtendSI32 struct{}

func (i *I64ExtendSI32) Opcode() opcode.Opcode {
	return opcode.OpcodeI64ExtendI32S
}

func (i *I64ExtendSI32) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64ExtendUI32 struct{}

func (i *I64ExtendUI32) Opcode() opcode.Opcode {
	return opcode.OpcodeI64ExtendI32U
}

func (i *I64ExtendUI32) ReadOperandsFrom(r io.Reader) error {
	return nil
}
