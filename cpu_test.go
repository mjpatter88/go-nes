package main

import "testing"

func TestBRK(t *testing.T) {
	cpu := Cpu{}
	cpu.Execute([]uint8{BRK})

	AssertBreak(t, &cpu, true)
}

func TestLDA(t *testing.T) {
	t.Run("LDA", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDA, 0x4a, BRK})

		AssertRegisterA(t, &cpu, 0x4a)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDA, 0x00, BRK})

		AssertRegisterA(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.Execute([]uint8{LDA, 0xf0, BRK})
		AssertRegisterA(t, &cpu, 0xf0)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})
}

func TestTAX(t *testing.T) {
	t.Run("TAX", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x3a
		cpu.Execute([]uint8{TAX, BRK})

		AssertRegisterX(t, &cpu, 0x3a)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0x00
		cpu.Execute([]uint8{TAX, BRK})

		AssertRegisterX(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegA = 0xf0
		cpu.Execute([]uint8{TAX, BRK})

		AssertRegisterX(t, &cpu, 0xf0)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})
}

func TestINX(t *testing.T) {
	t.Run("INX", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegX = 0x3a
		cpu.Execute([]uint8{INX, BRK})

		AssertRegisterX(t, &cpu, 0x3b)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegX = 0xff
		cpu.Execute([]uint8{INX, BRK})

		AssertRegisterX(t, &cpu, 0x00)
		AssertZero(t, &cpu, true)
		AssertNegative(t, &cpu, false)
	})

	t.Run("Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.RegX = 0x7f
		cpu.Execute([]uint8{INX, BRK})

		AssertRegisterX(t, &cpu, 0x80)
		AssertZero(t, &cpu, false)
		AssertNegative(t, &cpu, true)
	})
}

func TestFiveInstructions(t *testing.T) {
	cpu := Cpu{}
	cpu.Execute([]uint8{LDA, 0xc0, TAX, INX, BRK})

	AssertRegisterX(t, &cpu, 0xc1)
	AssertProgramCounter(t, &cpu, 0x8005)
}

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
func AssertNegative(t *testing.T, cpu *Cpu, status bool) {
	if cpu.Status.Negative != status {
		t.Errorf("Expected Negative status to be %t but was %t", status, cpu.Status.Negative)
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

func AssertProgramCounter(t *testing.T, cpu *Cpu, value uint16) {
	if cpu.ProgramCounter != value {
		t.Errorf("Expected ProgramCounter to be %#x but was %#x", value, cpu.ProgramCounter)
	}
}

func TestLoad(t *testing.T) {
	cpu := Cpu{}
	programBytes := []uint8{0x01, 0x02, 0x03}
	cpu.load(programBytes)

	for i := 0; i < 3; i++ {
		memAddress := 0x8000 + i
		if cpu.memory[memAddress] != uint8(i+1) {
			t.Errorf("Expected memory[%#x] to be %#x but was %#x", memAddress, i, cpu.memory[memAddress])
		}
	}
}

func TestReadMemory(t *testing.T) {
	cpu := Cpu{}
	memBytes := []uint8{0x01, 0x02, 0x03}
	for i := 0; i < 3; i++ {
		cpu.memory[i] = memBytes[i]
	}

	for i := 0; i < 3; i++ {
		byte := cpu.readMemory(uint16(i))
		if byte != uint8(i+1) {
			t.Errorf("wanted %#x but got %#x", i+1, byte)
		}
	}
}
