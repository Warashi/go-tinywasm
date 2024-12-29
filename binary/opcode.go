package binary

type Opcode byte

const (
	OpcodeEnd      Opcode = 0x0b
	OpcodeLocalGet Opcode = 0x20
	OpcodeI32Add   Opcode = 0x6a
)
