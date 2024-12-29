package execution

import (
	"fmt"
	"io"

	"github.com/Warashi/go-tinywasm/binary"
)

type Frame struct {
	programCounter int
	stackPointer   int
	instructions   []binary.Instruction
	arity          int
	locals         []Value
}

type Runtime struct {
	store     *Store
	stack     []Value
	callStack []Frame
}

func NewRuntime(r io.Reader) (*Runtime, error) {
	module, err := binary.NewModule(r)
	if err != nil {
		return nil, fmt.Errorf("failed to create module: %w", err)
	}

	store, err := NewStore(module)
	if err != nil {
		return nil, fmt.Errorf("failed to create store: %w", err)
	}

	return &Runtime{
		store: store,
	}, nil
}
