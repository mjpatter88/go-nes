package main

type Bus struct {
	cpuVRam [2048]uint8
	// TODO(mjpatter88): fix this hack once roms are supported.
	startingPC uint16
}

func (b *Bus) ReadMemory(address uint16) uint8 {
	return b.cpuVRam[address]
}

func (b *Bus) WriteMemory(address uint16, value uint8) {
	b.cpuVRam[address] = value
}

// nes is little-endian so 16-bit values read from memory need to handle this byte order.
// NOTE: this just impacts the 16-bit values from memory, not the 16-bit memory index.
func (b *Bus) ReadMemory_u16(address uint16) uint16 {
	// TODO(mjpatter88): fix this hack once roms are supported.
	if address == PROG_REFERENCE_MEM_ADDRESS {
		return b.startingPC
	}
	firstByte := uint16(b.ReadMemory(address))
	secondByte := uint16(b.ReadMemory(address + 1))
	return (secondByte << 8) | (firstByte)
}

// nes is little-endian so 16-bit values written to memory need to handle this byte order.
// NOTE: this just impacts the 16-bit values written to memory, not the 16-bit memory index.
func (b *Bus) WriteMemory_u16(address uint16, value uint16) {
	firstByte := (value) & 0xFF
	secondByte := (value >> 8) & 0xFF

	// TODO(mjpatter88): fix this hack once roms are supported.
	if address == PROG_REFERENCE_MEM_ADDRESS {
		b.startingPC = value
		return
	}
	b.WriteMemory(address, uint8(firstByte))
	b.WriteMemory(address+1, uint8(secondByte))
}
