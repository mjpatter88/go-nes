package main

import "testing"

func TestDecode_BRK(t *testing.T) {
	instr := Decode(0x00)
	AssertAction(t, instr, "BRK")
	AssertAddressingMode(t, instr, IMPLICIT)
	AssertNumberOfBytes(t, instr, 1)
}

func TestDecode_LDA(t *testing.T) {
	instr := Decode(0xa9)
	AssertAction(t, instr, "LDA")
	AssertAddressingMode(t, instr, IMMEDIATE)
	AssertNumberOfBytes(t, instr, 2)
}

func TestDecode_LDA_ZERO(t *testing.T) {
	instr := Decode(0xa5)
	AssertAction(t, instr, "LDA")
	AssertAddressingMode(t, instr, ZERO)
	AssertNumberOfBytes(t, instr, 2)
}

func TestDecode_LDA_ZERO_X(t *testing.T) {
	instr := Decode(0xb5)
	AssertAction(t, instr, "LDA")
	AssertAddressingMode(t, instr, ZERO_X)
	AssertNumberOfBytes(t, instr, 2)
}

func TestDecode_LDA_ABS(t *testing.T) {
	instr := Decode(0xad)
	AssertAction(t, instr, "LDA")
	AssertAddressingMode(t, instr, ABSOLUTE)
	AssertNumberOfBytes(t, instr, 3)
}

func TestDecode_LDA_ABS_X(t *testing.T) {
	instr := Decode(0xbd)
	AssertAction(t, instr, "LDA")
	AssertAddressingMode(t, instr, ABSOLUTE_X)
	AssertNumberOfBytes(t, instr, 3)
}

func TestDecode_LDA_ABS_Y(t *testing.T) {
	instr := Decode(0xb9)
	AssertAction(t, instr, "LDA")
	AssertAddressingMode(t, instr, ABSOLUTE_Y)
	AssertNumberOfBytes(t, instr, 3)
}

func TestDecode_LDA_IND_X(t *testing.T) {
	instr := Decode(0xa1)
	AssertAction(t, instr, "LDA")
	AssertAddressingMode(t, instr, INDIRECT_X)
	AssertNumberOfBytes(t, instr, 2)
}

func TestDecode_LDA_IND_Y(t *testing.T) {
	instr := Decode(0xb1)
	AssertAction(t, instr, "LDA")
	AssertAddressingMode(t, instr, INDIRECT_Y)
	AssertNumberOfBytes(t, instr, 2)
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
