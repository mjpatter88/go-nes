package main

import "fmt"

// Opcodes
const (
	BRK        = 0x00
	LDA        = 0xa9
	LDA_ZERO   = 0xa5
	LDA_ZERO_X = 0xb5
	LDA_ABS    = 0xad
	LDA_ABS_X  = 0xbd
	TAX        = 0xaa
	TAY        = 0xa8
	INX        = 0xe8
)

//Memory Addresses
const (
	PROG_MEM_ADDRESS           = 0x8000
	PROG_REFERENCE_MEM_ADDRESS = 0xfffc
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
	RegY           uint8
	Status         StatusRegister
	ProgramCounter uint16
	memory         [0xffff]uint8
}

func (c *Cpu) Execute(program []uint8) {
	c.load(program)
	c.reset()
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
			c.instrLDA(param)
		case LDA_ZERO:
			address := c.readMemory(c.ProgramCounter)
			c.ProgramCounter++
			c.instrLDA(c.readMemory(uint16(address)))
		case LDA_ZERO_X:
			// Address is a byte and the overflow/wrap behavior is intentional.
			address := c.readMemory(c.ProgramCounter)
			address += c.RegX
			c.ProgramCounter++
			c.instrLDA(c.readMemory(uint16(address)))
		case LDA_ABS:
			// Address is two bytes little endian
			addressA := c.readMemory(c.ProgramCounter)
			c.ProgramCounter++
			addressB := c.readMemory(c.ProgramCounter)
			c.ProgramCounter++
			address := (uint16(addressA) << 8) | (uint16(addressB))
			c.instrLDA(c.readMemory(uint16(address)))
		case LDA_ABS_X:
			// Address is two bytes little endian
			addressA := c.readMemory(c.ProgramCounter)
			c.ProgramCounter++
			addressB := c.readMemory(c.ProgramCounter)
			c.ProgramCounter++
			address := (uint16(addressA) << 8) | (uint16(addressB)) + uint16(c.RegX)
			c.instrLDA(c.readMemory(uint16(address)))
		case TAX:
			c.instrTAX()
		case TAY:
			c.instrTAY()
		case INX:
			c.instrINX()
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
	// nes spec says to write program memory address into mem address 0xFFFC
	// this value is then read into the program counter on system reset
	// NOTE: presumably the program could be loaded into memory at a different address
	// otherwise this additional level of indirection seems pointless. (?)
	c.writeMemory_u16(PROG_REFERENCE_MEM_ADDRESS, PROG_MEM_ADDRESS)
	c.ProgramCounter = 0x8000
}

func (c *Cpu) reset() {
	c.resetStatus()
	c.RegA = 0
	c.RegX = 0
	c.ProgramCounter = c.readMemory_u16(PROG_REFERENCE_MEM_ADDRESS)
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

func (c *Cpu) resetStatus() {
	c.Status = StatusRegister{}
}

func (c *Cpu) instrLDA(param uint8) {
	c.RegA = param
	c.updateFlags(c.RegA)
}

func (c *Cpu) instrTAX() {
	c.RegX = c.RegA
	c.updateFlags(c.RegX)
}

func (c *Cpu) instrTAY() {
	c.RegY = c.RegA
	c.updateFlags(c.RegY)
}

func (c *Cpu) instrINX() {
	c.RegX++
	c.updateFlags(c.RegX)
}
