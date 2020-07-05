package main

import (
	"fmt"

	"github.com/aalquaiti/a-chip-8/chip"
)

var pixels []byte

func main() {
	fmt.Println("Hello from a-chip-8")

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
	chip.Start()
}

var playing bool = false
var t int = 0

func update() {

	chip.Tick()
}
