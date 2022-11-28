package main

import (
	"github.com/gen2brain/iup-go/iup"
)

const githubUrl = "https://github.com/twangodev/lfm-gui"

func usernameCallback(ih iup.Ihandle) int {
	return promptUsername()
}

func githubCallback(ih iup.Ihandle) int {
	openBrowser(githubUrl)
	return iup.DEFAULT
}

func menu() iup.Ihandle {
	return iup.Menu(
		iup.Submenu("File", iup.Menu(
			iup.Item("Import"),
			iup.Item("Export"),
			iup.Item("Save\tCtrl+S"),
			iup.Item("Auto Save"),
			iup.Separator(),
			iup.Item("Exit\tCtrl+Q").SetCallback("ACTION", iup.ActionFunc(exit)),
		)),
		iup.Submenu("Edit", iup.Menu(
			iup.Item("Fields"),
			iup.Item("Blacklist"),
		)),
		iup.Submenu("Settings", iup.Menu(
			iup.Item("Update Username").SetCallback("ACTION", iup.ActionFunc(usernameCallback)),
			iup.Separator(),
			iup.Submenu("Logs", iup.Menu(
				iup.Item("Open Log File"),
				iup.Item("Configure Log Level"),
			)),
		)),
		iup.Submenu("Help", iup.Menu(
			iup.Item("About"),
			iup.Item("View Documentation"),
			iup.Item("Github").SetCallback("ACTION", iup.ActionFunc(githubCallback)),
		)),
	)
}
