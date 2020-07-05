package chip

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var bufferedStream *beep.Buffer

func init() {

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
	bufferedStream = beep.NewBuffer(format)
	bufferedStream.Append(streamer)

	// Divided by zero to match update function
	// Can be adjusted as see fit
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/60))
}

func Play() {
	speaker.Play(bufferedStream.Streamer(0, bufferedStream.Len()))
}

func Stop() {
	speaker.Clear()
}
