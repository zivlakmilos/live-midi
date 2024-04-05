package gui

import (
	"sync"

	"github.com/go-p5/p5"
	"gitlab.com/gomidi/midi/v2"
)

type MainWindow struct {
	sheetWidget *SheetWidget
	mutex       *sync.Mutex
}

func NewMainWindow() *MainWindow {
	return &MainWindow{
		sheetWidget: NewSheetWidget(10, 10),
		mutex:       &sync.Mutex{},
	}
}

func (w *MainWindow) Setup() {
	p5.Canvas(300, 300)
	w.sheetWidget.Setup()
}

func (w *MainWindow) Draw() {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.sheetWidget.Draw()
}

func (w *MainWindow) SetupMidi(portName string) {
	port, err := midi.FindInPort(portName)
	if err != nil {
		return
	}

	midi.ListenTo(port, func(msg midi.Message, timestampms int32) {
		w.mutex.Lock()
		defer w.mutex.Unlock()

		var key uint8

		switch {
		case msg.GetNoteStart(nil, &key, nil):
			w.sheetWidget.NoteOn(key)
		case msg.GetNoteEnd(nil, &key):
			w.sheetWidget.NoteOff(key)
		}
	}, midi.UseSysEx())
}
