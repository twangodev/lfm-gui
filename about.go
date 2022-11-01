package main

import "github.com/gen2brain/iup-go/iup"

var aboutButtonSize = "75x"

func aboutFrame() iup.Ihandle {
	return iup.Frame(
		iup.Vbox(
			iup.Label(config.app.title+"\n"+version),
			iup.Button("Report a Bug").SetAttribute("SIZE", aboutButtonSize),
			iup.Button("Check for Updates").SetAttribute("SIZE", aboutButtonSize),
			iup.Fill(),
		),
	).SetAttributes("TITLE=About, SIZE=" + globalColumnSize)
}
