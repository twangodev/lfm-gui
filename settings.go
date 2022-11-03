package main

import "github.com/gen2brain/iup-go/iup"

func settingsFrame() iup.Ihandle {
	return iup.Frame(
		iup.Hbox(
			iup.Toggle("Close to tray"),
			iup.Toggle("Run on Startup"),
			iup.Fill(),
		),
	).SetAttributes(`TITLE=Settings, SIZE=` + globalColumnSize)
}
