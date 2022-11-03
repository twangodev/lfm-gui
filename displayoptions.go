package main

import (
	"github.com/gen2brain/iup-go/iup"
)

var displayOptionsSize = "70x"

func presenceSettingsFrame() iup.Ihandle {
	return iup.Frame(
		iup.Hbox(
			iup.Vbox(
				albumPresenceFrame(),
				buttonsFrame(),
			),
			iup.Vbox(
				contentPresenceFrame(),
				controlFrame(),
			),
		),
	).SetAttributes(`TITLE="Display Options", SIZE=` + globalColumnSize)
}
