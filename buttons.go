package main

import (
	"github.com/gen2brain/iup-go/iup"
	"lfm-gui/concurrency"
)

func buttonsFrame() iup.Ihandle {
	return iup.Frame(
		iup.Hbox(
			iup.Vbox(
				iup.Toggle("Show Profile Button").
					SetAttribute("VALUE", ooState(config.buttons.profileButton)).
					SetCallback("VALUECHANGED_CB",
						updateConfigBoolCallback(concurrency.DISP_BUTTONS_PROFILE, &config.buttons.profileButton),
					),
				iup.Toggle("Show Song Button").
					SetAttribute("VALUE", ooState(config.buttons.songButton)).
					SetCallback("VALUECHANGED_CB", updateConfigBoolCallback(concurrency.DISP_BUTTONS_SONG, &config.buttons.songButton)),
				iup.Fill(),
			),
			iup.Fill(),
		),
	).SetAttributes(`TITLE="Buttons"`)
}
