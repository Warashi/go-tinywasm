package binary

type Opcode byte

const (
	OpcodeEnd      Opcode = 0x0b
	OpcodeCall     Opcode = 0x10
	OpcodeLocalGet Opcode = 0x20
	OpcodeLocalSet Opcode = 0x21
	OpcodeI32Const Opcode = 0x41
	OpcodeI32Add   Opcode = 0x6a
)
