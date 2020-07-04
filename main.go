package main

import (
	"fmt"

	"github.com/aalquaiti/a-chip-8/chip"
)

var pixels []byte

const (
	screenWidth  = 64
	screenHeight = 32
	windowWidth  = 640
	windowHeight = 320
)

func main() {
	fmt.Println("Hello from a-chip-8")
	pixels = make([]byte, screenWidth*screenHeight*4)
	screen := chip.NewScreen(screenWidth, screenHeight, windowWidth,
		windowHeight, func() {})

	// for y := 0; y < screenHeight; y++ {
	// 	for x := 0; x < screenWidth; x++ {
	// 		if y == (screenHeight/2)-1 || y == screenHeight/2 {
	// 			position := (screenWidth * y * 4) + (x * 4)
	// 			pixels[position] = 0xFF
	// 			pixels[position+1] = 0xFF
	// 			pixels[position+2] = 0xFF
	// 			pixels[position+3] = 0xFF
	// 		}
	// 	}
	// }

	for y := 0; y < screenHeight; y++ {
		for x := 0; x < screenWidth; x++ {
			if y == (screenHeight/2)-1 || y == screenHeight/2 {
				screen.Draw(x, y, true)
			}
		}
	}

	// screen.SetPixels(pixels)
	screen.Show()

}

func update() {}
