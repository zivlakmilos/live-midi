package gui

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gitlab.com/gomidi/midi/v2"
)

type MidiWindow struct {
	win        *app.Window
	ports      midi.InPorts
	group      *widget.Enum
	btnConnect *widget.Clickable
	mainWindow *MainWindow
}

func NewMidiWindow(ports midi.InPorts, mainWindow *MainWindow) *MidiWindow {
	return &MidiWindow{
		win: app.NewWindow(
			app.Title("Live MIDI"),
			app.Size(unit.Dp(300), unit.Dp(300)),
		),
		ports:      ports,
		group:      &widget.Enum{},
		btnConnect: &widget.Clickable{},
		mainWindow: mainWindow,
	}
}

func (w *MidiWindow) Run() error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops

	for {
		e := <-w.win.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			items := []layout.FlexChild{}
			for _, port := range w.ports {
				items = append(items, layout.Rigid(material.RadioButton(th, w.group, port.String(), port.String()).Layout))
			}

			items = append(items, layout.Rigid(layout.Spacer{Height: unit.Dp(25)}.Layout))
			items = append(items, layout.Rigid(material.Button(th, w.btnConnect, "Connect").Layout))

			layout.Inset{
				Top:    unit.Dp(25),
				Right:  unit.Dp(25),
				Bottom: unit.Dp(25),
				Left:   unit.Dp(25),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:    layout.Vertical,
					Spacing: layout.SpaceEnd,
				}.Layout(gtx, items...)
			})

			if w.btnConnect.Clicked() {
				w.mainWindow.SetupMidi(w.group.Value)
			}

			e.Frame(gtx.Ops)
		}
	}
}
