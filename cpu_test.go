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
		t.Errorf("Expected registerX to be %#x but was %#x", value, cpu.RegA)
	}
}
