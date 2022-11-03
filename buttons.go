package main

import "github.com/gen2brain/iup-go/iup"

func buttonsFrame() iup.Ihandle {
	return iup.Frame(
		iup.Hbox(
			iup.Vbox(
				iup.Toggle("Show Profile Button").SetAttribute("VALUE", ooState(config.buttons.profileButtonEnabled)),
				iup.Toggle("Show Song Button").SetAttribute("VALUE", ooState(config.buttons.songButtonEnabled)),
				iup.Fill(),
			),
			iup.Fill(),
		),
	).SetAttributes(`TITLE="Buttons"`)
}
