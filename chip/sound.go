// Package chip provide functionality for the chip-8 interpreter
package chip

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// Sound used to play a beep or stop from playing one
type Sound struct {
	buffStream *beep.Buffer // Used to buffer beep file
}

// NewSound create a new reference for sound
// This method should only be called once in the life of the app
func NewSound() (s *Sound) {
	s = &Sound{}

	// Open the beep file
	f, err := os.Open("beep.mp3")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)

	}
	defer streamer.Close()

	// Buffer file into memory
	s.buffStream = beep.NewBuffer(format)
	s.buffStream.Append(streamer)

	// Divided by zero to match update function
	// Can be adjusted as see fit
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/60))

	return
}

// PlaySound plays a beep
func (s *Sound) PlaySound() {
	speaker.Play(s.buffStream.Streamer(0, s.buffStream.Len()))
}

// StopSound stops a beep
func (s *Sound) StopSound() {
	speaker.Clear()
}
