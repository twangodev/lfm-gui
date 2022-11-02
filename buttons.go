package main

import "github.com/gen2brain/iup-go/iup"

func buttonsFrame() iup.Ihandle {
	return iup.Frame(
		iup.Hbox(
			iup.Vbox(
				iup.Toggle("Show Profile Button"),
				iup.Toggle("Show Song Button"),
				iup.Fill(),
			),
			iup.Fill(),
		),
	).SetAttributes(`TITLE="Buttons"`)
}
