package gui

import "github.com/go-p5/p5"

type MainWindow struct{}

func NewMainWindow() *MainWindow {
	return &MainWindow{}
}

func (w *MainWindow) Setup() {
	p5.Canvas(300, 300)
}

func (w *MainWindow) Draw() {
}
