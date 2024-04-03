package gui

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/go-p5/p5"
	"github.com/zivlakmilos/live-midi/resources"
)

type SheetWidget struct {
	x     float64
	y     float64
	clefG image.Image
	clefF image.Image
}

func NewSheetWidget(x, y float64) *SheetWidget {
	return &SheetWidget{
		x: x,
		y: y,
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
	w.drawClef(w.clefG, 0, w.y+45, 0.17)
	w.drawClef(w.clefF, 10, w.y+165, 0.1)

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
