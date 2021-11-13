package main

// Opcodes
const (
	BRK = 0x00
	CLC = 0x18
	JSR = 0x20
	RTS = 0x60

	BIT_ZERO = 0x24
	BIT_ABS  = 0x2c

	LDA        = 0xa9
	LDA_ZERO   = 0xa5
	LDA_ZERO_X = 0xb5
	LDA_ABS    = 0xad
	LDA_ABS_X  = 0xbd
	LDA_ABS_Y  = 0xb9
	LDA_IND_X  = 0xa1
	LDA_IND_Y  = 0xb1

	LDX        = 0xa2
	LDX_ZERO   = 0xa6
	LDX_ZERO_Y = 0xb6
	LDX_ABS    = 0xae
	LDX_ABS_Y  = 0xbe

	LDY        = 0xa0
	LDY_ZERO   = 0xa4
	LDY_ZERO_X = 0xb4
	LDY_ABS    = 0xac
	LDY_ABS_X  = 0xbc

	LSR        = 0x4a
	LSR_ZERO   = 0x46
	LSR_ZERO_X = 0x56
	LSR_ABS    = 0x4e
	LSR_ABS_X  = 0x5e

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

	INC_ZERO   = 0xe6
	INC_ZERO_X = 0xf6
	INC_ABS    = 0xee
	INC_ABS_X  = 0xfe

	TAX = 0xaa
	TXA = 0x8a
	DEX = 0xca
	INX = 0xe8
	TAY = 0xa8
	TYA = 0x98
	DEY = 0x88
	INY = 0xc8

	BPL = 0x10
	BMI = 0x30

	BVC = 0x50
	BVS = 0x70
	BCC = 0x90
	BCS = 0xb0

	BEQ = 0xf0
	BNE = 0xd0
)

// AddressingModes
const (
	IMPLICIT    = 0
	ABSOLUTE    = 1
	ABSOLUTE_X  = 2
	ABSOLUTE_Y  = 3
	ZERO        = 4
	ZERO_X      = 5
	ZERO_Y      = 6
	IMMEDIATE   = 7
	RELATIVE    = 8
	INDIRECT    = 9
	INDIRECT_X  = 10
	INDIRECT_Y  = 11
	ACCUMULATOR = 12
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
	0x24: {"BIT", ZERO, 2},
	0x2c: {"BIT", ABSOLUTE, 3},
	0x60: {"RTS", IMPLICIT, 1},
	0xa9: {"LDA", IMMEDIATE, 2},
	0xa5: {"LDA", ZERO, 2},
	0xb5: {"LDA", ZERO_X, 2},
	0xad: {"LDA", ABSOLUTE, 3},
	0xbd: {"LDA", ABSOLUTE_X, 3},
	0xb9: {"LDA", ABSOLUTE_Y, 3},
	0xa1: {"LDA", INDIRECT_X, 2},
	0xb1: {"LDA", INDIRECT_Y, 2},
	0xa2: {"LDX", IMMEDIATE, 2},
	0xa6: {"LDX", ZERO, 2},
	0xb6: {"LDX", ZERO_Y, 2},
	0xae: {"LDX", ABSOLUTE, 3},
	0xbe: {"LDX", ABSOLUTE_Y, 3},
	0xa0: {"LDY", IMMEDIATE, 2},
	0xa4: {"LDY", ZERO, 2},
	0xb4: {"LDY", ZERO_X, 2},
	0xac: {"LDY", ABSOLUTE, 3},
	0xbc: {"LDY", ABSOLUTE_X, 3},
	0x4a: {"LSR", ACCUMULATOR, 1},
	0x46: {"LSR", ZERO, 2},
	0x56: {"LSR", ZERO_X, 2},
	0x4e: {"LSR", ABSOLUTE, 3},
	0x5e: {"LSR", ABSOLUTE_X, 3},
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
	0xe6: {"INC", ZERO, 2},
	0xf6: {"INC", ZERO_X, 2},
	0xee: {"INC", ABSOLUTE, 3},
	0xfe: {"INC", ABSOLUTE_X, 3},
	0xaa: {"TAX", IMPLICIT, 1},
	0x8a: {"TXA", IMPLICIT, 1},
	0xca: {"DEX", IMPLICIT, 1},
	0xe8: {"INX", IMPLICIT, 1},
	0xa8: {"TAY", IMPLICIT, 1},
	0x98: {"TYA", IMPLICIT, 1},
	0x88: {"DEY", IMPLICIT, 1},
	0xc8: {"INY", IMPLICIT, 1},
	0x10: {"BPL", RELATIVE, 2},
	0x30: {"BMI", RELATIVE, 2},
	0x50: {"BVC", RELATIVE, 2},
	0x70: {"BVS", RELATIVE, 2},
	0x90: {"BCC", RELATIVE, 2},
	0xb0: {"BCS", RELATIVE, 2},
	0xf0: {"BEQ", RELATIVE, 2},
	0xd0: {"BNE", RELATIVE, 2},
}
