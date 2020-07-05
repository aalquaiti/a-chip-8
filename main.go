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
		windowHeight, update)

	for y := 0; y < screenHeight; y++ {
		for x := 0; x < screenWidth; x++ {
			if y == (screenHeight/2)-1 || y == screenHeight/2 {
				screen.Draw(x, y, true)
			}
		}
	}

	// f, err := os.Open("beep.mp3")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// streamer, format, err := mp3.Decode(f)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer streamer.Close()
	// speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	// loop := beep.Loop(-1, streamer)
	// speaker.Play(loop)

	// chip.Play()
	screen.Show()
}

var playing bool = false
var t int = 0

func update() {

	chip.Play()

	// if !playing && t < 7 {
	// 	chip.Play()
	// 	playing = true
	// }

	// if playing && t >= 7 {
	// 	chip.Stop()
	// 	playing = false
	// }

	// inst := chip.Inst
	// inst[0xE](0, 9, 0xE)
	// t++
	// t = t % 14
}
