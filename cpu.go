package main

import "fmt"

// Opcodes
const (
	BRK = 0x00
	LDA = 0xa9
	TAX = 0xaa
	INX = 0xe8
)

const PROG_MEM_ADDRESS = 0x8000

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
	memory         [0xffff]uint8
}

func (c *Cpu) Execute(program []uint8) {
	c.load(program)
	c.run()
}

func (c *Cpu) run() {
	for {
		opcode := c.readMemory(c.ProgramCounter)
		c.ProgramCounter++

		switch opcode {
		case LDA:
			param := c.readMemory(c.ProgramCounter)
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

func (c *Cpu) PrintState() {
	fmt.Printf("Program Counter: %#x\n", c.ProgramCounter)
	fmt.Printf("Register A: %#x\n", c.RegA)
	fmt.Printf("Register X: %#x\n", c.RegX)
}

func (c *Cpu) load(program []uint8) {
	for index, byte := range program {
		memIndex := PROG_MEM_ADDRESS + index
		c.memory[memIndex] = byte
	}
	c.ProgramCounter = PROG_MEM_ADDRESS
}

// nes is little endian, but we deal with that when reading roms.
// index is just a numerical value that we use directly to index into memory.
func (c *Cpu) readMemory(index uint16) uint8 {
	return c.memory[index]
}

func (c *Cpu) writeMemory(index uint16, value uint8) {
	c.memory[index] = value
}

func (c *Cpu) updateFlags(result uint8) {
	c.Status.Zero = (result == 0)
	c.Status.Negative = ((result & (1 << 7)) != 0)
}
