package instruction

import (
	"io"

	"github.com/Warashi/go-tinywasm/opcode"
)

type I32RemS struct{}

func (i *I32RemS) Opcode() opcode.Opcode {
	return opcode.OpcodeI32RemS
}

func (i *I32RemS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I32RemU struct{}

func (i *I32RemU) Opcode() opcode.Opcode {
	return opcode.OpcodeI32RemU
}

func (i *I32RemU) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64RemS struct{}

func (i *I64RemS) Opcode() opcode.Opcode {
	return opcode.OpcodeI64RemS
}

func (i *I64RemS) ReadOperandsFrom(r io.Reader) error {
	return nil
}

type I64RemU struct{}

func (i *I64RemU) Opcode() opcode.Opcode {
	return opcode.OpcodeI64RemU
}

func (i *I64RemU) ReadOperandsFrom(r io.Reader) error {
	return nil
}
