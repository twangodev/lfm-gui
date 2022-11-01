package main

import "github.com/gen2brain/iup-go/iup"

func settingsFrame() iup.Ihandle {
	return iup.Frame(
		iup.Vbox(
			iup.Toggle("Close to tray"),
			iup.Toggle("Run on Startup"),
			iup.Toggle("Use Custom Discord\nApplication"),
			customApplicationFrame(),
		),
	).SetAttributes(`TITLE=Settings, SIZE=` + globalColumnSize)
}
