package main

import (
	"github.com/gen2brain/iup-go/iup"
	"path/filepath"
)

const version = "DEV"

var albumDefaultPreviewPath, _ = filepath.Abs("./assets/lfm_logo.png")

var config = configuration{
	app: appConfiguration{
		title:     "Last.FM Discord RPC",
		version:   version,
		discordID: "970003417277812736",
	},
	username: "",
	preview: previewConfiguration{
		enabled: true,
		albumConfig: albumConfiguration{
			cover:                    "$COVERURL",
			albumDefaultHoverEnabled: true,
			albumDefaultHoverText:    "$ALBUM",
		},
		smallImageConfig: smallImageConfiguration{
			enabled:                true,
			smallImageDefaultKey:   "lfm_logo",
			smallImageHoverEnabled: true,
			smallImageHoverText:    "lfm-gui - " + version,
			lovedEnabled:           true,
			lovedKey:               "heart",
		},
		albumDefaultPreviewPath: albumDefaultPreviewPath,
	},
	refreshTime: 12000,
	rows: rowsConfiguration{
		rowOneEnabled:      true,
		rowOne:             "$TRACK",
		rowTwoEnabled:      true,
		rowTwo:             "by $ARTIST",
		timeElapsedEnabled: true,
	},
	buttons: buttonsConfiguration{
		profileButtonEnabled: true,
		profileButton:        "Visit Last.FM profile",
		songButtonEnabled:    true,
		songButton:           "View scrobble on Last.FM",
	},
}

const globalColumnSize = "100x"

func main() {
	iup.Open()
	defer iup.Close()

	menu().SetHandle("menu")

	hboxR1 := iup.Hbox(
		credentialsFrame(),
		settingsFrame(),
		aboutFrame(),
	)

	vbox := iup.Vbox(
		hboxR1,
		previewFrame(),
		presenceSettingsFrame(),
	).SetAttributes("MARGIN=4x5, GAP=5")
	dlg := iup.Dialog(vbox)

	dlg.SetAttributes(`TITLE="` + config.app.title + `", MENU=menu, RESIZE=NO,`)

	iup.Show(dlg)
	iup.MainLoop()
}
