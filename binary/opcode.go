package binary

type Opcode byte

const (
	OpcodeIf       Opcode = 0x04
	OpcodeEnd      Opcode = 0x0b
	OpcodeReturn   Opcode = 0x0f
	OpcodeCall     Opcode = 0x10
	OpcodeLocalGet Opcode = 0x20
	OpcodeLocalSet Opcode = 0x21
	OpcodeI32Store Opcode = 0x36
	OpcodeI32Const Opcode = 0x41
	OpcodeI32LtS   Opcode = 0x48
	OpcodeI32Add   Opcode = 0x6a
	OpcodeI32Sub   Opcode = 0x6b
)
