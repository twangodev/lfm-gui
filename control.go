package main

import "github.com/gen2brain/iup-go/iup"

const controlButtonSize = "100x"

func controlFrame() iup.Ihandle {
	return iup.Frame(
		iup.Hbox(
			iup.Vbox(
				iup.Button("Enable").SetAttribute("SIZE", controlButtonSize),
				iup.Button("Force Update").SetAttribute("SIZE", controlButtonSize),
				iup.Button("Force Reconnect").SetAttribute("SIZE", controlButtonSize),
				iup.Fill(),
			),
			iup.Fill(),
		),
	).SetAttributes("TITLE=Controls")
}
