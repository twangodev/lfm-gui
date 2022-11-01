package main

import "github.com/gen2brain/iup-go/iup"

const customApplicationSize = "100x"

func customApplicationFrame() iup.Ihandle {
	return iup.Frame(iup.Vbox(
		iup.Label("Application ID:"),
		iup.Text().SetAttributes(`VALUE="`+config.app.discordID+`", SIZE=`+customApplicationSize),
		iup.Button("Apply Changes").SetAttributes(`SIZE=`+customApplicationSize),
	)).SetAttributes(`TITLE="Custom Application Settings"`)
}
