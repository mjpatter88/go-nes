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

func (c *Cpu) readMemory(index uint16) uint8 {
	return c.memory[index]
}

func (c *Cpu) writeMemory(index uint16, value uint8) {
	c.memory[index] = value
}

// nes is little-endian so 16-bit values read from memory need to handle this byte order.
// NOTE: this just impacts the 16-bit values from memory, not the 16-bit memory index.
func (c *Cpu) readMemory_u16(index uint16) uint16 {
	firstByte := uint16(c.readMemory(index))
	secondByte := uint16(c.readMemory(index + 1))
	return (secondByte << 8) | (firstByte)
}

// nes is little-endian so 16-bit values written to memory need to handle this byte order.
// NOTE: this just impacts the 16-bit values written to memory, not the 16-bit memory index.
func (c *Cpu) writeMemory_u16(index uint16, value uint16) {
	firstByte := (value) & 0xFF
	secondByte := (value >> 8) & 0xFF
	c.writeMemory(index, uint8(firstByte))
	c.writeMemory(index+1, uint8(secondByte))
}

func (c *Cpu) updateFlags(result uint8) {
	c.Status.Zero = (result == 0)
	c.Status.Negative = ((result & (1 << 7)) != 0)
}
