package main

import "testing"

func TestDecode_BRK(t *testing.T) {
	instr := Decode(0x00)
	AssertAction(t, instr, "BRK")
	AssertAddressingMode(t, instr, IMPLICIT)
	AssertNumberOfBytes(t, instr, 1)
}

func TestDecodeLDA(t *testing.T) {
	t.Run("LDA", func(t *testing.T) {
		instr := Decode(0xa9)
		AssertAction(t, instr, "LDA")
		AssertAddressingMode(t, instr, IMMEDIATE)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("Zero", func(t *testing.T) {
		instr := Decode(0xa5)
		AssertAction(t, instr, "LDA")
		AssertAddressingMode(t, instr, ZERO)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("Zero X", func(t *testing.T) {
		instr := Decode(0xb5)
		AssertAction(t, instr, "LDA")
		AssertAddressingMode(t, instr, ZERO_X)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("ABSOLUTE", func(t *testing.T) {
		instr := Decode(0xad)
		AssertAction(t, instr, "LDA")
		AssertAddressingMode(t, instr, ABSOLUTE)
		AssertNumberOfBytes(t, instr, 3)
	})

	t.Run("ABSOLUTE X", func(t *testing.T) {
		instr := Decode(0xbd)
		AssertAction(t, instr, "LDA")
		AssertAddressingMode(t, instr, ABSOLUTE_X)
		AssertNumberOfBytes(t, instr, 3)
	})

	t.Run("ABSOLUTE Y", func(t *testing.T) {
		instr := Decode(0xb9)
		AssertAction(t, instr, "LDA")
		AssertAddressingMode(t, instr, ABSOLUTE_Y)
		AssertNumberOfBytes(t, instr, 3)
	})

	t.Run("INDIRECT X", func(t *testing.T) {
		instr := Decode(0xa1)
		AssertAction(t, instr, "LDA")
		AssertAddressingMode(t, instr, INDIRECT_X)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("INDIRECT Y", func(t *testing.T) {
		instr := Decode(0xb1)
		AssertAction(t, instr, "LDA")
		AssertAddressingMode(t, instr, INDIRECT_Y)
		AssertNumberOfBytes(t, instr, 2)
	})
}

func TestDecodeLDX(t *testing.T) {
	t.Run("LDX", func(t *testing.T) {
		instr := Decode(0xa2)
		AssertAction(t, instr, "LDX")
		AssertAddressingMode(t, instr, IMMEDIATE)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("Zero", func(t *testing.T) {
		instr := Decode(0xa6)
		AssertAction(t, instr, "LDX")
		AssertAddressingMode(t, instr, ZERO)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("Zero Y", func(t *testing.T) {
		instr := Decode(0xb6)
		AssertAction(t, instr, "LDX")
		AssertAddressingMode(t, instr, ZERO_Y)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("ABSOLUTE", func(t *testing.T) {
		instr := Decode(0xae)
		AssertAction(t, instr, "LDX")
		AssertAddressingMode(t, instr, ABSOLUTE)
		AssertNumberOfBytes(t, instr, 3)
	})

	t.Run("ABSOLUTE Y", func(t *testing.T) {
		instr := Decode(0xbe)
		AssertAction(t, instr, "LDX")
		AssertAddressingMode(t, instr, ABSOLUTE_Y)
		AssertNumberOfBytes(t, instr, 3)
	})
}

func TestDecodeLDY(t *testing.T) {
	t.Run("LDY", func(t *testing.T) {
		instr := Decode(0xa0)
		AssertAction(t, instr, "LDY")
		AssertAddressingMode(t, instr, IMMEDIATE)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("Zero", func(t *testing.T) {
		instr := Decode(0xa4)
		AssertAction(t, instr, "LDY")
		AssertAddressingMode(t, instr, ZERO)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("Zero X", func(t *testing.T) {
		instr := Decode(0xb4)
		AssertAction(t, instr, "LDY")
		AssertAddressingMode(t, instr, ZERO_X)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("ABSOLUTE", func(t *testing.T) {
		instr := Decode(0xac)
		AssertAction(t, instr, "LDY")
		AssertAddressingMode(t, instr, ABSOLUTE)
		AssertNumberOfBytes(t, instr, 3)
	})

	t.Run("ABSOLUTE X", func(t *testing.T) {
		instr := Decode(0xbc)
		AssertAction(t, instr, "LDY")
		AssertAddressingMode(t, instr, ABSOLUTE_X)
		AssertNumberOfBytes(t, instr, 3)
	})
}

func TestDecodeLSR(t *testing.T) {
	t.Run("LSR", func(t *testing.T) {
		instr := Decode(0x4a)
		AssertAction(t, instr, "LSR")
		AssertAddressingMode(t, instr, ACCUMULATOR)
		AssertNumberOfBytes(t, instr, 1)
	})

	t.Run("Zero", func(t *testing.T) {
		instr := Decode(0x46)
		AssertAction(t, instr, "LSR")
		AssertAddressingMode(t, instr, ZERO)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("Zero X", func(t *testing.T) {
		instr := Decode(0x56)
		AssertAction(t, instr, "LSR")
		AssertAddressingMode(t, instr, ZERO_X)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("ABSOLUTE", func(t *testing.T) {
		instr := Decode(0x4e)
		AssertAction(t, instr, "LSR")
		AssertAddressingMode(t, instr, ABSOLUTE)
		AssertNumberOfBytes(t, instr, 3)
	})

	t.Run("ABSOLUTE X", func(t *testing.T) {
		instr := Decode(0x5e)
		AssertAction(t, instr, "LSR")
		AssertAddressingMode(t, instr, ABSOLUTE_X)
		AssertNumberOfBytes(t, instr, 3)
	})
}

func TestDecodeAND(t *testing.T) {
	t.Run("AND", func(t *testing.T) {
		instr := Decode(0x29)
		AssertAction(t, instr, "AND")
		AssertAddressingMode(t, instr, IMMEDIATE)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("Zero", func(t *testing.T) {
		instr := Decode(0x25)
		AssertAction(t, instr, "AND")
		AssertAddressingMode(t, instr, ZERO)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("Zero X", func(t *testing.T) {
		instr := Decode(0x35)
		AssertAction(t, instr, "AND")
		AssertAddressingMode(t, instr, ZERO_X)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("ABSOLUTE", func(t *testing.T) {
		instr := Decode(0x2d)
		AssertAction(t, instr, "AND")
		AssertAddressingMode(t, instr, ABSOLUTE)
		AssertNumberOfBytes(t, instr, 3)
	})

	t.Run("ABSOLUTE X", func(t *testing.T) {
		instr := Decode(0x3d)
		AssertAction(t, instr, "AND")
		AssertAddressingMode(t, instr, ABSOLUTE_X)
		AssertNumberOfBytes(t, instr, 3)
	})

	t.Run("ABSOLUTE Y", func(t *testing.T) {
		instr := Decode(0x39)
		AssertAction(t, instr, "AND")
		AssertAddressingMode(t, instr, ABSOLUTE_Y)
		AssertNumberOfBytes(t, instr, 3)
	})

	t.Run("INDIRECT X", func(t *testing.T) {
		instr := Decode(0x21)
		AssertAction(t, instr, "AND")
		AssertAddressingMode(t, instr, INDIRECT_X)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("INDIRECT Y", func(t *testing.T) {
		instr := Decode(0x31)
		AssertAction(t, instr, "AND")
		AssertAddressingMode(t, instr, INDIRECT_Y)
		AssertNumberOfBytes(t, instr, 2)
	})
}

func TestDecodeSTA(t *testing.T) {
	t.Run("Zero", func(t *testing.T) {
		instr := Decode(0x85)
		AssertAction(t, instr, "STA")
		AssertAddressingMode(t, instr, ZERO)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("Zero X", func(t *testing.T) {
		instr := Decode(0x95)
		AssertAction(t, instr, "STA")
		AssertAddressingMode(t, instr, ZERO_X)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("ABSOLUTE", func(t *testing.T) {
		instr := Decode(0x8d)
		AssertAction(t, instr, "STA")
		AssertAddressingMode(t, instr, ABSOLUTE)
		AssertNumberOfBytes(t, instr, 3)
	})

	t.Run("ABSOLUTE X", func(t *testing.T) {
		instr := Decode(0x9d)
		AssertAction(t, instr, "STA")
		AssertAddressingMode(t, instr, ABSOLUTE_X)
		AssertNumberOfBytes(t, instr, 3)
	})

	t.Run("ABSOLUTE Y", func(t *testing.T) {
		instr := Decode(0x99)
		AssertAction(t, instr, "STA")
		AssertAddressingMode(t, instr, ABSOLUTE_Y)
		AssertNumberOfBytes(t, instr, 3)
	})

	t.Run("INDIRECT X", func(t *testing.T) {
		instr := Decode(0x81)
		AssertAction(t, instr, "STA")
		AssertAddressingMode(t, instr, INDIRECT_X)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("INDIRECT Y", func(t *testing.T) {
		instr := Decode(0x91)
		AssertAction(t, instr, "STA")
		AssertAddressingMode(t, instr, INDIRECT_Y)
		AssertNumberOfBytes(t, instr, 2)
	})
}

func TestDecodeADC(t *testing.T) {
	t.Run("Immediate", func(t *testing.T) {
		instr := Decode(0x69)
		AssertAction(t, instr, "ADC")
		AssertAddressingMode(t, instr, IMMEDIATE)
		AssertNumberOfBytes(t, instr, 2)
	})
	t.Run("Zero", func(t *testing.T) {
		instr := Decode(0x65)
		AssertAction(t, instr, "ADC")
		AssertAddressingMode(t, instr, ZERO)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("Zero X", func(t *testing.T) {
		instr := Decode(0x75)
		AssertAction(t, instr, "ADC")
		AssertAddressingMode(t, instr, ZERO_X)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("ABSOLUTE", func(t *testing.T) {
		instr := Decode(0x6d)
		AssertAction(t, instr, "ADC")
		AssertAddressingMode(t, instr, ABSOLUTE)
		AssertNumberOfBytes(t, instr, 3)
	})

	t.Run("ABSOLUTE X", func(t *testing.T) {
		instr := Decode(0x7d)
		AssertAction(t, instr, "ADC")
		AssertAddressingMode(t, instr, ABSOLUTE_X)
		AssertNumberOfBytes(t, instr, 3)
	})

	t.Run("ABSOLUTE Y", func(t *testing.T) {
		instr := Decode(0x79)
		AssertAction(t, instr, "ADC")
		AssertAddressingMode(t, instr, ABSOLUTE_Y)
		AssertNumberOfBytes(t, instr, 3)
	})

	t.Run("INDIRECT X", func(t *testing.T) {
		instr := Decode(0x61)
		AssertAction(t, instr, "ADC")
		AssertAddressingMode(t, instr, INDIRECT_X)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("INDIRECT Y", func(t *testing.T) {
		instr := Decode(0x71)
		AssertAction(t, instr, "ADC")
		AssertAddressingMode(t, instr, INDIRECT_Y)
		AssertNumberOfBytes(t, instr, 2)
	})
}

func TestDecodeCMP(t *testing.T) {
	t.Run("Immediate", func(t *testing.T) {
		instr := Decode(0xc9)
		AssertAction(t, instr, "CMP")
		AssertAddressingMode(t, instr, IMMEDIATE)
		AssertNumberOfBytes(t, instr, 2)
	})
	t.Run("Zero", func(t *testing.T) {
		instr := Decode(0xc5)
		AssertAction(t, instr, "CMP")
		AssertAddressingMode(t, instr, ZERO)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("Zero X", func(t *testing.T) {
		instr := Decode(0xd5)
		AssertAction(t, instr, "CMP")
		AssertAddressingMode(t, instr, ZERO_X)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("ABSOLUTE", func(t *testing.T) {
		instr := Decode(0xcd)
		AssertAction(t, instr, "CMP")
		AssertAddressingMode(t, instr, ABSOLUTE)
		AssertNumberOfBytes(t, instr, 3)
	})

	t.Run("ABSOLUTE X", func(t *testing.T) {
		instr := Decode(0xdd)
		AssertAction(t, instr, "CMP")
		AssertAddressingMode(t, instr, ABSOLUTE_X)
		AssertNumberOfBytes(t, instr, 3)
	})

	t.Run("ABSOLUTE Y", func(t *testing.T) {
		instr := Decode(0xd9)
		AssertAction(t, instr, "CMP")
		AssertAddressingMode(t, instr, ABSOLUTE_Y)
		AssertNumberOfBytes(t, instr, 3)
	})

	t.Run("INDIRECT X", func(t *testing.T) {
		instr := Decode(0xc1)
		AssertAction(t, instr, "CMP")
		AssertAddressingMode(t, instr, INDIRECT_X)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("INDIRECT Y", func(t *testing.T) {
		instr := Decode(0xd1)
		AssertAction(t, instr, "CMP")
		AssertAddressingMode(t, instr, INDIRECT_Y)
		AssertNumberOfBytes(t, instr, 2)
	})
}

func TestDecodeCPX(t *testing.T) {
	t.Run("Immediate", func(t *testing.T) {
		instr := Decode(0xe0)
		AssertAction(t, instr, "CPX")
		AssertAddressingMode(t, instr, IMMEDIATE)
		AssertNumberOfBytes(t, instr, 2)
	})
	t.Run("Zero", func(t *testing.T) {
		instr := Decode(0xe4)
		AssertAction(t, instr, "CPX")
		AssertAddressingMode(t, instr, ZERO)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("ABSOLUTE", func(t *testing.T) {
		instr := Decode(0xec)
		AssertAction(t, instr, "CPX")
		AssertAddressingMode(t, instr, ABSOLUTE)
		AssertNumberOfBytes(t, instr, 3)
	})
}

func TestDecodeCPY(t *testing.T) {
	t.Run("Immediate", func(t *testing.T) {
		instr := Decode(0xc0)
		AssertAction(t, instr, "CPY")
		AssertAddressingMode(t, instr, IMMEDIATE)
		AssertNumberOfBytes(t, instr, 2)
	})
	t.Run("Zero", func(t *testing.T) {
		instr := Decode(0xc4)
		AssertAction(t, instr, "CPY")
		AssertAddressingMode(t, instr, ZERO)
		AssertNumberOfBytes(t, instr, 2)
	})

	t.Run("ABSOLUTE", func(t *testing.T) {
		instr := Decode(0xcc)
		AssertAction(t, instr, "CPY")
		AssertAddressingMode(t, instr, ABSOLUTE)
		AssertNumberOfBytes(t, instr, 3)
	})
}

func TestDecode_TAX(t *testing.T) {
	instr := Decode(0xaa)
	AssertAction(t, instr, "TAX")
	AssertAddressingMode(t, instr, IMPLICIT)
	AssertNumberOfBytes(t, instr, 1)
}

func TestDecode_TXA(t *testing.T) {
	instr := Decode(0x8a)
	AssertAction(t, instr, "TXA")
	AssertAddressingMode(t, instr, IMPLICIT)
	AssertNumberOfBytes(t, instr, 1)
}

func TestDecode_DEX(t *testing.T) {
	instr := Decode(0xca)
	AssertAction(t, instr, "DEX")
	AssertAddressingMode(t, instr, IMPLICIT)
	AssertNumberOfBytes(t, instr, 1)
}

func TestDecode_INX(t *testing.T) {
	instr := Decode(0xe8)
	AssertAction(t, instr, "INX")
	AssertAddressingMode(t, instr, IMPLICIT)
	AssertNumberOfBytes(t, instr, 1)
}

func TestDecode_TAY(t *testing.T) {
	instr := Decode(0xa8)
	AssertAction(t, instr, "TAY")
	AssertAddressingMode(t, instr, IMPLICIT)
	AssertNumberOfBytes(t, instr, 1)
}

func TestDecode_TYA(t *testing.T) {
	instr := Decode(0x98)
	AssertAction(t, instr, "TYA")
	AssertAddressingMode(t, instr, IMPLICIT)
	AssertNumberOfBytes(t, instr, 1)
}

func TestDecode_DEY(t *testing.T) {
	instr := Decode(0x88)
	AssertAction(t, instr, "DEY")
	AssertAddressingMode(t, instr, IMPLICIT)
	AssertNumberOfBytes(t, instr, 1)
}

func TestDecode_INY(t *testing.T) {
	instr := Decode(0xc8)
	AssertAction(t, instr, "INY")
	AssertAddressingMode(t, instr, IMPLICIT)
	AssertNumberOfBytes(t, instr, 1)
}

func TestDecode_CLC(t *testing.T) {
	instr := Decode(0x18)
	AssertAction(t, instr, "CLC")
	AssertAddressingMode(t, instr, IMPLICIT)
	AssertNumberOfBytes(t, instr, 1)
}

func TestDecode_JSR(t *testing.T) {
	instr := Decode(0x20)
	AssertAction(t, instr, "JSR")
	AssertAddressingMode(t, instr, ABSOLUTE)
	AssertNumberOfBytes(t, instr, 3)
}

func TestDecode_RTS(t *testing.T) {
	instr := Decode(0x60)
	AssertAction(t, instr, "RTS")
	AssertAddressingMode(t, instr, IMPLICIT)
	AssertNumberOfBytes(t, instr, 1)
}

func TestDecode_BPL(t *testing.T) {
	instr := Decode(0x10)
	AssertAction(t, instr, "BPL")
	AssertAddressingMode(t, instr, RELATIVE)
	AssertNumberOfBytes(t, instr, 2)
}

func TestDecode_BMI(t *testing.T) {
	instr := Decode(0x30)
	AssertAction(t, instr, "BMI")
	AssertAddressingMode(t, instr, RELATIVE)
	AssertNumberOfBytes(t, instr, 2)
}

func TestDecode_BVC(t *testing.T) {
	instr := Decode(0x50)
	AssertAction(t, instr, "BVC")
	AssertAddressingMode(t, instr, RELATIVE)
	AssertNumberOfBytes(t, instr, 2)
}

func TestDecode_BVS(t *testing.T) {
	instr := Decode(0x70)
	AssertAction(t, instr, "BVS")
	AssertAddressingMode(t, instr, RELATIVE)
	AssertNumberOfBytes(t, instr, 2)
}

func TestDecode_BCC(t *testing.T) {
	instr := Decode(0x90)
	AssertAction(t, instr, "BCC")
	AssertAddressingMode(t, instr, RELATIVE)
	AssertNumberOfBytes(t, instr, 2)
}

func TestDecode_BCS(t *testing.T) {
	instr := Decode(0xB0)
	AssertAction(t, instr, "BCS")
	AssertAddressingMode(t, instr, RELATIVE)
	AssertNumberOfBytes(t, instr, 2)
}

func TestDecode_BEQ(t *testing.T) {
	instr := Decode(0xf0)
	AssertAction(t, instr, "BEQ")
	AssertAddressingMode(t, instr, RELATIVE)
	AssertNumberOfBytes(t, instr, 2)
}

func TestDecode_BNE(t *testing.T) {
	instr := Decode(0xd0)
	AssertAction(t, instr, "BNE")
	AssertAddressingMode(t, instr, RELATIVE)
	AssertNumberOfBytes(t, instr, 2)
}

func AssertAction(t *testing.T, instr Instruction, action string) {
	if instr.Action != action {
		t.Errorf("Expected Action to be %s but was %s", action, instr.Action)
	}
}

func AssertAddressingMode(t *testing.T, instr Instruction, addressingMode int) {
	if instr.AddressingMode != addressingMode {
		t.Errorf("Expected AddressingMode to be %d but was %d", addressingMode, instr.AddressingMode)
	}
}

func AssertNumberOfBytes(t *testing.T, instr Instruction, number int) {
	if instr.NumberOfBytes != number {
		t.Errorf("Expected NumberOfBytes to be %d but was %d", number, instr.NumberOfBytes)
	}
}
