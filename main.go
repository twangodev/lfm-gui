package main

import (
	"github.com/gen2brain/iup-go/iup"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
	"time"
)

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
		profileButton:         true,
		profileButtonText:     "Visit Last.FM profile",
		songButton:            true,
		previewSongButtonText: "View scrobble on Last.FM",
	},
	state:    true,
	logLevel: 0,
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
		generateLogContext("presenceCycle").Trace("Starting presence cycle")
		updatePresence(false)
		time.Sleep(time.Duration(config.refreshTime) * time.Millisecond)
	}
}

func elapsedCycle() {
	for {
		generateLogContext("elapsedCycle").Trace("Starting elapsed cycle")
		if config.state {
			go updateElapsed()
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {

	// Windows Colorful Logs
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetLevel(log.TraceLevel)
	log.SetOutput(colorable.NewColorableStdout())
	mainContext := generateLogContext("main").WithFields(log.Fields{
		"name":    config.app.title,
		"version": version,
		"appID":   config.app.discordID,
	})
	mainContext.Info("Starting Application")

	if defaultDiscordId == config.app.discordID {
		mainContext.Debugln("Detected default Discord Application ID")
	} else {
		mainContext.Info("Detected custom Discord Application ID")
	}

	// Start ongoing background schedulers
	mainContext.Trace("Starting channelReceiver")
	go channelReceiver()
	mainContext.Trace("Starting presenceCycle")
	go presenceCycle()
	mainContext.Trace("Starting elapsedCycle")
	go elapsedCycle()

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

	if config.username == "" {
		mainContext.Warnln("No username set")
		go iup.Message("No username set", "Please set your Last.FM username in the settings menu")
	}

	iup.MainLoop()
}
