package main

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	"github.com/xtlx2000/golib/log"
)

func PlayMp3(filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		log.Errorf("os open error: %v", err)
		return err
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Errorf("mp3 decode error: %v", err)
		return err
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done

	time.Sleep(time.Millisecond * 200)
	return nil
}
