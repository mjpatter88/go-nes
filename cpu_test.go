package main

import "testing"

func TestCpu(t *testing.T) {
	AssertBreak := func(cpu *Cpu, status bool) {
		if cpu.Status.Break != status {
			t.Errorf("Expected Break status to be %t but was %t", status, cpu.Status.Break)
		}
	}
	AssertZero := func(cpu *Cpu, status bool) {
		if cpu.Status.Zero != status {
			t.Errorf("Expected Zero status to be %t but was %t", status, cpu.Status.Zero)
		}
	}
	AssertNegative := func(cpu *Cpu, status bool) {
		if cpu.Status.Negative != status {
			t.Errorf("Expected Negative status to be %t but was %t", status, cpu.Status.Negative)
		}
	}
	AssertRegisterA := func(cpu *Cpu, value uint8) {
		if cpu.RegA != value {
			t.Errorf("Expected registerA to be %#x but was %#x", value, cpu.RegA)
		}
	}
	t.Run("Test BRK", func(t *testing.T) {
		cpu := Cpu{}
		cpu.execute([]uint8{BRK})

		AssertBreak(&cpu, true)
	})

	t.Run("Test LDA", func(t *testing.T) {
		cpu := Cpu{}
		cpu.execute([]uint8{LDA, 0x4a, BRK})

		AssertRegisterA(&cpu, 0x4a)
		AssertZero(&cpu, false)
		AssertNegative(&cpu, false)
	})

	t.Run("Test LDA - Zero flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.execute([]uint8{LDA, 0x00, BRK})

		AssertRegisterA(&cpu, 0x00)
		AssertZero(&cpu, true)
		AssertNegative(&cpu, false)
	})

	t.Run("Test LDA - Negative flag", func(t *testing.T) {
		cpu := Cpu{}
		cpu.execute([]uint8{LDA, 0xF0, BRK})
		AssertRegisterA(&cpu, 0xF0)
		AssertZero(&cpu, false)
		AssertNegative(&cpu, true)
	})
}
