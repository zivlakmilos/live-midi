package main

import (
	"log"

	"github.com/go-p5/p5"
	"github.com/zivlakmilos/live-midi/internal/gui"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

func main() {
	defer midi.CloseDriver()

	w := gui.NewMainWindow()

	go func() {
		ports := midi.GetInPorts()
		midiW := gui.NewMidiWindow(ports, w)
		err := midiW.Run()
		if err != nil {
			log.Fatalf("error: %v", err)
			return
		}
	}()

	p5.Run(w.Setup, w.Draw)
}
