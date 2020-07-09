// Package chip provide functionality for the chip-8 interpreter
package chip

import (
	"fmt"
	"os"
)

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

	if len(os.Args) < 2 {
		fmt.Println("Usage: chip8 [file]")
		os.Exit(0)
	}
	fmt.Println(os.Args[1])
	Load(os.Args[1])
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

	IncTime()
}
