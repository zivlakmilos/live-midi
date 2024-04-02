package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/zivlakmilos/live-midi/internal/gui"
)

func main() {
	a := app.New()

	w := gui.NewMainWindow(a)
	w.Show()

	a.Run()
}
