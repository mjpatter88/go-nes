package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// In order to work, this needs to be loaded at 0x600 rather than the "normal" 0x8000.
// See: https://github.com/bugzmanov/nes_ebook/blob/master/code/ch3.4/src/cpu.rs#L244-L258
const MEM_ADDRESS = 0x600

const VIDEO_MEM_ADDRESS = 0x200

// TODO(mjpatter88): make the window 640x640 and scale the 32x32 nes cpu video output to fill it
const windowWidth = 32
const windowHeight = 32
const windowPositionX = 500
const windowPositionY = 500

const pixelWidth = 32
const pixelHeight = 32

// Target fps is 60 -> 1,000,000 / 60 = 16,666.666
const usPerFrame = 16666

// Example tetris game from: https://bugzmanov.github.io/nes_ebook/chapter_3_4.html
var instr = []uint8{
	0x20, 0x06, 0x06, 0x20, 0x38, 0x06, 0x20, 0x0d, 0x06, 0x20, 0x2a, 0x06, 0x60, 0xa9, 0x02, 0x85,
	0x02, 0xa9, 0x04, 0x85, 0x03, 0xa9, 0x11, 0x85, 0x10, 0xa9, 0x10, 0x85, 0x12, 0xa9, 0x0f, 0x85,
	0x14, 0xa9, 0x04, 0x85, 0x11, 0x85, 0x13, 0x85, 0x15, 0x60, 0xa5, 0xfe, 0x85, 0x00, 0xa5, 0xfe,
	0x29, 0x03, 0x18, 0x69, 0x02, 0x85, 0x01, 0x60, 0x20, 0x4d, 0x06, 0x20, 0x8d, 0x06, 0x20, 0xc3,
	0x06, 0x20, 0x19, 0x07, 0x20, 0x20, 0x07, 0x20, 0x2d, 0x07, 0x4c, 0x38, 0x06, 0xa5, 0xff, 0xc9,
	0x77, 0xf0, 0x0d, 0xc9, 0x64, 0xf0, 0x14, 0xc9, 0x73, 0xf0, 0x1b, 0xc9, 0x61, 0xf0, 0x22, 0x60,
	0xa9, 0x04, 0x24, 0x02, 0xd0, 0x26, 0xa9, 0x01, 0x85, 0x02, 0x60, 0xa9, 0x08, 0x24, 0x02, 0xd0,
	0x1b, 0xa9, 0x02, 0x85, 0x02, 0x60, 0xa9, 0x01, 0x24, 0x02, 0xd0, 0x10, 0xa9, 0x04, 0x85, 0x02,
	0x60, 0xa9, 0x02, 0x24, 0x02, 0xd0, 0x05, 0xa9, 0x08, 0x85, 0x02, 0x60, 0x60, 0x20, 0x94, 0x06,
	0x20, 0xa8, 0x06, 0x60, 0xa5, 0x00, 0xc5, 0x10, 0xd0, 0x0d, 0xa5, 0x01, 0xc5, 0x11, 0xd0, 0x07,
	0xe6, 0x03, 0xe6, 0x03, 0x20, 0x2a, 0x06, 0x60, 0xa2, 0x02, 0xb5, 0x10, 0xc5, 0x10, 0xd0, 0x06,
	0xb5, 0x11, 0xc5, 0x11, 0xf0, 0x09, 0xe8, 0xe8, 0xe4, 0x03, 0xf0, 0x06, 0x4c, 0xaa, 0x06, 0x4c,
	0x35, 0x07, 0x60, 0xa6, 0x03, 0xca, 0x8a, 0xb5, 0x10, 0x95, 0x12, 0xca, 0x10, 0xf9, 0xa5, 0x02,
	0x4a, 0xb0, 0x09, 0x4a, 0xb0, 0x19, 0x4a, 0xb0, 0x1f, 0x4a, 0xb0, 0x2f, 0xa5, 0x10, 0x38, 0xe9,
	0x20, 0x85, 0x10, 0x90, 0x01, 0x60, 0xc6, 0x11, 0xa9, 0x01, 0xc5, 0x11, 0xf0, 0x28, 0x60, 0xe6,
	0x10, 0xa9, 0x1f, 0x24, 0x10, 0xf0, 0x1f, 0x60, 0xa5, 0x10, 0x18, 0x69, 0x20, 0x85, 0x10, 0xb0,
	0x01, 0x60, 0xe6, 0x11, 0xa9, 0x06, 0xc5, 0x11, 0xf0, 0x0c, 0x60, 0xc6, 0x10, 0xa5, 0x10, 0x29,
	0x1f, 0xc9, 0x1f, 0xf0, 0x01, 0x60, 0x4c, 0x35, 0x07, 0xa0, 0x00, 0xa5, 0xfe, 0x91, 0x00, 0x60,
	0xa6, 0x03, 0xa9, 0x00, 0x81, 0x10, 0xa2, 0x00, 0xa9, 0x01, 0x81, 0x10, 0x60, 0xa2, 0x00, 0xea,
	0xea, 0xca, 0xd0, 0xfb, 0x60,
}

var colorPalette = map[uint8]color.RGBA{
	0:  {0x00, 0x00, 0x00, 0xFF}, // Black
	1:  {0xFF, 0xFF, 0xFF, 0xFF}, // White
	2:  {0x51, 0x2D, 0x38, 0xFF}, // Mauve
	3:  {0x29, 0x78, 0xA0, 0xFF}, // Blue
	4:  {0xD3, 0x4F, 0x73, 0xFF}, // Redish
	5:  {0xD7, 0xAF, 0x70, 0xFF}, // Yellowish
	6:  {0xFF, 0x78, 0x5A, 0xFF}, // Orange
	7:  {0xCC, 0x97, 0x8E, 0xFF}, // Brown
	8:  {0x7C, 0xEA, 0x9C, 0xFF}, // Light Green
	9:  {0xE2, 0x4E, 0x1B, 0xFF}, // Red
	10: {0x00, 0x4E, 0x98, 0xFF}, // Royal Blue
	11: {0xA5, 0xFF, 0xD6, 0xFF}, // Marine
	12: {0x94, 0xFB, 0xAB, 0xFF}, // Mint
	13: {0x89, 0x80, 0xF5, 0xFF}, // Purple
	14: {0xB5, 0x65, 0x76, 0xFF}, // Rose
	15: {0xF8, 0x8D, 0xAD, 0xFF}, // Pink
}

func initSDL() (*sdl.Window, *sdl.Renderer) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow(
		"6502 Tetris",
		windowPositionX,
		windowPositionY,
		windowWidth,
		windowHeight,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		panic(err)
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}
	return window, renderer

}

func main() {
	window, renderer := initSDL()

	tex, err := renderer.CreateTexture(
		uint32(sdl.PIXELFORMAT_RGBA32),
		sdl.TEXTUREACCESS_STREAMING,
		windowWidth,
		windowHeight,
	)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()
	defer window.Destroy()
	defer renderer.Destroy()
	var screenBytes = [windowWidth * windowHeight * 4]byte{}

	cpu := Cpu{}
	cpu.LoadAtAddress(instr, MEM_ADDRESS)

	startTime := time.Now()
	lastDrawTime := time.Now()
	running := true
	steps := 0
	frameCount := 0
	framesProcessed := 0

	// Print fps each second
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				frames := frameCount - framesProcessed
				framesProcessed = frameCount
				fmt.Println("fps: ", frames)
			}
		}
	}()

	for running && !cpu.Status.Break {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}

		cpu.Step()
		steps += 1

		// Cap at 60 fps.
		elapsedTime := time.Since(lastDrawTime).Microseconds()
		if elapsedTime > usPerFrame {
			lastDrawTime = time.Now()
			fillScreen(&screenBytes, &cpu)
			drawFrame(renderer, tex, &screenBytes)
			frameCount += 1
		}

		// TODO(mjpatter88): think about how to actually manage frequency rather than harcoding a delay here.
		time.Sleep(100 * time.Microsecond)
	}

	// For short executions, these numbers might be surprising.
	// I've found that the first event handling loop and last event handling loop
	// take a long time (up to 300ms). This skews the overall time especially
	// badly on short runs. As the overall run increases, this impact is lessened and
	// the fps approaches 60 as it should. I tried various hacks to work around it, but
	// in the end I decided it wasn't worth it.
	elapsedMS := time.Since(startTime).Milliseconds()
	fmt.Printf("Cpu Steps: %d\n", steps)
	fmt.Printf("Elapsed MS: %d\n", elapsedMS)
	fmt.Printf("Frames Drawn: %d\n", frameCount)
	fmt.Printf("FPS: %f\n", (float64(frameCount) / (float64(elapsedMS) / 1000.0)))
}

func drawFrame(renderer *sdl.Renderer, texture *sdl.Texture, screen *[windowWidth * windowHeight * 4]byte) {
	bytes, _, err := texture.Lock(nil)
	if err != nil {
		panic(err)
	}
	for i := 0; i < int(windowWidth*windowHeight*4); i++ {
		bytes[i] = screen[i]
	}
	texture.Unlock()
	rect := sdl.Rect{X: 0, Y: 0, W: int32(windowWidth), H: int32(windowHeight)}
	err = renderer.Copy(texture, nil, &rect)
	if err != nil {
		panic(err)
	}

	renderer.Present()
}

// Use the nes cpu video memory to render the display based on an arbitrary color palette
func fillScreen(screenBytes *[windowWidth * windowHeight * 4]byte, cpu *Cpu) {
	// TODO(mjpatter88): Eventually each pixel of cpu video memory should correspond to a 20x20 square of pixels
	for i := 0; i < int(pixelWidth*pixelHeight); i += 1 {
		color := colorPalette[cpu.readMemory(uint16(VIDEO_MEM_ADDRESS+i))]
		screenBytes[i*4] = color.R
		screenBytes[(i*4)+1] = color.G
		screenBytes[(i*4)+2] = color.B
		screenBytes[(i*4)+3] = color.A
	}
}
