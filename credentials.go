package main

import "github.com/gen2brain/iup-go/iup"

var loginFrameSize = "75x"

func credentialsFrame() iup.Ihandle {
	return iup.Frame(
		iup.Vbox(
			iup.Label("Last.FM Username:").SetAttributes(`VALUE="`+config.username+`", SIZE=`+loginFrameSize),
			iup.Text().SetAttribute("SIZE", loginFrameSize),
			iup.Button("Apply").SetAttribute("SIZE", loginFrameSize),
			iup.Fill(),
		),
	).SetAttributes("TITLE=Credentials, SIZE=" + globalColumnSize)

}
