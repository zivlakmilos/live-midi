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
)

type MidiWindow struct {
	win *app.Window
}

func NewMidiWindow() *MidiWindow {
	return &MidiWindow{
		win: app.NewWindow(
			app.Title("Live MIDI"),
			app.Size(unit.Dp(300), unit.Dp(300)),
		),
	}
}

func (w *MidiWindow) Run() error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops

	grp := &widget.Enum{}
	clickable := &widget.Clickable{}

	for {
		e := <-w.win.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			items := []layout.FlexChild{}
			items = append(items, layout.Rigid(material.RadioButton(th, grp, "midi1", "MIDI 1").Layout))
			items = append(items, layout.Rigid(material.RadioButton(th, grp, "midi2", "MIDI 2").Layout))
			items = append(items, layout.Rigid(material.RadioButton(th, grp, "midi3", "MIDI 3").Layout))

			items = append(items, layout.Rigid(layout.Spacer{Height: unit.Dp(25)}.Layout))
			items = append(items, layout.Rigid(material.Button(th, clickable, "Save").Layout))

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

			e.Frame(gtx.Ops)
		}
	}
}
