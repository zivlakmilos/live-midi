package main

import (
	"log"

	"github.com/go-p5/p5"
	"github.com/zivlakmilos/live-midi/internal/gui"
)

func main() {
	go func() {
		midiW := gui.NewMidiWindow()
		err := midiW.Run()
		if err != nil {
			log.Fatalf("error: %v", err)
			return
		}
	}()

	w := gui.NewMainWindow()
	p5.Run(w.Setup, w.Draw)
}
