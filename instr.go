package main

// Opcodes
const (
	BRK = 0x00
	JSR = 0x20

	LDA        = 0xa9
	LDA_ZERO   = 0xa5
	LDA_ZERO_X = 0xb5
	LDA_ABS    = 0xad
	LDA_ABS_X  = 0xbd
	LDA_ABS_Y  = 0xb9
	LDA_IND_X  = 0xa1
	LDA_IND_Y  = 0xb1

	STA_ZERO   = 0x85
	STA_ZERO_X = 0x95
	STA_ABS    = 0x8d
	STA_ABS_X  = 0x9d
	STA_ABS_Y  = 0x99
	STA_IND_X  = 0x81
	STA_IND_Y  = 0x91

	TAX = 0xaa
	TAY = 0xa8
	INX = 0xe8
)

// AddressingModes
const (
	IMPLICIT   = 0
	ABSOLUTE   = 1
	ABSOLUTE_X = 2
	ABSOLUTE_Y = 3
	ZERO       = 4
	ZERO_X     = 5
	ZERO_Y     = 6
	IMMEDIATE  = 7
	RELATIVE   = 8
	INDIRECT   = 9
	INDIRECT_X = 10
	INDIRECT_Y = 11
)

type Instruction struct {
	Action         string
	AddressingMode int
	NumberOfBytes  int
}

func Decode(opcode uint8) Instruction {
	return instructionMap[opcode]
}

var instructionMap = map[uint8]Instruction{
	0x00: {"BRK", IMPLICIT, 1},
	0x20: {"JSR", ABSOLUTE, 3},
	0xa9: {"LDA", IMMEDIATE, 2},
	0xa5: {"LDA", ZERO, 2},
	0xb5: {"LDA", ZERO_X, 2},
	0xad: {"LDA", ABSOLUTE, 3},
	0xbd: {"LDA", ABSOLUTE_X, 3},
	0xb9: {"LDA", ABSOLUTE_Y, 3},
	0xa1: {"LDA", INDIRECT_X, 2},
	0xb1: {"LDA", INDIRECT_Y, 2},
	0x85: {"STA", ZERO, 2},
	0x95: {"STA", ZERO_X, 2},
	0x8d: {"STA", ABSOLUTE, 3},
	0x9d: {"STA", ABSOLUTE_X, 3},
	0x99: {"STA", ABSOLUTE_Y, 3},
	0x81: {"STA", INDIRECT_X, 2},
	0x91: {"STA", INDIRECT_Y, 2},
	0xaa: {"TAX", IMPLICIT, 1},
	0xa8: {"TAY", IMPLICIT, 1},
	0xe8: {"INX", IMPLICIT, 1},
}
