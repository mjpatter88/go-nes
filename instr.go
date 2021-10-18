package main

// Opcodes
const (
	BRK = 0x00
	CLC = 0x18
	JSR = 0x20
	RTS = 0x60

	LDA        = 0xa9
	LDA_ZERO   = 0xa5
	LDA_ZERO_X = 0xb5
	LDA_ABS    = 0xad
	LDA_ABS_X  = 0xbd
	LDA_ABS_Y  = 0xb9
	LDA_IND_X  = 0xa1
	LDA_IND_Y  = 0xb1

	AND        = 0x29
	AND_ZERO   = 0x25
	AND_ZERO_X = 0x35
	AND_ABS    = 0x2d
	AND_ABS_X  = 0x3d
	AND_ABS_Y  = 0x39
	AND_IND_X  = 0x21
	AND_IND_Y  = 0x31

	ADC        = 0x69
	ADC_ZERO   = 0x65
	ADC_ZERO_X = 0x75
	ADC_ABS    = 0x6d
	ADC_ABS_X  = 0x7d
	ADC_ABS_Y  = 0x79
	ADC_IND_X  = 0x61
	ADC_IND_Y  = 0x71

	CMP        = 0xc9
	CMP_ZERO   = 0xc5
	CMP_ZERO_X = 0xd5
	CMP_ABS    = 0xcd
	CMP_ABS_X  = 0xdd
	CMP_ABS_Y  = 0xd9
	CMP_IND_X  = 0xc1
	CMP_IND_Y  = 0xd1

	CPX      = 0xe0
	CPX_ZERO = 0xe4
	CPX_ABS  = 0xec

	CPY      = 0xc0
	CPY_ZERO = 0xc4
	CPY_ABS  = 0xcc

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
	0x18: {"CLC", IMPLICIT, 1},
	0x20: {"JSR", ABSOLUTE, 3},
	0x60: {"RTS", IMPLICIT, 1},
	0xa9: {"LDA", IMMEDIATE, 2},
	0xa5: {"LDA", ZERO, 2},
	0xb5: {"LDA", ZERO_X, 2},
	0xad: {"LDA", ABSOLUTE, 3},
	0xbd: {"LDA", ABSOLUTE_X, 3},
	0xb9: {"LDA", ABSOLUTE_Y, 3},
	0xa1: {"LDA", INDIRECT_X, 2},
	0xb1: {"LDA", INDIRECT_Y, 2},
	0x29: {"AND", IMMEDIATE, 2},
	0x25: {"AND", ZERO, 2},
	0x35: {"AND", ZERO_X, 2},
	0x2d: {"AND", ABSOLUTE, 3},
	0x3d: {"AND", ABSOLUTE_X, 3},
	0x39: {"AND", ABSOLUTE_Y, 3},
	0x21: {"AND", INDIRECT_X, 2},
	0x31: {"AND", INDIRECT_Y, 2},
	0x69: {"ADC", IMMEDIATE, 2},
	0x65: {"ADC", ZERO, 2},
	0x75: {"ADC", ZERO_X, 2},
	0x6d: {"ADC", ABSOLUTE, 3},
	0x7d: {"ADC", ABSOLUTE_X, 3},
	0x79: {"ADC", ABSOLUTE_Y, 3},
	0x61: {"ADC", INDIRECT_X, 2},
	0x71: {"ADC", INDIRECT_Y, 2},
	0xc9: {"CMP", IMMEDIATE, 2},
	0xc5: {"CMP", ZERO, 2},
	0xd5: {"CMP", ZERO_X, 2},
	0xcd: {"CMP", ABSOLUTE, 3},
	0xdd: {"CMP", ABSOLUTE_X, 3},
	0xd9: {"CMP", ABSOLUTE_Y, 3},
	0xc1: {"CMP", INDIRECT_X, 2},
	0xd1: {"CMP", INDIRECT_Y, 2},
	0xe0: {"CPX", IMMEDIATE, 2},
	0xe4: {"CPX", ZERO, 2},
	0xec: {"CPX", ABSOLUTE, 3},
	0xc0: {"CPY", IMMEDIATE, 2},
	0xc4: {"CPY", ZERO, 2},
	0xcc: {"CPY", ABSOLUTE, 3},
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
