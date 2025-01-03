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
	case opcode.OpcodeUnreachable:
		return new(instruction.Unreachable), nil
	case opcode.OpcodeNop:
		return new(instruction.Nop), nil
	case opcode.OpcodeBlock:
		return new(instruction.Block), nil
	case opcode.OpcodeLoop:
		return new(instruction.Loop), nil
	case opcode.OpcodeIf:
		return new(instruction.If), nil
	case opcode.OpcodeElse:
		return new(instruction.Else), nil
	case opcode.OpcodeEnd:
		return new(instruction.End), nil
	case opcode.OpcodeBr:
		return new(instruction.Br), nil
	case opcode.OpcodeBrIf:
		return new(instruction.BrIf), nil
	case opcode.OpcodeBrTable:
		return new(instruction.BrTable), nil
	case opcode.OpcodeReturn:
		return new(instruction.Return), nil
	case opcode.OpcodeCall:
		return new(instruction.Call), nil
	case opcode.OpcodeDrop:
		return new(instruction.Drop), nil
	case opcode.OpcodeSelect:
		return new(instruction.Select), nil
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
	case opcode.OpcodeF32Store:
		return new(instruction.F32Store), nil
	case opcode.OpcodeF64Store:
		return new(instruction.F64Store), nil
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
	case opcode.OpcodeI32Eqz:
		return new(instruction.I32Eqz), nil
	case opcode.OpcodeI32Eq:
		return new(instruction.I32Eq), nil
	case opcode.OpcodeI32Ne:
		return new(instruction.I32Ne), nil
	case opcode.OpcodeI32LtS:
		return new(instruction.I32LtS), nil
	case opcode.OpcodeI32LtU:
		return new(instruction.I32LtU), nil
	case opcode.OpcodeI32GtS:
		return new(instruction.I32GtS), nil
	case opcode.OpcodeI32GtU:
		return new(instruction.I32GtU), nil
	case opcode.OpcodeI32LeS:
		return new(instruction.I32LeS), nil
	case opcode.OpcodeI32LeU:
		return new(instruction.I32LeU), nil
	case opcode.OpcodeI32GeS:
		return new(instruction.I32GeS), nil
	case opcode.OpcodeI32GeU:
		return new(instruction.I32GeU), nil
	case opcode.OpcodeI64Eqz:
		return new(instruction.I64Eqz), nil
	case opcode.OpcodeI64Eq:
		return new(instruction.I64Eq), nil
	case opcode.OpcodeI64Ne:
		return new(instruction.I64Ne), nil
	case opcode.OpcodeI64LtS:
		return new(instruction.I64LtS), nil
	case opcode.OpcodeI64LtU:
		return new(instruction.I64LtU), nil
	case opcode.OpcodeI64GtS:
		return new(instruction.I64GtS), nil
	case opcode.OpcodeI64GtU:
		return new(instruction.I64GtU), nil
	case opcode.OpcodeI64LeS:
		return new(instruction.I64LeS), nil
	case opcode.OpcodeI64LeU:
		return new(instruction.I64LeU), nil
	case opcode.OpcodeI64GeS:
		return new(instruction.I64GeS), nil
	case opcode.OpcodeI64GeU:
		return new(instruction.I64GeU), nil
	case opcode.OpcodeF32Eq:
		return new(instruction.F32Eq), nil
	case opcode.OpcodeF32Ne:
		return new(instruction.F32Ne), nil
	case opcode.OpcodeF32Lt:
		return new(instruction.F32Lt), nil
	case opcode.OpcodeF32Gt:
		return new(instruction.F32Gt), nil
	case opcode.OpcodeF32Le:
		return new(instruction.F32Le), nil
	case opcode.OpcodeF32Ge:
		return new(instruction.F32Ge), nil
	case opcode.OpcodeF64Eq:
		return new(instruction.F64Eq), nil
	case opcode.OpcodeF64Ne:
		return new(instruction.F64Ne), nil
	case opcode.OpcodeF64Lt:
		return new(instruction.F64Lt), nil
	case opcode.OpcodeF64Gt:
		return new(instruction.F64Gt), nil
	case opcode.OpcodeF64Le:
		return new(instruction.F64Le), nil
	case opcode.OpcodeF64Ge:
		return new(instruction.F64Ge), nil
	case opcode.OpcodeI32Add:
		return new(instruction.I32Add), nil
	case opcode.OpcodeI32Sub:
		return new(instruction.I32Sub), nil
	case opcode.OpcodeFCPrefix:
		return new(instruction.FCPrefix), nil
	default:
		return nil, fmt.Errorf("unknown opcode: %v", op)
	}
}
