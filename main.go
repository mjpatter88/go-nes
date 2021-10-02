package main

// a9 c0 aa e8 00
func main() {
	cpu := Cpu{}
	cpu.execute([]uint8{0xa9, 0xc0, 0xaa, 0xe8, 0x00})
}
