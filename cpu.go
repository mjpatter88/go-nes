package main

import "fmt"

//Memory Addresses
const (
	DEFAULT_PROG_MEM_ADDRESS   = 0x8000
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
	StackPointer   uint8
}

func (c *Cpu) Execute(program []uint8) {
	c.ExecuteAtAddress(program, DEFAULT_PROG_MEM_ADDRESS)
}

// Loads the program into memory  at the specified address and executes it
func (c *Cpu) ExecuteAtAddress(program []uint8, address uint16) {
	c.load(program, address)
	c.reset()
	c.run()
}

func (c *Cpu) run() {

	for !c.Status.Break {

		opcode := c.readMemory(c.ProgramCounter)
		instr := Decode(opcode)

		didJump := false
		var param uint16

		switch instr.AddressingMode {
		case IMPLICIT:
			param = 0
		case ABSOLUTE:
			param = c.AbsoluteMode()
		case ABSOLUTE_X:
			param = c.AbsoluteXMode()
		case ABSOLUTE_Y:
			param = c.AbsoluteYMode()
		case ZERO:
			param = c.ZeroMode()
		case ZERO_X:
			param = c.ZeroXMode()
		case ZERO_Y:
			panic(fmt.Errorf("unsuppored addressing mode %q", instr.AddressingMode))
		case IMMEDIATE:
			param = c.ImmediateMode()
		case RELATIVE:
			panic(fmt.Errorf("unsuppored addressing mode %q", instr.AddressingMode))
		case INDIRECT:
			panic(fmt.Errorf("unsuppored addressing mode %q", instr.AddressingMode))
		case INDIRECT_X:
			param = c.IndirectXMode()
		case INDIRECT_Y:
			param = c.IndirectYMode()
		}

		switch instr.Action {
		case "LDA":
			c.instrLDA(param)
		case "AND":
			c.instrAND(param)
		case "ADC":
			c.instrADC(param)
		case "STA":
			c.instrSTA(param)
		case "TAX":
			c.instrTAX()
		case "TAY":
			c.instrTAY()
		case "INX":
			c.instrINX()
		case "JSR":
			c.instrJSR(param)
			didJump = true
		case "RTS":
			c.instrRTS(param)
			didJump = true
		case "CLC":
			c.instrCLC()
		case "BRK":
			c.Status.Break = true
		default:
			panic(fmt.Errorf("unsuppored opcode %#x at pc: %#x", opcode, c.ProgramCounter))
		}

		// Jump instructions are expected to manually update the program counter themselves
		if !didJump {
			for i := 0; i < instr.NumberOfBytes; i++ {
				c.ProgramCounter++
			}
		}
	}
}

func (c *Cpu) PrintState() {
	fmt.Printf("Program Counter: %#x\n", c.ProgramCounter)
	fmt.Printf("Register A: %#x\n", c.RegA)
	fmt.Printf("Register X: %#x\n", c.RegX)
}

func (c *Cpu) load(program []uint8, address uint16) {
	for index, byte := range program {
		memIndex := address + uint16(index)
		c.memory[memIndex] = byte
	}
	// nes spec says to write program memory address into mem address 0xFFFC
	// this value is then read into the program counter on system reset
	c.writeMemory_u16(PROG_REFERENCE_MEM_ADDRESS, address)
	c.ProgramCounter = 0x8000
}

func (c *Cpu) reset() {
	c.resetStatus()
	c.RegA = 0
	c.RegX = 0
	c.ProgramCounter = c.readMemory_u16(PROG_REFERENCE_MEM_ADDRESS)
	c.StackPointer = 0xff
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

func (c *Cpu) instrLDA(param uint16) {
	c.RegA = c.readMemory(param)
	c.updateFlags(c.RegA)
}

func (c *Cpu) instrAND(param uint16) {
	value := c.readMemory(param)
	c.RegA &= value
	c.updateFlags(c.RegA)
}

func (c *Cpu) instrADC(param uint16) {
	// TODO(mjpatter) handle overflow and carry flags correctly
	value := c.readMemory(param)
	c.RegA += value
	c.updateFlags(c.RegA)
}

func (c *Cpu) instrSTA(param uint16) {
	c.writeMemory(param, c.RegA)
	c.resetStatus()
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

func (c *Cpu) instrCLC() {
	c.Status.Carry = false
}

func (c *Cpu) instrJSR(param uint16) {
	// We'll write the two bytes at once, so write it to SP - 1 (ends up writing to SP-1 and SP)
	index := 0x0100 | uint16((c.StackPointer - 1))
	// JSR length is 3 and we want to store the address of the next insturction - 1.
	value := c.ProgramCounter + 3 - 1
	c.writeMemory_u16(index, value)

	c.StackPointer -= 2
	c.ProgramCounter = param
}

func (c *Cpu) instrRTS(param uint16) {
	// Read two bytes from the top of the stack.
	index := 0x0100 | uint16((c.StackPointer + 1))
	value := c.readMemory_u16(index)
	c.StackPointer += 2
	c.ProgramCounter = value + 1
}

// https://skilldrick.github.io/easy6502/#addressing
func (c *Cpu) ImmediateMode() uint16 {
	// Return the address of the value directly after the opcode.
	return c.ProgramCounter + 1
}

func (c *Cpu) ZeroMode() uint16 {
	// Use the value stored directly after the opcode as an index into memory and return the value stored there.
	address := c.readMemory(c.ProgramCounter + 1)
	return uint16(address)
}

func (c *Cpu) ZeroXMode() uint16 {
	// Calculate a memory address by adding the value stored directly after the opcode
	// add the value in the x register.

	// Address is a byte and the overflow/wrap behavior is intentional.
	address := c.readMemory(c.ProgramCounter + 1)
	address += c.RegX
	return uint16(address)
}

func (c *Cpu) AbsoluteMode() uint16 {
	// Use the two bytes stored directly after the opcode as an index into memory.
	// Treat them as litte endian (LSB first).

	address := c.readMemory_u16(c.ProgramCounter + 1)
	return address
}

func (c *Cpu) AbsoluteXMode() uint16 {
	// Same as AbsoluteMode but the value in the X register is added to
	// the memory address.

	address := c.readMemory_u16(c.ProgramCounter + 1)
	address += uint16(c.RegX)
	return address
}

func (c *Cpu) AbsoluteYMode() uint16 {
	// Same as AbsoluteMode but the value in the Y register is added to
	// the memory address.

	address := c.readMemory_u16(c.ProgramCounter + 1)
	address += uint16(c.RegY)
	return address
}

func (c *Cpu) IndirectXMode() uint16 {
	// Use the two bytes stored directly after the opcode as an index into memory.
	// Treat them as litte endian (LSB first). Add the value in the X register.
	// Use this sum as an initial index. Lookup the value stored in memory at
	// this index and return it.

	// Initial address is a byte and the overflow/wrap behavior is intentional.
	index := c.readMemory(c.ProgramCounter + 1)
	index += c.RegX

	// Use the initial address to read an address from memory.
	// Address is two bytes little endian (LSB first)
	address := c.readMemory_u16(uint16(index))
	return address
}

func (c *Cpu) IndirectYMode() uint16 {
	// Same as IndirectXMode but with the Y register.

	// Initial address is a byte and the overflow/wrap behavior is intentional.
	index := c.readMemory(c.ProgramCounter + 1)
	index += c.RegY

	// Use the initial address to read an address from memory.
	// Address is two bytes little endian (LSB first)
	address := c.readMemory_u16(uint16(index))
	return address
}

// https://skilldrick.github.io/easy6502/#stack
// Stack is 0x0100 to 0x01ff in memory.
// Stack pointer starts at 0xff refers to 0x01ff in memory.
// It grows downwards, so when a byte is added the next SP value is 0xfe.
// When adding addresses (such as JSR) the MSB is added first: 0x8000 -> 0x80 then 0x00
// TODO(mjpatter88): refactor stack management into helpers.
