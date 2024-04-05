package gui

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"slices"

	"github.com/go-p5/p5"
	"github.com/zivlakmilos/live-midi/internal/midinote"
	"github.com/zivlakmilos/live-midi/resources"
)

type SheetWidget struct {
	x     float64
	y     float64
	clefG image.Image
	clefF image.Image
	notes []uint8
}

func NewSheetWidget(x, y float64) *SheetWidget {
	return &SheetWidget{
		x:     x,
		y:     y,
		notes: []uint8{},
	}
}

func (w *SheetWidget) Setup() {
	clefG, _, err := image.Decode(bytes.NewReader(resources.IconClefG))
	if err != nil {
		log.Fatalf("errro: %v", err)
	}
	w.clefG = clefG

	clefF, _, err := image.Decode(bytes.NewReader(resources.IconClefF))
	if err != nil {
		log.Fatalf("errro: %v", err)
	}
	w.clefF = clefF
}

func (w *SheetWidget) Draw() {
	w.drawLines()

	w.drawClef(w.clefG, 0, w.y+45, 0.17)
	w.drawClef(w.clefF, 10, w.y+165, 0.1)

	for _, note := range w.notes {
		w.drawNote(note)
	}
}

func (w *SheetWidget) NoteOn(key uint8) {
	w.notes = append(w.notes, key)
}

func (w *SheetWidget) NoteOff(key uint8) {
	w.notes = slices.DeleteFunc(w.notes, func(e uint8) bool {
		return key == e
	})
}

func (w *SheetWidget) drawLines() {
	for i := 0; i < 5; i++ {
		x := w.x
		y := w.y + 40 + 20*float64(i)
		p5.Line(x, y, x+200, y)
	}

	for i := 0; i < 5; i++ {
		x := w.x
		y := w.y + 160 + 20*float64(i)
		p5.Line(x, y, x+200, y)
	}
}

func (w *SheetWidget) drawClef(img image.Image, x, y, scale float64) {
	p5.Push()
	p5.Translate(x, y)
	p5.Scale(scale, scale)
	p5.DrawImage(img, 0, 0)
	p5.Pop()
}

func (w *SheetWidget) drawNote(note uint8) {
	note, octave := midinote.Normalize(note)

	x := w.x + 120
	y := w.y + 140 - float64(note)*10 - float64(octave)*70
	width := 30.0
	height := 20.0
	p5.Fill(color.Black)
	p5.Ellipse(x, y, width, height)

	if note%2 == 0 {
		p5.Line(x-25, y, x+25, y)
	}
}
