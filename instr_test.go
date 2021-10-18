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

func TestDecode_TAY(t *testing.T) {
	instr := Decode(0xa8)
	AssertAction(t, instr, "TAY")
	AssertAddressingMode(t, instr, IMPLICIT)
	AssertNumberOfBytes(t, instr, 1)
}

func TestDecode_INX(t *testing.T) {
	instr := Decode(0xe8)
	AssertAction(t, instr, "INX")
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
