package main

import (
	"github.com/gen2brain/iup-go/iup"
	"github.com/hugolgst/rich-go/client"
	log "github.com/sirupsen/logrus"
	lfm "github.com/twangodev/lfm-api"
	"strings"
	"time"
)

func promptUsername() int {
	dlg := iup.Dialog(
		iup.Vbox(
			iup.Hbox(
				iup.Fill(),
				iup.Label("Username:"),
				iup.Text().SetAttributes(`CONFIG="`+config.username+`", SIZE=75x`).SetHandle("username"),
				iup.Fill(),
			),
			iup.Hbox(
				iup.Fill(),
				iup.Button("OK").SetAttribute("SIZE", "40x").
					SetCallback("ACTION", iup.ActionFunc(func(ih iup.Ihandle) int {
						rawUsername := iup.GetHandle("username").GetAttribute("VALUE")
						trimmed := strings.TrimSpace(rawUsername)
						if trimmed == "" {
							log.Error("Username cannot be empty")
							go iup.Message("Error", "Username is empty")
							return iup.CLOSE
						}
						if strings.Contains(trimmed, " ") {
							log.Error("Username contains spaces, skipping update and sending message")
							go iup.Message("Invalid Username", "Username cannot contain spaces")
							return iup.DEFAULT
						}
						config.username = trimmed
						log.Info("Username updated to: ", trimmed)
						go updatePresence(true)
						return iup.CLOSE
					})),
				iup.Button("Cancel").SetAttribute("SIZE", "40x").SetCallback("ACTION", iup.ActionFunc(exit)),
				iup.Fill(),
			),
		).SetAttributes(`MARGIN=4x5, GAP=5`),
	)
	iup.Popup(dlg, iup.CENTER, iup.CENTER)
	return iup.DEFAULT
}

func safeSetActivity(activity client.Activity) {
	logContext := generateLogContext("safeSetActivity").WithField("activity", activity)

	defer func() {
		if r := recover(); r != nil {
			logContext.WithField("panic", r).Error("Panic while updating presence. Restart recommended")
		}
	}()

	err := client.SetActivity(activity)
	if err != nil {
		logContext.Error("Presence failed to update")
	} else {
		logContext.Info("Presence updated")
		updateRpcPreview(activity)
	}

}

func updatePresenceReceiver(force bool) {

	logContext := generateLogContext("updatePresenceReceiver").WithFields(log.Fields{
		"force": force, "username": config.username, "state": config.state,
	})

	if !config.state {
		logContext.Debug("State is disabled, skipping update")
		return
	}

	if config.username == "" {
		logContext.Warn("No username set, skipping presence update")
		return
	}

	currentScrobble, err := lfm.GetActiveScrobble(config.username)
	logContext.WithFields(log.Fields{
		"scrobble": currentScrobble,
		"active":   currentScrobble.Active,
	}).Debugln("Received scrobble")
	logContext = logContext.WithField("currentScrobble", currentScrobble)
	if err != nil {
		logContext.Warnln("Failed to get active scrobble:", err)
		return
	}

	if currentScrobble.Active {
		if !loggedIn {
			logContext.Info("New scrobble detected. Logging in")
			login()
		}
	} else {
		if loggedIn {
			logContext.Info("No scrobble detected. Logging out")
			logout()
			updateRpcPreviewImage("") // Reset image to default
		} else { // Retain logout state
			logContext.Trace("No new scrobble detected")
		}
		return
	}

	if timestamp != currentScrobble.DataTimestamp || force { // Update old scrobble to match current scrobble
		timestamp = currentScrobble.DataTimestamp
		logContext.Debug("Updating presence")
	} else {
		return
	}

	baseActivity := createActivity(currentScrobble, false)
	safeSetActivity(baseActivity)

	buttonedActivity := createActivity(currentScrobble, true)
	safeSetActivity(buttonedActivity)
}

func elapsedTimeReceiver() {
	elapsed := time.Since(timestamp)
	minutes := int(elapsed.Minutes())
	seconds := int(elapsed.Seconds()) % 60
	generateLogContext("elapsedTimeReceiver").WithFields(log.Fields{
		"minutes": minutes, "seconds": seconds},
	).Trace("Updating elapsed time")
	if elapsed < time.Second {
		return
	}
	go updateRpcPreviewElapsed(minutes, seconds)
}
