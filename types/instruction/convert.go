package instruction

import (
	"fmt"

	"github.com/Warashi/go-tinywasm/types/binary"
	"github.com/Warashi/go-tinywasm/types/runtime"
)

var ErrInvalidInstruction = fmt.Errorf("invalid instruction")

func Convert(insts []binary.Instruction) ([]runtime.Instruction, error) {
	result := make([]runtime.Instruction, 0, len(insts))
	for _, inst := range insts {
		o, ok := inst.(runtime.Instruction)
		if !ok {
			return nil, ErrInvalidInstruction
		}
		result = append(result, o)
	}

	return result, nil
}
