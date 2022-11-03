package main

import (
	"github.com/gen2brain/iup-go/iup"
	log "github.com/sirupsen/logrus"
)

const controlButtonSize = "100x"

func enabledDisabled(boolean bool) string {
	if boolean {
		return "Enabled"
	}
	return "Disabled"
}

func enableDisableCallback(ih iup.Ihandle) int {
	err := changeState()
	if err != nil {
		log.Warn("Error changing state: ", err)
		err = kill()
		if err != nil {
			log.Error("Error killing:  ", err)
		}
	}
	ih.SetAttribute("TITLE", enabledDisabled(config.state))
	return iup.DEFAULT
}

func controlFrame() iup.Ihandle {
	return iup.Frame(
		iup.Hbox(
			iup.Vbox(
				iup.Button(enabledDisabled(config.state)).SetAttribute("SIZE", controlButtonSize).SetCallback("ACTION", iup.ActionFunc(enableDisableCallback)),
				iup.Button("Force Update").SetAttribute("SIZE", controlButtonSize),
				iup.Button("Force Reconnect").SetAttribute("SIZE", controlButtonSize),
				iup.Fill(),
			),
			iup.Fill(),
		),
	).SetAttributes("TITLE=Controls")
}
