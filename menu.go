package main

import "github.com/gen2brain/iup-go/iup"

func menu() iup.Ihandle {
	return iup.Menu(
		iup.Submenu("File", iup.Menu(
			iup.Item("Import"),
			iup.Item("Export"),
			iup.Item("Auto Save"),
			iup.Separator(),
			iup.Item("Exit").SetCallback("ACTION", exit),
		)),
		iup.Submenu("Edit", iup.Menu(
			iup.Item("Fields"),
			iup.Item("Blacklist"),
		)),
		iup.Submenu("Last.FM", iup.Menu(
			iup.Item("Update Username"),
			iup.Separator(),
			iup.Item("Custom Discord Application"),
			iup.Item("Configure Custom Application").SetAttribute("ACTIVE", ynState(config.app.discordID != defaultDiscordId)),
		)),
		iup.Submenu("Settings", iup.Menu(
			iup.Item("Close to Tray"),
			iup.Item("Run on Startup"),
		)),
		iup.Submenu("Help", iup.Menu(
			iup.Item("About"),
			iup.Item("View Documentation"),
			iup.Item("Github"),
			iup.Item("Logs"),
		)),
	)
}
