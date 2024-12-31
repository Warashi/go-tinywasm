package binary

import (
	"fmt"
	"io"

	"github.com/Warashi/go-tinywasm/opcode"
)

type instruction interface {
	Opcode() opcode.Opcode
	ReadOperandsFrom(r io.Reader) error
}

func fromOpcode(code opcode.Opcode) (instruction, error) {
	switch code {
	default:
		return nil, fmt.Errorf("unknown opcode: %2x", code)
	}
}
