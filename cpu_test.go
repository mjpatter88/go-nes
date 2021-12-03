package main

import (
	"testing"
)

func TestBRK(t *testing.T) {
	cpu := Cpu{}
	cpu.Execute([]uint8{BRK})

	AssertBreak(t, &cpu, true)
}

func TestNOP(t *testing.T) {
	cpu := Cpu{}
	cpu.Execute([]uint8{NOP, BRK})

	// No assertions.
}

func TestBIT(t *testing.T) {
	t.Run("BIT Zero", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x00
		cpu.bus.cpuVRam[0xaa] = 0xff
		cpu.instrBIT(0xaa)

		AssertZero(t, &cpu, true)
	})

	t.Run("BIT Non-Zero", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x4a
		cpu.bus.cpuVRam[0xaa] = 0x4a
		cpu.instrBIT(0xaa)

		AssertZero(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		// Negative flag is set to the value of bit 7 of the operand
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0x80
		cpu.instrBIT(0xaa)

		AssertNegative(t, &cpu, true)
	})

	t.Run("Overflow flag", func(t *testing.T) {
		// Overflow flag is set to the value of bit 6 of the operand
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0x40
		cpu.instrBIT(0xaa)

		AssertOverflow(t, &cpu, true)
	})

	t.Run("BIT Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0xff
		cpu.Execute([]uint8{BIT_ZERO, 0xaa, BRK})

		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, true)
		AssertOverflow(t, &cpu, true)
	})
}

func TestLDA(t *testing.T) {
	t.Run("LDA", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0x4a
		cpu.instrLDA(0xaa)

		AssertRegisterA(t, &cpu, 0x4a)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0x00
		cpu.instrLDA(0xaa)

		AssertRegisterA(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0xf0
		cpu.instrLDA(0xaa)

		AssertRegisterA(t, &cpu, 0xf0)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})

	t.Run("LDA Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDA, 0x0e, BRK})

		AssertRegisterA(t, &cpu, 0x0e)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})
}

func TestLDX(t *testing.T) {
	t.Run("LDX", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0x4a
		cpu.instrLDX(0xaa)

		AssertRegisterX(t, &cpu, 0x4a)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0x00
		cpu.instrLDX(0xaa)

		AssertRegisterX(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0xf0
		cpu.instrLDX(0xaa)

		AssertRegisterX(t, &cpu, 0xf0)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})

	t.Run("LDX Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDX, 0x0e, BRK})

		AssertRegisterX(t, &cpu, 0x0e)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})
}

func TestLDY(t *testing.T) {
	t.Run("LDY", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0x4a
		cpu.instrLDY(0xaa)

		AssertRegisterY(t, &cpu, 0x4a)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0x00
		cpu.instrLDY(0xaa)

		AssertRegisterY(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0xf0
		cpu.instrLDY(0xaa)

		AssertRegisterY(t, &cpu, 0xf0)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})

	t.Run("LDY Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDY, 0x0e, BRK})

		AssertRegisterY(t, &cpu, 0x0e)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})
}

func TestLSR(t *testing.T) {
	t.Run("LSR", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0xff
		cpu.instrLSR_acc()

		AssertRegisterA(t, &cpu, 0x7f)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
		AssertCarry(t, &cpu, true)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x01
		cpu.instrLSR_acc()

		AssertRegisterA(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
		AssertCarry(t, &cpu, true)
	})

	t.Run("LSR Zero Page", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0xff
		cpu.instrLSR(0xaa)

		AssertMemoryValue(t, &cpu, 0xaa, 0x7f)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
		AssertCarry(t, &cpu, true)
	})

	t.Run("LSR Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDA, 0xff, LSR, BRK})

		AssertRegisterA(t, &cpu, 0x7f)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})

	t.Run("LSR Zero Page Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0xff
		cpu.Execute([]uint8{LSR_ZERO, 0xaa, BRK})

		AssertMemoryValue(t, &cpu, 0xaa, 0x7f)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
		AssertCarry(t, &cpu, true)
	})
}

func TestINC(t *testing.T) {
	t.Run("INC", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0xf1
		cpu.instrINC(0xaa)

		AssertMemoryValue(t, &cpu, 0xaa, 0xf2)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0xff
		cpu.instrINC(0xaa)

		AssertMemoryValue(t, &cpu, 0xaa, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0x7f
		cpu.instrINC(0xaa)

		AssertMemoryValue(t, &cpu, 0xaa, 0x80)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})

	t.Run("INC Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0xf1
		cpu.Execute([]uint8{INC_ZERO, 0xaa, BRK})

		AssertMemoryValue(t, &cpu, 0xaa, 0xf2)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})
}

func TestDEC(t *testing.T) {
	t.Run("DEC", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0xf1
		cpu.instrDEC(0xaa)

		AssertMemoryValue(t, &cpu, 0xaa, 0xf0)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0x01
		cpu.instrDEC(0xaa)

		AssertMemoryValue(t, &cpu, 0xaa, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0x00
		cpu.instrDEC(0xaa)

		AssertMemoryValue(t, &cpu, 0xaa, 0xFF)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})

	t.Run("DEC Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0xaa] = 0xf1
		cpu.Execute([]uint8{DEC_ZERO, 0xaa, BRK})

		AssertMemoryValue(t, &cpu, 0xaa, 0xf0)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})
}

func TestAND(t *testing.T) {
	t.Run("AND", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x0f
		cpu.bus.cpuVRam[0xaa] = 0xff
		cpu.instrAND(0xaa)

		AssertRegisterA(t, &cpu, 0x0f)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x0f
		cpu.bus.cpuVRam[0xaa] = 0x00
		cpu.instrAND(0xaa)

		AssertRegisterA(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0xff
		cpu.bus.cpuVRam[0xaa] = 0xff
		cpu.instrAND(0xaa)

		AssertRegisterA(t, &cpu, 0xff)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})

	t.Run("AND Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDA, 0xf0, AND, 0xff, BRK})

		AssertRegisterA(t, &cpu, 0xf0)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})
}

func TestADC(t *testing.T) {
	t.Run("ADC", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x02
		cpu.bus.cpuVRam[0xaa] = 0x31
		cpu.instrADC(0xaa)

		AssertRegisterA(t, &cpu, 0x33)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
		AssertOverflow(t, &cpu, false)
		AssertCarry(t, &cpu, false)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x00
		cpu.bus.cpuVRam[0xaa] = 0x00
		cpu.instrADC(0xaa)

		AssertRegisterA(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0xfa
		cpu.bus.cpuVRam[0xaa] = 0x01
		cpu.instrADC(0xaa)

		AssertRegisterA(t, &cpu, 0xfb)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})

	t.Run("Overflow", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0xff
		cpu.bus.cpuVRam[0xaa] = 0x01
		cpu.instrADC(0xaa)

		AssertRegisterA(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("ADC Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDA, 0xf0, ADC, 0x1, BRK})

		AssertRegisterA(t, &cpu, 0xf1)
	})
}

func TestSBC(t *testing.T) {
	t.Run("SBC", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x32
		cpu.bus.cpuVRam[0xaa] = 0x02
		cpu.instrSBC(0xaa)

		AssertRegisterA(t, &cpu, 0x30)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
		AssertOverflow(t, &cpu, false)
		AssertCarry(t, &cpu, false)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x01
		cpu.bus.cpuVRam[0xaa] = 0x01
		cpu.instrSBC(0xaa)

		AssertRegisterA(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x01
		cpu.bus.cpuVRam[0xaa] = 0x02
		cpu.instrSBC(0xaa)

		AssertRegisterA(t, &cpu, 0xff)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})

	t.Run("SBC Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDA, 0x88, SBC, 0x1, BRK})

		AssertRegisterA(t, &cpu, 0x87)
	})
}

func TestCMP(t *testing.T) {
	t.Run("CMP - equal", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x4a
		cpu.bus.cpuVRam[0xaa] = 0x4a
		cpu.instrCMP(0xaa)

		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
		AssertCarry(t, &cpu, false)
	})

	t.Run("CMP - regA less than", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x0a
		cpu.bus.cpuVRam[0xaa] = 0x4a
		cpu.instrCMP(0xaa)

		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
		AssertCarry(t, &cpu, false)
	})

	t.Run("CMP - regA greater than", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x5a
		cpu.bus.cpuVRam[0xaa] = 0x4a
		cpu.instrCMP(0xaa)

		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
		AssertCarry(t, &cpu, true)
	})

	t.Run("CMP Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDA, 0x0e, CMP, 0x0e, BRK})

		AssertRegisterA(t, &cpu, 0x0e)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
		AssertCarry(t, &cpu, false)
	})
}

func TestCPX(t *testing.T) {
	t.Run("CPX - equal", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegX = 0x4a
		cpu.bus.cpuVRam[0xaa] = 0x4a
		cpu.instrCPX(0xaa)

		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
		AssertCarry(t, &cpu, false)
	})

	t.Run("CPX - regX less than", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegX = 0x0a
		cpu.bus.cpuVRam[0xaa] = 0x4a
		cpu.instrCPX(0xaa)

		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
		AssertCarry(t, &cpu, false)
	})

	t.Run("CPX - regX greater than", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegX = 0x5a
		cpu.bus.cpuVRam[0xaa] = 0x4a
		cpu.instrCPX(0xaa)

		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
		AssertCarry(t, &cpu, true)
	})

	t.Run("CPX Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDA, 0x0e, TAX, CPX, 0x0e, BRK})

		AssertRegisterX(t, &cpu, 0x0e)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
		AssertCarry(t, &cpu, false)
	})
}

func TestCPY(t *testing.T) {
	t.Run("CPY - equal", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegY = 0x4a
		cpu.bus.cpuVRam[0xaa] = 0x4a
		cpu.instrCPY(0xaa)

		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
		AssertCarry(t, &cpu, false)
	})

	t.Run("CPY - regY less than", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegY = 0x0a
		cpu.bus.cpuVRam[0xaa] = 0x4a
		cpu.instrCPY(0xaa)

		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
		AssertCarry(t, &cpu, false)
	})

	t.Run("CPY - regY greater than", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegY = 0x5a
		cpu.bus.cpuVRam[0xaa] = 0x4a
		cpu.instrCPY(0xaa)

		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
		AssertCarry(t, &cpu, true)
	})

	t.Run("CPY Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDA, 0x0e, TAY, CPY, 0x0e, BRK})

		AssertRegisterY(t, &cpu, 0x0e)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
		AssertCarry(t, &cpu, false)
	})
}

func TestSTA(t *testing.T) {
	t.Run("STA", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0xff
		cpu.instrSTA(0x01)

		AssertMemoryValue(t, &cpu, 0x01, 0xff)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})

	t.Run("STA Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDA, 0xfe, STA_ZERO, 0x01, BRK})
		AssertMemoryValue(t, &cpu, 0x01, 0xfe)
	})
}

func TestTAX(t *testing.T) {
	t.Run("TAX", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x3a
		cpu.instrTAX()

		AssertRegisterX(t, &cpu, 0x3a)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x00
		cpu.instrTAX()

		AssertRegisterX(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0xf0
		cpu.instrTAX()

		AssertRegisterX(t, &cpu, 0xf0)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})

	t.Run("TAX Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDA, 0x0e, TAX, BRK})

		AssertRegisterX(t, &cpu, 0x0e)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})
}

func TestTXA(t *testing.T) {
	t.Run("TXA", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegX = 0x3a
		cpu.instrTXA()

		AssertRegisterA(t, &cpu, 0x3a)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegX = 0x00
		cpu.instrTXA()

		AssertRegisterA(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegX = 0xf0
		cpu.instrTXA()

		AssertRegisterA(t, &cpu, 0xf0)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})

	t.Run("TXA Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDX, 0x0e, TXA, BRK})

		AssertRegisterA(t, &cpu, 0x0e)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})
}

func TestTAY(t *testing.T) {
	t.Run("TAY", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x3a
		cpu.instrTAY()

		AssertRegisterY(t, &cpu, 0x3a)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x00
		cpu.instrTAY()

		AssertRegisterY(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0xf0
		cpu.instrTAY()

		AssertRegisterY(t, &cpu, 0xf0)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})

	t.Run("TAY Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDA, 0x0e, TAY, BRK})

		AssertRegisterY(t, &cpu, 0x0e)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})
}

func TestTYA(t *testing.T) {
	t.Run("TYA", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegY = 0x3a
		cpu.instrTYA()

		AssertRegisterA(t, &cpu, 0x3a)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegY = 0x00
		cpu.instrTYA()

		AssertRegisterA(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegY = 0xf0
		cpu.instrTYA()

		AssertRegisterA(t, &cpu, 0xf0)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})

	t.Run("TYA Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDY, 0x0e, TYA, BRK})

		AssertRegisterA(t, &cpu, 0x0e)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})
}

func TestINX(t *testing.T) {
	t.Run("INX", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegX = 0x3a
		cpu.instrINX()

		AssertRegisterX(t, &cpu, 0x3b)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegX = 0xff
		cpu.instrINX()

		AssertRegisterX(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegX = 0x7f
		cpu.instrINX()

		AssertRegisterX(t, &cpu, 0x80)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})

	t.Run("INX Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDA, 0x0e, TAX, INX, BRK})

		AssertRegisterX(t, &cpu, 0x0f)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})
}

func TestDEX(t *testing.T) {
	t.Run("DEX", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegX = 0x3a
		cpu.instrDEX()

		AssertRegisterX(t, &cpu, 0x39)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegX = 0x01
		cpu.instrDEX()

		AssertRegisterX(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegX = 0xff
		cpu.instrDEX()

		AssertRegisterX(t, &cpu, 0xfe)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})

	t.Run("DEX Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDA, 0x0e, TAX, DEX, BRK})

		AssertRegisterX(t, &cpu, 0x0d)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})
}

func TestINY(t *testing.T) {
	t.Run("INY", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegY = 0x3a
		cpu.instrINY()

		AssertRegisterY(t, &cpu, 0x3b)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegY = 0xff
		cpu.instrINY()

		AssertRegisterY(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegY = 0x7f
		cpu.instrINY()

		AssertRegisterY(t, &cpu, 0x80)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})

	t.Run("INY Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDA, 0x0e, TAY, INY, BRK})

		AssertRegisterY(t, &cpu, 0x0f)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})
}

func TestDEY(t *testing.T) {
	t.Run("DEY", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegY = 0x3a
		cpu.instrDEY()

		AssertRegisterY(t, &cpu, 0x39)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegY = 0x01
		cpu.instrDEY()

		AssertRegisterY(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegY = 0xff
		cpu.instrDEY()

		AssertRegisterY(t, &cpu, 0xfe)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})

	t.Run("DEY Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDA, 0x0e, TAY, DEY, BRK})

		AssertRegisterY(t, &cpu, 0x0d)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})
}

func TestCLC(t *testing.T) {
	t.Run("CLC", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Status.Carry = true
		cpu.instrCLC()

		AssertCarry(t, &cpu, false)
	})

	t.Run("CLC Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{CLC, BRK})
		AssertCarry(t, &cpu, false)
	})
}

func TestSEC(t *testing.T) {
	t.Run("SEC", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Status.Carry = true
		cpu.instrSEC()

		AssertCarry(t, &cpu, true)
	})

	t.Run("SEC Instruction", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{SEC, BRK})
		AssertCarry(t, &cpu, true)
	})
}

func TestJSR(t *testing.T) {
	cpu := Cpu{}
	cpu.ProgramCounter = 0x0200
	cpu.StackPointer = 0x00ff
	cpu.instrJSR(0x9000)

	AssertProgramCounter(t, &cpu, 0x9000)
	// 0x0200 + 3 (len of current instr) - 1 => 0x0202
	// MSB goes onto stack first then LSB second
	AssertMemoryValue(t, &cpu, 0x01ff, 0x02)
	AssertMemoryValue(t, &cpu, 0x01fe, 0x02)

	AssertStackPointer(t, &cpu, 0x00fd)
}

// Exercise sets of instructions that utilize various addressing modes
func TestAddressingModeInstructionExecution(t *testing.T) {
	t.Run("Zero Page", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0x0000] = 0xee
		cpu.Execute([]uint8{LDA_ZERO, 0x00, BRK})
		AssertRegisterA(t, &cpu, 0xee)
	})

	t.Run("Zero Page X", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0x00ff] = 0xee
		// Load 0xfe into a, transfer to x, the load from (0xfe + 1) into a
		cpu.Execute([]uint8{LDA, 0xfe, TAX, LDA_ZERO_X, 0x01, BRK})
		AssertRegisterA(t, &cpu, 0xee)
	})

	t.Run("Zero Page X - Address Overflow", func(t *testing.T) {
		// If the summed address overflows one byte, then it should wrap around.
		// Ex: 0xff + 0x05 -> 0x04
		cpu := Cpu{}
		cpu.bus.cpuVRam[0x0004] = 0xee
		cpu.Execute([]uint8{LDA, 0xff, TAX, LDA_ZERO_X, 0x05, BRK})
		AssertRegisterA(t, &cpu, 0xee)
	})

	t.Run("Absolute", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0x01aa] = 0xee
		// Remember little-endian applies to the absolute address
		cpu.Execute([]uint8{LDA_ABS, 0xaa, 0x01, BRK})
		AssertRegisterA(t, &cpu, 0xee)
	})

	t.Run("Absolute X", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0x01ab] = 0xee
		// Remember little-endian applies to the absolute address
		cpu.Execute([]uint8{LDA, 0x01, TAX, LDA_ABS_X, 0xaa, 0x01, BRK})
		AssertRegisterA(t, &cpu, 0xee)
	})

	t.Run("Absolute Y", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0x01ab] = 0xee
		// Remember little-endian applies to the absolute address
		cpu.Execute([]uint8{LDA, 0x01, TAY, LDA_ABS_Y, 0xaa, 0x01, BRK})
		AssertRegisterA(t, &cpu, 0xee)
	})

	t.Run("Indirect X", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0x00fe] = 0xaa
		cpu.bus.cpuVRam[0x00ff] = 0x01
		cpu.bus.cpuVRam[0x01aa] = 0xee
		cpu.Execute([]uint8{LDA, 0x01, TAX, LDA_IND_X, 0xfd, BRK})
		AssertRegisterA(t, &cpu, 0xee)
	})

	t.Run("Indirect Y", func(t *testing.T) {
		cpu := Cpu{}
		cpu.bus.cpuVRam[0x00fe] = 0xaa
		cpu.bus.cpuVRam[0x00ff] = 0x01
		cpu.bus.cpuVRam[0x01aa] = 0xee
		cpu.Execute([]uint8{LDA, 0x01, TAY, LDA_IND_Y, 0xfd, BRK})
		AssertRegisterA(t, &cpu, 0xee)
	})
}

// Excercise sets of jumping instructions to test overall behavior
func TestJumpingInstructionExecution(t *testing.T) {
	t.Run("JSR", func(t *testing.T) {
		cpu := Cpu{}
		// Jump to the INX instruction
		// This assumes a starting address of 0x0200 so the address to jump to is 0x0205
		cpu.Execute([]uint8{JSR, 0x05, 0x02, LDA, 0x05, INX, BRK})
		AssertRegisterA(t, &cpu, 0x00)
		AssertRegisterX(t, &cpu, 0x01)
	})

	t.Run("JSR and RTS", func(t *testing.T) {
		cpu := Cpu{}
		// Jump to the INX instruction and then return.
		// Expect one INX to run and the LDA instr to run but nothing else.
		cpu.Execute([]uint8{JSR, 0x06, 0x02, LDA, 0x05, BRK, INX, RTS, INX, BRK})
		AssertRegisterA(t, &cpu, 0x05)
		AssertRegisterX(t, &cpu, 0x01)
	})
}

func TestBMI(t *testing.T) {
	t.Run("BMI", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x8000
		cpu.Status.Negative = true
		branchTaken := cpu.instrBMI(0x1234)

		if !branchTaken {
			t.Fatalf("Expected to take branch but did not.")
		}
		AssertProgramCounter(t, &cpu, 0x1234)
	})

	t.Run("BMI - No branch", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x8000
		cpu.Status.Negative = false
		branchTaken := cpu.instrBMI(0x1234)

		if branchTaken {
			t.Fatalf("Expected to not take branch but did.")
		}
		AssertProgramCounter(t, &cpu, 0x8000)
	})
}

func TestBPL(t *testing.T) {
	t.Run("BPL", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x8000
		cpu.Status.Negative = false
		branchTaken := cpu.instrBPL(0x1234)

		if !branchTaken {
			t.Fatalf("Expected to take branch but did not.")
		}
		AssertProgramCounter(t, &cpu, 0x1234)
	})

	t.Run("BPL - No branch", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x8000
		cpu.Status.Negative = true
		branchTaken := cpu.instrBPL(0x1234)

		if branchTaken {
			t.Fatalf("Expected to not take branch but did.")
		}
		AssertProgramCounter(t, &cpu, 0x8000)
	})
}

func TestBEQ(t *testing.T) {
	t.Run("BEQ", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x8000
		cpu.Status.Zero = true
		branchTaken := cpu.instrBEQ(0x1234)

		if !branchTaken {
			t.Fatalf("Expected to take branch but did not.")
		}
		AssertProgramCounter(t, &cpu, 0x1234)
	})

	t.Run("BEQ - No branch", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x8000
		cpu.Status.Zero = false
		branchTaken := cpu.instrBEQ(0x1234)

		if branchTaken {
			t.Fatalf("Expected to not take branch but did.")
		}
		AssertProgramCounter(t, &cpu, 0x8000)
	})
}

func TestBNE(t *testing.T) {
	t.Run("BNE", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x8000
		cpu.Status.Zero = false
		branchTaken := cpu.instrBNE(0x1234)

		if !branchTaken {
			t.Fatalf("Expected to take branch but did not.")
		}
		AssertProgramCounter(t, &cpu, 0x1234)
	})
	t.Run("BNE - No branch", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x8000
		cpu.Status.Zero = true
		branchTaken := cpu.instrBNE(0x1234)

		if branchTaken {
			t.Fatalf("Expected to not take branch but did.")
		}
		AssertProgramCounter(t, &cpu, 0x8000)
	})

}

func TestBVC(t *testing.T) {
	t.Run("BVC", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x8000
		cpu.Status.Overflow = false
		branchTaken := cpu.instrBVC(0x1234)

		if !branchTaken {
			t.Fatalf("Expected to take branch but did not.")
		}
		AssertProgramCounter(t, &cpu, 0x1234)
	})
	t.Run("BVC - No branch", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x8000
		cpu.Status.Overflow = true
		branchTaken := cpu.instrBVC(0x1234)

		if branchTaken {
			t.Fatalf("Expected to not take branch but did.")
		}
		AssertProgramCounter(t, &cpu, 0x8000)
	})

}

func TestBVS(t *testing.T) {
	t.Run("BVS", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x8000
		cpu.Status.Overflow = true
		branchTaken := cpu.instrBVS(0x1234)

		if !branchTaken {
			t.Fatalf("Expected to take branch but did not.")
		}
		AssertProgramCounter(t, &cpu, 0x1234)
	})
	t.Run("BVS - No branch", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x8000
		cpu.Status.Overflow = false
		branchTaken := cpu.instrBVS(0x1234)

		if branchTaken {
			t.Fatalf("Expected to not take branch but did.")
		}
		AssertProgramCounter(t, &cpu, 0x8000)
	})

}

func TestBCC(t *testing.T) {
	t.Run("BCC", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x8000
		cpu.Status.Carry = false
		branchTaken := cpu.instrBCC(0x1234)

		if !branchTaken {
			t.Fatalf("Expected to take branch but did not.")
		}
		AssertProgramCounter(t, &cpu, 0x1234)
	})
	t.Run("BCC - No branch", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x8000
		cpu.Status.Carry = true
		branchTaken := cpu.instrBCC(0x1234)

		if branchTaken {
			t.Fatalf("Expected to not take branch but did.")
		}
		AssertProgramCounter(t, &cpu, 0x8000)
	})
}

func TestBCS(t *testing.T) {
	t.Run("BCS", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x8000
		cpu.Status.Carry = true
		branchTaken := cpu.instrBCS(0x1234)

		if !branchTaken {
			t.Fatalf("Expected to take branch but did not.")
		}
		AssertProgramCounter(t, &cpu, 0x1234)
	})
	t.Run("BCS - No branch", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x8000
		cpu.Status.Carry = false
		branchTaken := cpu.instrBCS(0x1234)

		if branchTaken {
			t.Fatalf("Expected to not take branch but did.")
		}
		AssertProgramCounter(t, &cpu, 0x8000)
	})
}

func TestJMP(t *testing.T) {
	t.Run("JMP", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x8000
		cpu.instrJMP(0x1234)

		AssertProgramCounter(t, &cpu, 0x1234)
	})
	t.Run("JMP Instruction - Absolute", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{JMP_ABS, 0x34, 0x02, BRK})
		// Assert jump target +1 since it will execute the following instruction which is 0x00 (BRK)
		AssertProgramCounter(t, &cpu, 0x0235)
	})

	t.Run("JMP Instruction - Indirect", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{JMP_IND, 0x03, 0x02, 0x34, 0x02, BRK})
		// Assert jump target +1 since it will execute the following instruction which is 0x00 (BRK)
		AssertProgramCounter(t, &cpu, 0x0235)
	})
}

// Test combinations of compare instructions with branching instructions
func TestCompareAndBranch(t *testing.T) {
	t.Run("CMP + BEQ", func(t *testing.T) {
		cpu := Cpu{}
		// Branch over the INX instruction.
		cpu.Execute([]uint8{LDA, 0x05, CMP, 0x05, BEQ, 0x01, INX, BRK})
		AssertRegisterX(t, &cpu, 0x00)

		AssertProgramCounter(t, &cpu, 0x0208)
	})
	t.Run("CMP + BEQ - No Branch", func(t *testing.T) {
		cpu := Cpu{}
		// Don't branch over the INX instruction.
		cpu.Execute([]uint8{LDA, 0x05, CMP, 0x06, BEQ, 0x01, INX, BRK})
		AssertRegisterX(t, &cpu, 0x01)

		AssertProgramCounter(t, &cpu, 0x0208)
	})
	t.Run("CMP + BNE", func(t *testing.T) {
		cpu := Cpu{}
		// Branch over the INX instruction.
		cpu.Execute([]uint8{LDA, 0x05, CMP, 0x06, BNE, 0x01, INX, BRK})
		AssertRegisterX(t, &cpu, 0x00)

		AssertProgramCounter(t, &cpu, 0x0208)
	})
	t.Run("CMP + BNE - No Branch", func(t *testing.T) {
		cpu := Cpu{}
		// Don't branch over the INX instruction.
		cpu.Execute([]uint8{LDA, 0x05, CMP, 0x05, BNE, 0x01, INX, BRK})
		AssertRegisterX(t, &cpu, 0x01)

		AssertProgramCounter(t, &cpu, 0x0208)
	})

	t.Run("BPL loop", func(t *testing.T) {
		cpu := Cpu{}
		// Decrement X until it is no longer positive
		// Branch back -3 (1 for dex and 2 for bpl)
		cpu.Execute([]uint8{LDX, 0x05, DEX, BPL, 0xfd, BRK})
		AssertRegisterX(t, &cpu, 0xff)

		AssertProgramCounter(t, &cpu, 0x0206)
	})

	t.Run("BMI loop", func(t *testing.T) {
		cpu := Cpu{}
		// Increment X until it is no longer negative
		// Branch back -3 (1 for dex and 2 for bmi)
		cpu.Execute([]uint8{LDX, 0xfd, INX, BMI, 0xfd, BRK})
		AssertRegisterX(t, &cpu, 0x00)

		AssertProgramCounter(t, &cpu, 0x0206)
	})
}

func TestFiveInstructions(t *testing.T) {
	cpu := Cpu{}
	cpu.Execute([]uint8{LDA, 0xc0, TAX, INX, BRK})

	AssertRegisterA(t, &cpu, 0xc0)
	AssertRegisterX(t, &cpu, 0xc1)
	AssertProgramCounter(t, &cpu, 0x0205)
}

func TestLoad(t *testing.T) {
	t.Run("Load Program", func(t *testing.T) {
		cpu := Cpu{}
		programBytes := []uint8{0x01, 0x02, 0x03}
		var startingAddress uint16 = 0x0200
		cpu.LoadAtAddress(programBytes, startingAddress)

		for i := 0; i < 3; i++ {
			memAddress := startingAddress + uint16(i)
			if cpu.bus.cpuVRam[memAddress] != uint8(i+1) {
				t.Errorf("Expected memory[%#x] to be %#x but was %#x", memAddress, i, cpu.bus.cpuVRam[memAddress])
			}
		}
	})
	t.Run("Sets Program Reference in Memory", func(t *testing.T) {
		cpu := Cpu{}
		programBytes := []uint8{0x01, 0x02, 0x03}
		cpu.LoadAtAddress(programBytes, 0x0200)

		value := cpu.bus.ReadMemory_u16(0xfffc)
		if value != 0x0200 {
			t.Errorf("Expected memory[0xfffc] to be %#x but was %#x", 0x0200, value)
		}
	})
}

func TestReset(t *testing.T) {
	cpu := Cpu{}
	cpu.Status.Zero = true
	cpu.RegA = 0x11
	cpu.RegX = 0x22
	cpu.bus.startingPC = 0x1234

	cpu.reset()

	AssertZero(t, &cpu, false)
	AssertRegisterX(t, &cpu, 0)
	AssertRegisterA(t, &cpu, 0)
	AssertProgramCounter(t, &cpu, 0x1234)
	AssertStackPointer(t, &cpu, 0xff)
}

func TestImmediateMode(t *testing.T) {
	cpu := Cpu{}
	cpu.ProgramCounter = 0x01
	value := cpu.ImmediateMode()

	if value != 0x02 {
		t.Errorf("Expected %#x but got %#x", 0x02, value)
	}
}

func TestZeroMode(t *testing.T) {
	cpu := Cpu{}
	cpu.ProgramCounter = 0x01
	cpu.bus.cpuVRam[0x02] = 0x05
	value := cpu.ZeroMode()

	if value != 0x05 {
		t.Errorf("Expected %#x but got %#x", 0x05, value)
	}
}

func TestZeroXMode(t *testing.T) {
	cpu := Cpu{}
	cpu.ProgramCounter = 0x01
	cpu.RegX = 0x01
	cpu.bus.cpuVRam[0x02] = 0x05
	value := cpu.ZeroXMode()

	if value != 0x06 {
		t.Errorf("Expected %#x but got %#x", 0x06, value)
	}
}

func TestAbsoluteMode(t *testing.T) {
	cpu := Cpu{}
	cpu.ProgramCounter = 0x01
	cpu.bus.cpuVRam[0x02] = 0x34
	cpu.bus.cpuVRam[0x03] = 0x12
	value := cpu.AbsoluteMode()

	if value != 0x1234 {
		t.Errorf("Expected %#x but got %#x", 0x1234, value)
	}
}

func TestAbsoluteXMode(t *testing.T) {
	cpu := Cpu{}
	cpu.ProgramCounter = 0x01
	cpu.RegX = 0x01
	cpu.bus.cpuVRam[0x02] = 0x34
	cpu.bus.cpuVRam[0x03] = 0x12
	value := cpu.AbsoluteXMode()

	if value != 0x1235 {
		t.Errorf("Expected %#x but got %#x", 0x1235, value)
	}
}

func TestAbsoluteYMode(t *testing.T) {
	cpu := Cpu{}
	cpu.ProgramCounter = 0x01
	cpu.RegY = 0x01
	cpu.bus.cpuVRam[0x02] = 0x34
	cpu.bus.cpuVRam[0x03] = 0x12
	value := cpu.AbsoluteYMode()

	if value != 0x1235 {
		t.Errorf("Expected %#x but got %#x", 0x1235, value)
	}
}
func TestIndirectMode(t *testing.T) {
	cpu := Cpu{}
	cpu.ProgramCounter = 0x01
	cpu.bus.cpuVRam[0x02] = 0xf0
	cpu.bus.cpuVRam[0x03] = 0x00

	cpu.bus.cpuVRam[0xf0] = 0x01
	cpu.bus.cpuVRam[0xf1] = 0xcc
	value := cpu.IndirectMode()

	if value != 0xcc01 {
		t.Errorf("Expected %#x but got %#x", 0xcc01, value)
	}
}

func TestIndirectXMode(t *testing.T) {
	cpu := Cpu{}
	cpu.ProgramCounter = 0x01
	cpu.RegX = 0x01
	cpu.bus.cpuVRam[0x02] = 0x34
	cpu.bus.cpuVRam[0x35] = 0xab
	value := cpu.IndirectXMode()

	if value != 0xab {
		t.Errorf("Expected %#x but got %#x", 0xab, value)
	}
}

func TestIndirectYMode(t *testing.T) {
	cpu := Cpu{}
	cpu.ProgramCounter = 0x01
	cpu.RegY = 0x01
	cpu.bus.cpuVRam[0x02] = 0x34
	cpu.bus.cpuVRam[0x35] = 0xab
	value := cpu.IndirectYMode()

	if value != 0xab {
		t.Errorf("Expected %#x but got %#x", 0xab, value)
	}
}

func TestRelativeMode(t *testing.T) {
	t.Run("Relative Mode - positive offset", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x01
		cpu.bus.cpuVRam[0x02] = 0x34
		value := cpu.RelativeMode()

		if value != 0x37 {
			t.Errorf("Expected %#x but got %#x", 0x37, value)
		}
	})
	t.Run("Relative Mode - negative offset", func(t *testing.T) {
		cpu := Cpu{}
		cpu.ProgramCounter = 0x05
		cpu.bus.cpuVRam[0x06] = 0xfd
		value := cpu.RelativeMode()

		// 5 + 2 - 3 = 4
		if value != 0x04 {
			t.Errorf("Expected %#x but got %#x", 0x04, value)
		}
	})
}

// Test helpers
func AssertBreak(t *testing.T, cpu *Cpu, status bool) {
	if cpu.Status.Break != status {
		t.Errorf("Expected Break status to be %t but was %t", status, cpu.Status.Break)
	}
}
func AssertZero(t *testing.T, cpu *Cpu, status bool) {
	if cpu.Status.Zero != status {
		t.Errorf("Expected Zero status to be %t but was %t", status, cpu.Status.Zero)
	}
}
func AssertCarry(t *testing.T, cpu *Cpu, status bool) {
	if cpu.Status.Carry != status {
		t.Errorf("Expected Carry status to be %t but was %t", status, cpu.Status.Carry)
	}
}
func AssertNegative(t *testing.T, cpu *Cpu, status bool) {
	if cpu.Status.Negative != status {
		t.Errorf("Expected Negative status to be %t but was %t", status, cpu.Status.Negative)
	}
}

func AssertOverflow(t *testing.T, cpu *Cpu, status bool) {
	if cpu.Status.Overflow != status {
		t.Errorf("Expected Overflow status to be %t but was %t", status, cpu.Status.Overflow)
	}
}

func AssertRegisterA(t *testing.T, cpu *Cpu, value uint8) {
	if cpu.RegA != value {
		t.Errorf("Expected registerA to be %#x but was %#x", value, cpu.RegA)
	}
}

func AssertRegisterX(t *testing.T, cpu *Cpu, value uint8) {
	if cpu.RegX != value {
		t.Errorf("Expected registerX to be %#x but was %#x", value, cpu.RegX)
	}
}

func AssertRegisterY(t *testing.T, cpu *Cpu, value uint8) {
	if cpu.RegY != value {
		t.Errorf("Expected registerY to be %#x but was %#x", value, cpu.RegY)
	}
}

func AssertProgramCounter(t *testing.T, cpu *Cpu, value uint16) {
	if cpu.ProgramCounter != value {
		t.Errorf("Expected ProgramCounter to be %#x but was %#x", value, cpu.ProgramCounter)
	}
}

func AssertMemoryValue(t *testing.T, cpu *Cpu, address uint16, value uint8) {
	if cpu.bus.cpuVRam[address] != value {
		t.Errorf("Expected memory value at %#x to be %#x but was %#x", address, value, cpu.bus.cpuVRam[address])
	}
}

func AssertStackPointer(t *testing.T, cpu *Cpu, value uint8) {
	if cpu.StackPointer != value {
		t.Errorf("Expected stack pointer to be %#x but was %#x", value, cpu.StackPointer)
	}
}
