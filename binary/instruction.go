package binary

import (
	"fmt"
	"io"

	"github.com/Warashi/wasmium/instruction"
	"github.com/Warashi/wasmium/opcode"
)

type inst interface {
	Opcode() opcode.Opcode
	ReadOperandsFrom(r io.Reader) error
}

func fromOpcode(op opcode.Opcode) (inst, error) {
	switch op {
	case opcode.OpcodeLoop:
		return new(instruction.Loop), nil
	case opcode.OpcodeIf:
		return new(instruction.If), nil
	case opcode.OpcodeEnd:
		return new(instruction.End), nil
	case opcode.OpcodeBr:
		return new(instruction.Br), nil
	case opcode.OpcodeReturn:
		return new(instruction.Return), nil
	case opcode.OpcodeCall:
		return new(instruction.Call), nil
	case opcode.OpcodeDrop:
		return new(instruction.Drop), nil
	case opcode.OpcodeLocalGet:
		return new(instruction.LocalGet), nil
	case opcode.OpcodeLocalSet:
		return new(instruction.LocalSet), nil
	case opcode.OpcodeGlobalGet:
		return new(instruction.GlobalGet), nil
	case opcode.OpcodeGlobalSet:
		return new(instruction.GlobalSet), nil
	case opcode.OpcodeI32Load:
		return new(instruction.I32Load), nil
	case opcode.OpcodeI64Load:
		return new(instruction.I64Load), nil
	case opcode.OpcodeF32Load:
		return new(instruction.F32Load), nil
	case opcode.OpcodeF64Load:
		return new(instruction.F64Load), nil
	case opcode.OpcodeI32Load8S:
		return new(instruction.I32Load8S), nil
	case opcode.OpcodeI32Load8U:
		return new(instruction.I32Load8U), nil
	case opcode.OpcodeI32Load16S:
		return new(instruction.I32Load16S), nil
	case opcode.OpcodeI32Load16U:
		return new(instruction.I32Load16U), nil
	case opcode.OpcodeI64Load8S:
		return new(instruction.I64Load8S), nil
	case opcode.OpcodeI64Load8U:
		return new(instruction.I64Load8U), nil
	case opcode.OpcodeI64Load16S:
		return new(instruction.I64Load16S), nil
	case opcode.OpcodeI64Load16U:
		return new(instruction.I64Load16U), nil
	case opcode.OpcodeI64Load32S:
		return new(instruction.I64Load32S), nil
	case opcode.OpcodeI64Load32U:
		return new(instruction.I64Load32U), nil
	case opcode.OpcodeI32Store:
		return new(instruction.I32Store), nil
	case opcode.OpcodeI64Store:
		return new(instruction.I64Store), nil
	case opcode.OpcodeI32Store8:
		return new(instruction.I32Store8), nil
	case opcode.OpcodeI32Store16:
		return new(instruction.I32Store16), nil
	case opcode.OpcodeI64Store8:
		return new(instruction.I64Store8), nil
	case opcode.OpcodeI64Store16:
		return new(instruction.I64Store16), nil
	case opcode.OpcodeI64Store32:
		return new(instruction.I64Store32), nil
	case opcode.OpcodeI32Const:
		return new(instruction.I32Const), nil
	case opcode.OpcodeI64Const:
		return new(instruction.I64Const), nil
	case opcode.OpcodeF32Const:
		return new(instruction.F32Const), nil
	case opcode.OpcodeF64Const:
		return new(instruction.F64Const), nil
	case opcode.OpcodeI32LtS:
		return new(instruction.I32LtS), nil
	case opcode.OpcodeI32Add:
		return new(instruction.I32Add), nil
	case opcode.OpcodeI32Sub:
		return new(instruction.I32Sub), nil
	default:
		return nil, fmt.Errorf("unknown opcode: %v", op)
	}
}
