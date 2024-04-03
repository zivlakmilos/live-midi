package main

import (
	"github.com/go-p5/p5"
	"github.com/zivlakmilos/live-midi/internal/gui"
)

func main() {
	w := gui.NewMainWindow()
	p5.Run(w.Setup, w.Draw)
}
