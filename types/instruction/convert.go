package instruction

import (
	"fmt"

	"github.com/Warashi/wasmium/types/binary"
	"github.com/Warashi/wasmium/types/runtime"
)

var ErrInvalidInstruction = fmt.Errorf("invalid instruction")

func Convert(insts []binary.Instruction) ([]runtime.Instruction, error) {
	result := make([]runtime.Instruction, 0, len(insts))
	for _, inst := range insts {
		o, ok := inst.(runtime.Instruction)
		if !ok {
			return nil, fmt.Errorf("%w: %T", ErrInvalidInstruction, inst)
		}
		result = append(result, o)
	}

	return result, nil
}
