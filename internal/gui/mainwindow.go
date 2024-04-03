package gui

import "github.com/go-p5/p5"

type MainWindow struct {
	sheetWidget *SheetWidget
}

func NewMainWindow() *MainWindow {
	return &MainWindow{
		sheetWidget: NewSheetWidget(10, 10),
	}
}

func (w *MainWindow) Setup() {
	p5.Canvas(300, 300)
	w.sheetWidget.Setup()
}

func (w *MainWindow) Draw() {
	w.sheetWidget.Draw()
}
