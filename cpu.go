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
	LDA_ABS_Y  = 0xb9
	LDA_IND_X  = 0xa1
	LDA_IND_Y  = 0xb1
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
		// TODO(mjpatter88): future refactor idea: first decode opcode into instr + mode.
		// Then have two switch statements. First uses the addressing mode to calculate
		// a "value". Second switches on the instruction and calls the
		// corresponding instruction with the "value" calculated above.
		opcode := c.readMemory(c.ProgramCounter)
		c.ProgramCounter++

		switch opcode {
		case LDA:
			c.instrLDA(c.ImmediateMode())
		case LDA_ZERO:
			c.instrLDA(c.ZeroMode())
		case LDA_ZERO_X:
			c.instrLDA(c.ZeroXMode())
		case LDA_ABS:
			c.instrLDA(c.AbsoluteMode())
		case LDA_ABS_X:
			c.instrLDA(c.AbsoluteXMode())
		case LDA_ABS_Y:
			c.instrLDA(c.AbsoluteYMode())
		case LDA_IND_X:
			c.instrLDA(c.IndirectXMode())
		case LDA_IND_Y:
			// Initial address is a byte and the overflow/wrap behavior is intentional.
			address := c.readMemory(c.ProgramCounter)
			address += c.RegY
			c.ProgramCounter++

			// Use the initial address to read an address from memory.
			index := uint16(address)
			// Address is two bytes little endian (LSB first)
			addressA := c.readMemory(index)
			addressB := c.readMemory(index + 1)
			finalAddress := (uint16(addressB) << 8) | (uint16(addressA))
			c.instrLDA(c.readMemory(uint16(finalAddress)))
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

// https://skilldrick.github.io/easy6502/#addressing
func (c *Cpu) ImmediateMode() uint8 {
	// Use the program counter to return the value directly after the opcode.
	//
	// Incrememnts program counter by 1.
	value := c.readMemory(c.ProgramCounter)
	c.ProgramCounter++
	return value
}

func (c *Cpu) ZeroMode() uint8 {
	// Use the value stored directly after the opcode as an index into memory and return the value stored there.
	//
	// Incrememnts program counter by 1.
	address := c.readMemory(c.ProgramCounter)
	c.ProgramCounter++
	return c.readMemory(uint16(address))
}

func (c *Cpu) ZeroXMode() uint8 {
	// Calculate a memory address by adding the value stored directly after the opcode
	// add the value in the x register.
	// Return the value stored at that address.
	//
	// Incrememnts program counter by 1.

	// Address is a byte and the overflow/wrap behavior is intentional.
	address := c.readMemory(c.ProgramCounter)
	address += c.RegX
	c.ProgramCounter++
	return c.readMemory(uint16(address))
}

func (c *Cpu) AbsoluteMode() uint8 {
	// Use the two bytes stored directly after the opcode as an index into memory.
	// Treat them as litte endian (LSB first).
	//
	// Incrememnts program counter by 2.

	// TODO(mjpatter88): In all these cases, use the readMem_u16 function.
	addressA := c.readMemory(c.ProgramCounter)
	c.ProgramCounter++
	addressB := c.readMemory(c.ProgramCounter)
	c.ProgramCounter++
	address := (uint16(addressB) << 8) | (uint16(addressA))
	return c.readMemory(address)
}

func (c *Cpu) AbsoluteXMode() uint8 {
	// Same as AbsoluteMode but the value in the X register is added to
	// the memory address.
	//
	// Incrememnts program counter by 2.

	addressA := c.readMemory(c.ProgramCounter)
	c.ProgramCounter++
	addressB := c.readMemory(c.ProgramCounter)
	c.ProgramCounter++
	address := (uint16(addressB) << 8) | (uint16(addressA)) + uint16(c.RegX)
	return c.readMemory(uint16(address))
}

func (c *Cpu) AbsoluteYMode() uint8 {
	// Same as AbsoluteMode but the value in the Y register is added to
	// the memory address.
	//
	// Incrememnts program counter by 2.

	addressA := c.readMemory(c.ProgramCounter)
	c.ProgramCounter++
	addressB := c.readMemory(c.ProgramCounter)
	c.ProgramCounter++
	address := (uint16(addressB) << 8) | (uint16(addressA)) + uint16(c.RegY)
	return c.readMemory(uint16(address))
}

func (c *Cpu) IndirectXMode() uint8 {
	// Use the two bytes stored directly after the opcode as an index into memory.
	// Treat them as litte endian (LSB first). Add the value in the X register.
	// Use this sum as an initial index. Lookup the value stored in memory at
	// this index and then use that value as another index.
	// Return the value stored in memory at that second index.
	//
	// Incrememnts program counter by 2.

	// Initial address is a byte and the overflow/wrap behavior is intentional.
	address := c.readMemory(c.ProgramCounter)
	address += c.RegX
	c.ProgramCounter++

	// Use the initial address to read an address from memory.
	index := uint16(address)
	// Address is two bytes little endian (LSB first)
	addressA := c.readMemory(index)
	addressB := c.readMemory(index + 1)
	finalAddress := (uint16(addressB) << 8) | (uint16(addressA))
	return c.readMemory(uint16(finalAddress))
}
