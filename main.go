package main

import (
	"github.com/gen2brain/iup-go/iup"
	log "github.com/sirupsen/logrus"
	"path/filepath"
)

var albumDefaultPreviewPath, _ = filepath.Abs("./assets/lfm_logo.png")

var config = configuration{
	app: appConfiguration{
		title:     "Last.FM Discord RPC",
		version:   version,
		discordID: defaultDiscordId,
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
	state: true,
}

const globalColumnSize = "100x"

func main() {
	iup.Open()
	defer iup.Close()

	log.SetLevel(log.TraceLevel)
	log.SetReportCaller(true)

	menu().SetHandle("menu")

	vbox := iup.Vbox(
		previewFrame(),
		presenceSettingsFrame(),
	).SetAttributes("MARGIN=4x5, GAP=5")
	dlg := iup.Dialog(vbox)

	dlg.SetAttributes(`TITLE="` + config.app.title + ` - ` + version + `", MENU=menu, RESIZE=NO,`)

	iup.Show(dlg)
	iup.MainLoop()

}
