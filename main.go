package main

import (
	"github.com/gen2brain/iup-go/iup"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"time"
)

var albumDefaultPreviewPath, _ = filepath.Abs("./assets/lfm_logo.png")

var config = configuration{
	app: appConfiguration{
		title:     "Last.FM Discord RPC",
		discordID: defaultDiscordId,
	},
	username: "",
	preview: previewConfiguration{
		enabled: true,
		albumConfig: albumConfiguration{
			cover:                 "$COVERURL",
			hover:                 true,
			albumDefaultHoverText: "$ALBUM",
		},
		smallImageConfig: smallImageConfiguration{
			enabled:              true,
			smallImageDefaultKey: "lfm_logo",
			smallImageHover:      true,
			smallImageHoverText:  "lfm-gui - " + version,
			lovedEnabled:         true,
		},
		albumDefaultPreviewPath: albumDefaultPreviewPath,
	},
	refreshTime: 12000,
	rows: rowsConfiguration{
		rowOne:      true,
		rowOneText:  "$TRACK",
		rowTwo:      true,
		rowTwoText:  "by $ARTIST",
		timeElapsed: true,
	},
	buttons: buttonsConfiguration{
		profileButton:     true,
		profileButtonText: "Visit Last.FM profile",
		songButton:        true,
		songButtonText:    "View scrobble on Last.FM",
	},
	state: true,
}

const globalColumnSize = "100x"

func channelReceiver() {
	for {
		msg := <-channel
		handleMsg(msg)
	}
}

func presenceCycle() {
	for {
		updatePresence()
		time.Sleep(time.Duration(config.refreshTime) * time.Millisecond)
	}
}

func main() {

	log.SetLevel(log.TraceLevel)
	// Windows Colorful Logs
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(colorable.NewColorableStdout())
	mainContext := log.WithField("ctx", "main")
	mainContext.WithFields(log.Fields{
		"name":    config.app.title,
		"version": version,
		"appID":   config.app.discordID,
	}).Info("Starting Application")

	if defaultDiscordId == config.app.discordID {
		mainContext.Debugln("Using default Discord Application ID")
	} else {
		mainContext.Info("Detected custom Discord Application ID")
	}

	if config.username == "" {
		mainContext.Warnln("No username set")

	}

	mainContext.Trace("Starting channelReceiver")
	go channelReceiver()
	mainContext.Trace("Starting presenceCycle")
	go presenceCycle()

	mainContext.Trace("Starting IUP")
	iup.Open()
	defer iup.Close()

	mainContext.Trace("Creating main dialog")
	menu().SetHandle("menu")

	vbox := iup.Vbox(
		previewFrame(),
		presenceSettingsFrame(),
	).SetAttributes("MARGIN=4x5, GAP=5")
	dlg := iup.Dialog(vbox)

	dlg.SetAttributes(`TITLE="` + config.app.title + ` - ` + version + `", MENU=menu, RESIZE=NO,`)

	mainContext.Trace("Showing main dialog")
	iup.Show(dlg)
	iup.MainLoop()

}
