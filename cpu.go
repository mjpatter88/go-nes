package main

import "fmt"

//Memory Addresses
const (
	// TODO(mjpatter88): revisit this once roms are supported.
	// Thid should be 0x8000, but that breaks since that address
	// will be part of the ROM address space.
	DEFAULT_PROG_MEM_ADDRESS   = 0x0200
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
	// TODO(mjpatter88): revist this. The cpu should probably just have a
	// pointer to the bus instead of owning it. Maybe introduce an "nes"
	// structure that owns the cpu, bus, ppu, etc?
	bus          Bus
	StackPointer uint8
}

func (c *Cpu) readMemory(index uint16) uint8 {
	return c.bus.ReadMemory(index)
}

func (c *Cpu) writeMemory(index uint16, value uint8) {
	c.bus.WriteMemory(index, value)
}

func (c *Cpu) Execute(program []uint8) {
	c.ExecuteAtAddress(program, DEFAULT_PROG_MEM_ADDRESS)
}

// Loads the program into memory at the specified address and executes it
func (c *Cpu) ExecuteAtAddress(program []uint8, address uint16) {
	c.LoadAtAddress(program, address)
	c.run()
}

func (c *Cpu) Load(program []uint8) {
	c.LoadAtAddress(program, DEFAULT_PROG_MEM_ADDRESS)
}

// Loads the program into memory at the specified address but does not execute it
func (c *Cpu) LoadAtAddress(program []uint8, address uint16) {
	for index, byte := range program {
		memIndex := address + uint16(index)
		c.bus.WriteMemory(memIndex, byte)
	}
	// nes spec says to write program memory address into mem address 0xFFFC
	// this value is then read into the program counter on system reset
	c.bus.WriteMemory_u16(PROG_REFERENCE_MEM_ADDRESS, address)
	c.reset()
}

func (c *Cpu) run() {
	for !c.Status.Break {
		c.Step()
	}
}

// Executes a single instruction.
// TODO(mjpatter88): maybe return the number of cycles?
func (c *Cpu) Step() {
	opcode := c.bus.ReadMemory(c.ProgramCounter)
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
		panic(fmt.Errorf("unsuppored addressing mode %b", instr.AddressingMode))
	case IMMEDIATE:
		param = c.ImmediateMode()
	case RELATIVE:
		param = c.RelativeMode()
	case INDIRECT:
		param = c.IndirectMode()
	case INDIRECT_X:
		param = c.IndirectXMode()
	case INDIRECT_Y:
		param = c.IndirectYMode()
	}

	switch instr.Action {
	case "BIT":
		c.instrBIT(param)
	case "LDA":
		c.instrLDA(param)
	case "LDX":
		c.instrLDX(param)
	case "LSR":
		if instr.AddressingMode == ACCUMULATOR {
			c.instrLSR_acc()
		} else {
			c.instrLSR(param)
		}
	case "LDY":
		c.instrLDY(param)
	case "AND":
		c.instrAND(param)
	case "ADC":
		c.instrADC(param)
	case "SBC":
		c.instrSBC(param)
	case "CMP":
		c.instrCMP(param)
	case "CPX":
		c.instrCPX(param)
	case "CPY":
		c.instrCPY(param)
	case "STA":
		c.instrSTA(param)
	case "TAX":
		c.instrTAX()
	case "TXA":
		c.instrTXA()
	case "TAY":
		c.instrTAY()
	case "TYA":
		c.instrTYA()
	case "DEX":
		c.instrDEX()
	case "INX":
		c.instrINX()
	case "DEY":
		c.instrDEY()
	case "INY":
		c.instrINY()
	case "INC":
		c.instrINC(param)
	case "DEC":
		c.instrDEC(param)
	case "JSR":
		c.instrJSR(param)
		didJump = true
	case "RTS":
		c.instrRTS(param)
		didJump = true
	case "BPL":
		didJump = c.instrBPL(param)
	case "BMI":
		didJump = c.instrBMI(param)
	case "BVC":
		didJump = c.instrBVC(param)
	case "BVS":
		didJump = c.instrBVS(param)
	case "BCC":
		didJump = c.instrBCC(param)
	case "BCS":
		didJump = c.instrBCS(param)
	case "BEQ":
		didJump = c.instrBEQ(param)
	case "BNE":
		didJump = c.instrBNE(param)
	case "JMP":
		c.instrJMP(param)
		didJump = true
	case "CLC":
		c.instrCLC()
	case "SEC":
		c.instrSEC()
	case "BRK":
		c.Status.Break = true
	case "NOP":
		// Intentionally empty
	default:
		panic(fmt.Errorf("unsuppored opcode %#x at pc: %#x", opcode, c.ProgramCounter))
	}

	// Jump instructions are expected to manually update the program counter themselves
	if !didJump {
		c.ProgramCounter += uint16(instr.NumberOfBytes)
	}

}

func (c *Cpu) PrintState() {
	fmt.Printf("Program Counter: %#x\n", c.ProgramCounter)
	fmt.Printf("Register A: %#x\n", c.RegA)
	fmt.Printf("Register X: %#x\n", c.RegX)
}

func (c *Cpu) reset() {
	c.resetStatus()
	c.RegA = 0
	c.RegX = 0
	c.ProgramCounter = c.bus.ReadMemory_u16(PROG_REFERENCE_MEM_ADDRESS)
	c.StackPointer = 0xff
}

func (c *Cpu) updateFlags(result uint8) {
	c.Status.Zero = (result == 0)
	c.Status.Negative = ((result & (1 << 7)) != 0)
}

func (c *Cpu) resetStatus() {
	c.Status = StatusRegister{}
}

// Compare the register value to another value.
// If  regVal == OtherValue then Zero reg = true
// If  regVal < OtherValue then Negative reg = true
// If  regVal > OtherValue then Carry reg = true
//
// See: http://6502.org/tutorials/compare_instructions.html and
// http://6502.org/tutorials/compare_beyond.html
func (c *Cpu) compare(regValue uint8, otherValue uint8) {
	if regValue == otherValue {
		c.Status.Zero = true
		c.Status.Carry = false
		c.Status.Negative = false
	}
	if regValue < otherValue {
		c.Status.Zero = false
		c.Status.Carry = false
		c.Status.Negative = true
	}
	if regValue > otherValue {
		c.Status.Zero = false
		c.Status.Carry = true
		c.Status.Negative = false
	}
}

func (c *Cpu) instrBIT(param uint16) {
	// Set the zero flag based on the result of a logical and of the value
	// loaded from memory and the accumulator register.
	//
	// Status flag handling here is complicated.
	// Zero is set based on the result of the operation, but
	// Negative and Overflow are set based on bits 7 and 6 of the operand respectively.
	//
	// See: https://www.masswerk.at/6502/6502_instruction_set.html#BIT

	value := c.bus.ReadMemory(param)
	result := value & c.RegA

	c.Status.Zero = (result == 0)
	c.Status.Negative = ((value & (1 << 7)) != 0)
	c.Status.Overflow = ((value & (1 << 6)) != 0)
}

func (c *Cpu) instrLDA(param uint16) {
	c.RegA = c.bus.ReadMemory(param)
	c.updateFlags(c.RegA)
}

func (c *Cpu) instrLDX(param uint16) {
	c.RegX = c.bus.ReadMemory(param)
	c.updateFlags(c.RegX)
}

func (c *Cpu) instrLDY(param uint16) {
	c.RegY = c.bus.ReadMemory(param)
	c.updateFlags(c.RegY)
}

func (c *Cpu) instrLSR(param uint16) {
	value := c.bus.ReadMemory(param)
	c.Status.Carry = (value & 0x01) != 0
	value = value >> 1
	c.bus.WriteMemory(param, value)
	c.updateFlags(value)
}

func (c *Cpu) instrLSR_acc() {
	c.Status.Carry = (c.RegA & 0x01) != 0
	c.RegA = c.RegA >> 1
	c.updateFlags(c.RegA)
}

func (c *Cpu) instrAND(param uint16) {
	value := c.bus.ReadMemory(param)
	c.RegA &= value
	c.updateFlags(c.RegA)
}

func (c *Cpu) instrINC(param uint16) {
	value := c.bus.ReadMemory(param)
	value += 1
	c.updateFlags(value)
	c.bus.WriteMemory(param, value)
}

func (c *Cpu) instrDEC(param uint16) {
	value := c.bus.ReadMemory(param)
	value -= 1
	c.updateFlags(value)
	c.bus.WriteMemory(param, value)
}

func (c *Cpu) instrADC(param uint16) {
	// TODO(mjpatter) handle overflow and carry flags correctly
	value := c.bus.ReadMemory(param)
	c.RegA += value
	c.updateFlags(c.RegA)
}

func (c *Cpu) instrSBC(param uint16) {
	// TODO(mjpatter) handle overflow and carry flags correctly
	value := c.bus.ReadMemory(param)
	c.RegA -= value
	c.updateFlags(c.RegA)
}

func (c *Cpu) instrCMP(param uint16) {
	value := c.bus.ReadMemory(param)
	c.compare(c.RegA, value)
}

func (c *Cpu) instrCPX(param uint16) {
	value := c.bus.ReadMemory(param)
	c.compare(c.RegX, value)
}

func (c *Cpu) instrCPY(param uint16) {
	value := c.bus.ReadMemory(param)
	c.compare(c.RegY, value)
}

func (c *Cpu) instrSTA(param uint16) {
	c.bus.WriteMemory(param, c.RegA)
	c.resetStatus()
}

func (c *Cpu) instrTAX() {
	c.RegX = c.RegA
	c.updateFlags(c.RegX)
}

func (c *Cpu) instrTXA() {
	c.RegA = c.RegX
	c.updateFlags(c.RegA)
}

func (c *Cpu) instrTAY() {
	c.RegY = c.RegA
	c.updateFlags(c.RegY)
}

func (c *Cpu) instrTYA() {
	c.RegA = c.RegY
	c.updateFlags(c.RegA)
}

func (c *Cpu) instrINX() {
	c.RegX++
	c.updateFlags(c.RegX)
}

func (c *Cpu) instrDEX() {
	c.RegX--
	c.updateFlags(c.RegX)
}

func (c *Cpu) instrINY() {
	c.RegY++
	c.updateFlags(c.RegY)
}

func (c *Cpu) instrDEY() {
	c.RegY--
	c.updateFlags(c.RegY)
}

func (c *Cpu) instrCLC() {
	c.Status.Carry = false
}

func (c *Cpu) instrSEC() {
	c.Status.Carry = true
}

func (c *Cpu) instrJSR(param uint16) {
	// We'll write the two bytes at once, so write it to SP - 1 (ends up writing to SP-1 and SP)
	index := 0x0100 | uint16((c.StackPointer - 1))
	// JSR length is 3 and we want to store the address of the next insturction - 1.
	value := c.ProgramCounter + 3 - 1
	c.bus.WriteMemory_u16(index, value)

	c.StackPointer -= 2
	c.ProgramCounter = param
}

func (c *Cpu) instrRTS(param uint16) {
	// Read two bytes from the top of the stack.
	index := 0x0100 | uint16((c.StackPointer + 1))
	value := c.bus.ReadMemory_u16(index)
	c.StackPointer += 2
	c.ProgramCounter = value + 1
}

// Returns true if branch was taken, false otherwise
func (c *Cpu) instrBPL(param uint16) bool {
	if !c.Status.Negative {
		c.ProgramCounter = param
		return true
	}
	return false
}

// Returns true if branch was taken, false otherwise
func (c *Cpu) instrBMI(param uint16) bool {
	if c.Status.Negative {
		c.ProgramCounter = param
		return true
	}
	return false
}

// Returns true if branch was taken, false otherwise
func (c *Cpu) instrBVC(param uint16) bool {
	if !c.Status.Overflow {
		c.ProgramCounter = param
		return true
	}
	return false
}

// Returns true if branch was taken, false otherwise
func (c *Cpu) instrBVS(param uint16) bool {
	if c.Status.Overflow {
		c.ProgramCounter = param
		return true
	}
	return false
}

// Returns true if branch was taken, false otherwise
func (c *Cpu) instrBCC(param uint16) bool {
	if !c.Status.Carry {
		c.ProgramCounter = param
		return true
	}
	return false
}

// Returns true if branch was taken, false otherwise
func (c *Cpu) instrBCS(param uint16) bool {
	if c.Status.Carry {
		c.ProgramCounter = param
		return true
	}
	return false
}

// Returns true if branch was taken, false otherwise
func (c *Cpu) instrBEQ(param uint16) bool {
	if c.Status.Zero {
		c.ProgramCounter = param
		return true
	}
	return false
}

// Returns true if branch was taken, false otherwise
func (c *Cpu) instrBNE(param uint16) bool {
	if !c.Status.Zero {
		c.ProgramCounter = param
		return true
	}
	return false
}

func (c *Cpu) instrJMP(param uint16) {
	c.ProgramCounter = param
}

// https://skilldrick.github.io/easy6502/#addressing
func (c *Cpu) ImmediateMode() uint16 {
	// Return the address of the value directly after the opcode.
	return c.ProgramCounter + 1
}

func (c *Cpu) ZeroMode() uint16 {
	// Use the value stored directly after the opcode as an index into memory and return the value stored there.
	address := c.bus.ReadMemory(c.ProgramCounter + 1)
	return uint16(address)
}

func (c *Cpu) ZeroXMode() uint16 {
	// Calculate a memory address by adding the value stored directly after the opcode
	// add the value in the x register.

	// Address is a byte and the overflow/wrap behavior is intentional.
	address := c.bus.ReadMemory(c.ProgramCounter + 1)
	address += c.RegX
	return uint16(address)
}

func (c *Cpu) AbsoluteMode() uint16 {
	// Use the two bytes stored directly after the opcode as an index into memory.
	// Treat them as litte endian (LSB first).

	address := c.bus.ReadMemory_u16(c.ProgramCounter + 1)
	return address
}

func (c *Cpu) AbsoluteXMode() uint16 {
	// Same as AbsoluteMode but the value in the X register is added to
	// the memory address.

	address := c.bus.ReadMemory_u16(c.ProgramCounter + 1)
	address += uint16(c.RegX)
	return address
}

func (c *Cpu) AbsoluteYMode() uint16 {
	// Same as AbsoluteMode but the value in the Y register is added to
	// the memory address.

	address := c.bus.ReadMemory_u16(c.ProgramCounter + 1)
	address += uint16(c.RegY)
	return address
}

func (c *Cpu) IndirectMode() uint16 {
	// Use the two bytes stored directly after the opcode as an index into memory.
	// Treat them as litte endian (LSB first). Lookup the value stored in memory at
	// this index and return it.

	// address is two bytes.
	address := c.bus.ReadMemory_u16(c.ProgramCounter + 1)

	// Use the address to read a value from memory.
	// Value is two bytes little endian (LSB first)
	value := c.bus.ReadMemory_u16(uint16(address))
	return value
}

func (c *Cpu) IndirectXMode() uint16 {
	// Use the byte stored directly after the opcode as an index into memory.
	// Add the value in the X register. Use this sum as an initial index.
	// Lookup the value stored in memory at this index and return it.

	// Initial address is a byte and the overflow/wrap behavior is intentional.
	index := c.bus.ReadMemory(c.ProgramCounter + 1)
	index += c.RegX

	// Use the initial address to read an address from memory.
	// Address is two bytes little endian (LSB first)
	address := c.bus.ReadMemory_u16(uint16(index))
	return address
}

func (c *Cpu) IndirectYMode() uint16 {
	// Same as IndirectXMode but with the Y register.

	// Initial address is a byte and the overflow/wrap behavior is intentional.
	index := c.bus.ReadMemory(c.ProgramCounter + 1)
	index += c.RegY

	// Use the initial address to read an address from memory.
	// Address is two bytes little endian (LSB first)
	address := c.bus.ReadMemory_u16(uint16(index))
	return address
}

func (c *Cpu) RelativeMode() uint16 {
	// Relative mode instructions are size 2.
	// Return pc + 2 + the offset.
	//
	// The offset can be positive or negative, so we need to use two's complement addition.
	// There may be a better way, but this series of casts does the trick.
	offset := c.bus.ReadMemory(c.ProgramCounter + 1)
	return uint16(int16(c.ProgramCounter)+int16(int8(offset))) + 2
}

// https://skilldrick.github.io/easy6502/#stack
// Stack is 0x0100 to 0x01ff in memory.
// Stack pointer starts at 0xff refers to 0x01ff in memory.
// It grows downwards, so when a byte is added the next SP value is 0xfe.
// When adding addresses (such as JSR) the MSB is added first: 0x8000 -> 0x80 then 0x00
// TODO(mjpatter88): refactor stack management into helpers.
