package main

import (
	"testing"
)

func TestReadMemory(t *testing.T) {
	bus := Bus{}
	memBytes := []uint8{0x01, 0x02, 0x03}
	for i := 0; i < 3; i++ {
		bus.cpuVRam[i] = memBytes[i]
	}

	for i := 0; i < 3; i++ {
		byte := bus.ReadMemory(uint16(i))
		if byte != uint8(i+1) {
			t.Errorf("wanted %#x but got %#x", i+1, byte)
		}
	}
}

func TestWriteMemory(t *testing.T) {
	bus := Bus{}
	memBytes := []uint8{0x01, 0x02, 0x03}
	for i := 0; i < 3; i++ {
		bus.WriteMemory(uint16(i), memBytes[i])
	}

	for i := 0; i < 3; i++ {
		byte := bus.cpuVRam[i]
		if byte != uint8(i+1) {
			t.Errorf("wanted %#x but got %#x", i+1, byte)
		}
	}
}

func TestReadMemory_u16(t *testing.T) {
	bus := Bus{}
	bus.cpuVRam[0x100] = 0x11
	bus.cpuVRam[0x101] = 0x22

	value := bus.ReadMemory_u16(0x100)

	if value != 0x2211 {
		t.Errorf("wanted %#x but got %#x", 0x2211, value)
	}
}

func TestWriteMemory_u16(t *testing.T) {
	bus := Bus{}
	bus.WriteMemory_u16(0x100, 0x1122)

	firstByte := bus.cpuVRam[0x100]
	secondByte := bus.cpuVRam[0x100+1]

	if firstByte != 0x22 {
		t.Errorf("wanted %#x but got %#x", 0x22, firstByte)
	}

	if secondByte != 0x11 {
		t.Errorf("wanted %#x but got %#x", 0x11, secondByte)
	}
}
