package main

import (
	"github.com/gen2brain/iup-go/iup"
)

const controlButtonSize = "100x"

func enabledDisabled(boolean bool) string {
	if boolean {
		return "Disable"
	}
	return "Enable"
}

func enableDisableCallback(ih iup.Ihandle) int {
	changeState()
	return iup.DEFAULT
}

func forceUpdateCallback(ih iup.Ihandle) int {
	updatePresence(true)
	return iup.DEFAULT
}

func forceReconnectCallback(ih iup.Ihandle) int {
	reconnect()
	return iup.DEFAULT
}

func controlFrame() iup.Ihandle {
	return iup.Frame(
		iup.Hbox(
			iup.Vbox(
				iup.Button(enabledDisabled(config.state)).SetAttribute("SIZE", controlButtonSize).SetCallback("ACTION", iup.ActionFunc(enableDisableCallback)).SetHandle("controlButton"),
				iup.Button("Force Update").SetAttribute("SIZE", controlButtonSize).SetCallback("ACTION", iup.ActionFunc(forceUpdateCallback)).SetHandle("forceUpdateButton"),
				iup.Button("Force Reconnect").SetAttribute("SIZE", controlButtonSize).SetCallback("ACTION", iup.ActionFunc(forceReconnectCallback)).SetHandle("forceReconnectButton"),
				iup.Fill(),
			),
			iup.Fill(),
		),
	).SetAttributes("TITLE=Controls")
}
