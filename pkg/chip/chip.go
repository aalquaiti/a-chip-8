// Package chip provide functionality for the chip-8 interpreter
package chip

const (
	screenWidth  = 64
	screenHeight = 32
	windowWidth  = 640
	windowHeight = 320
)

var screen *Screen
var sound *Sound

// Chip represents chip-9 interpreter
type Chip struct{}

func init() {
	screen = NewScreen(screenWidth, screenHeight, windowWidth,
		windowHeight, Update)
	sound = NewSound()
	// Load("test_opcode.ch8")
	// Load("BC_test.ch8")
	Load("roms/INVADERS")

	// for y := 0; y < screenHeight; y++ {
	// 	for x := 0; x < screenWidth; x++ {
	// 		if y == (screenHeight/2)-1 || y == screenHeight/2 {
	// 			screen.Draw(x, y, true)
	// 		}
	// 	}
	// }

	// // V0 = 0
	// ram[0x200] = 0x60
	// ram[0x201] = 0x00

	// // V1 = 0
	// ram[0x202] = 0x61
	// ram[0x203] = 0x00

	// // V2 = A
	// ram[0x204] = 0x62
	// ram[0x205] = 0x0A

	// // Set I to point to char A
	// ram[0x206] = 0xF2
	// ram[0x207] = 0x29 // constant

	// // Draw character A at (V0,V1)
	// ram[0x208] = 0xD0
	// ram[0x209] = 0x15

	// // Infinite Loop
	// ram[0x20A] = 0x12
	// ram[0x20B] = 0x0A
}

// Start the interpreter
func Start() {
	screen.Show()
}

// Update the current the state of the machine
func Update() {
	// chip-8 runs at about 500hz. As a tick is called 60 times a second,
	// this tick needs to be called 8 times to match the right speed
	for i := 0; i < 8; i++ {
		Tick()
	}
}
