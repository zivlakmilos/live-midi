package gui

import (
	"fyne.io/fyne/v2"
)

type MainWindow struct {
	app    fyne.App
	window fyne.Window
}

func NewMainWindow(app fyne.App) *MainWindow {
	w := app.NewWindow("Midi View")
	w.Resize(fyne.NewSize(100, 100))

	return &MainWindow{
		app:    app,
		window: w,
	}
}

func (w *MainWindow) Show() {
	w.window.Show()
}
