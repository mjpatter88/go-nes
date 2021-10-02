package main

import "fmt"

// Opcodes
const (
	BRK = 0x00
	LDA = 0xa9
	TAX = 0xaa
	INX = 0xe8
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

func (c *Cpu) Execute(instructions []uint8) {
	for {
		opcode := instructions[c.ProgramCounter]
		c.ProgramCounter++

		switch opcode {
		case LDA:
			param := instructions[c.ProgramCounter]
			c.ProgramCounter++
			c.RegA = param
			c.updateFlags(c.RegA)
		case TAX:
			c.RegX = c.RegA
			c.updateFlags(c.RegX)
		case INX:
			c.RegX++
			c.updateFlags(c.RegX)
		case BRK:
			c.Status.Break = true
			return
		default:
			panic(fmt.Errorf("unsuppored opcode %#x", opcode))
		}
	}
}

func (c *Cpu) printState() {
	fmt.Printf("Program Counter: %#x\n", c.ProgramCounter)
	fmt.Printf("Register A: %#x\n", c.RegA)
	fmt.Printf("Register X: %#x\n", c.RegX)
}

func (c *Cpu) updateFlags(result uint8) {
	c.Status.Zero = (result == 0)
	c.Status.Negative = ((result & (1 << 7)) != 0)
}
