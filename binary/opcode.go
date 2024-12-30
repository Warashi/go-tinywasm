package binary

type Opcode byte

const (
	OpcodeEnd      Opcode = 0x0b
	OpcodeCall     Opcode = 0x10
	OpcodeLocalGet Opcode = 0x20
	OpcodeI32Add   Opcode = 0x6a
)
