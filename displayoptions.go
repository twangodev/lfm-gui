package main

import (
	"github.com/gen2brain/iup-go/iup"
)

var displayOptionsSize = "70x"

func presenceSettingsFrame() iup.Ihandle {
	return iup.Frame(
		iup.Hbox(
			albumPresenceFrame(),
			contentPresenceFrame(),
		),
	).SetAttributes(`TITLE="Display Options", SIZE=` + globalColumnSize)
}
