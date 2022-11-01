package main

import "github.com/gen2brain/iup-go/iup"

func menu() iup.Ihandle {
	return iup.Menu(
		iup.Submenu("File", iup.Menu(
			iup.Item("Import"),
			iup.Item("Export"),
			iup.Separator(),
			iup.Item("Exit").SetCallback("ACTION", exit),
		)),
		iup.Submenu("Help", iup.Menu(
			iup.Item("View Documentation"),
			iup.Item("Github"),
			iup.Item("Logs"),
		)),
	)
}
