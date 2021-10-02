package main

import "fmt"

// Opcodes
const (
	BRK = 0x00
	LDA = 0xa9
	TAX = 0xaa
)

type StatusRegister struct {
	Carry     bool
	Zero      bool
	Interrupt bool
	Decimal   bool
	Break     bool
	Unused    bool
	Overflow  bool
	Negative  bool
}

type Cpu struct {
	RegA           uint8
	RegX           uint8
	Status         StatusRegister
	ProgramCounter uint16
}

func (c *Cpu) execute(instructions []uint8) {
	for {
		opcode := instructions[c.ProgramCounter]
		c.ProgramCounter++

		switch opcode {
		case LDA:
			param := instructions[c.ProgramCounter]
			c.ProgramCounter++
			c.RegA = param
			c.Status.Zero = (c.RegA == 0)
			c.Status.Negative = ((c.RegA & (1 << 7)) != 0)
		case TAX:
			c.RegX = c.RegA
			c.Status.Zero = (c.RegX == 0)
			c.Status.Negative = ((c.RegX & (1 << 7)) != 0)
		case BRK:
			c.Status.Break = true
			return
		default:
			panic(fmt.Errorf("unsuppored opcode %#x", opcode))
		}
	}
}
