package main

import (
	"fmt"
	"github.com/hugolgst/rich-go/client"
	"github.com/hugolgst/rich-go/ipc"
	log "github.com/sirupsen/logrus"
	lfm "github.com/twangodev/lfm-api"
	"golang.org/x/net/html"
	"strings"
	"time"
)

var loggedIn = false
var timestamp = time.Now()

func getLogContext() *log.Entry {
	return log.WithField("loggedIn", loggedIn)
}

func login() {
	err := ipc.CloseSocket()
	if err != nil {
		getLogContext().Warnln("Failed to close IPC socket:", err)
	}
	err = client.Login(config.app.discordID)
	if err != nil {
		getLogContext().Warnln("Failed to login:", err)
		return
	}
	loggedIn = true
	getLogContext().Infoln("Logged in to Discord RPC")
}

func logout() {
	client.Logout()
	loggedIn = false
	timestamp = time.Now()
	getLogContext().Info("Logged out of Discord RPC")
}

func queryVariable(query string, scrobble lfm.Scrobble) string {
	switch query {
	case "$TRACK":
		return scrobble.Name
	case "$ARTIST":
		return scrobble.Artist
	case "$ALBUM":
		return scrobble.Album
	case "$COVERURL":
		return scrobble.CoverArtUrl
	default:
		return query
	}
}

func formatString(text string, scrobble lfm.Scrobble) string {
	var formattedText string
	for _, query := range strings.Split(strings.TrimSpace(text), " ") {
		formattedText += queryVariable(query, scrobble) + " "
	}
	return strings.TrimSpace(formattedText)
}

func createActivity(scrobble lfm.Scrobble, buttonsEnabled bool) client.Activity {

	logContext := generateLogContext("createActivity").WithFields(log.Fields{
		"scrobble":       scrobble,
		"buttonsEnabled": buttonsEnabled,
	})
	logContext.Trace("Building activity")

	activity := client.Activity{}

	if config.preview.enabled {
		largeImage := formatString(config.preview.albumConfig.cover, scrobble)
		logContext.WithField("largeImage", largeImage).Trace("Setting large image")
		activity.LargeImage = largeImage

		if config.preview.albumConfig.hover {
			largeText := formatString(config.preview.albumConfig.albumDefaultHoverText, scrobble)
			logContext.WithField("largeText", largeText).Trace("Setting large image hover text")
			activity.LargeText = largeText
		}

		if config.preview.smallImageConfig.enabled {
			smallImage := formatString(config.preview.smallImageConfig.smallImageDefaultKey, scrobble)
			logContext.WithField("smallImage", smallImage).Trace("Setting small image")
			activity.SmallImage = smallImage

			if config.preview.smallImageConfig.smallImageHover {
				smallText := formatString(config.preview.smallImageConfig.smallImageHoverText, scrobble)
				logContext.WithField("smallText", smallText).Trace("Setting small image hover text")
				activity.SmallText = smallText
			}

			if config.preview.smallImageConfig.lovedEnabled && scrobble.Loved {
				logContext.Trace("Overriding small image to loved")
				activity.SmallImage = "heart"
			}

		}

	}

	if config.rows.rowOne {
		rowOne := formatString(config.rows.rowOneText, scrobble)
		logContext.WithField("rowOne", rowOne).Trace("Setting row one")
		activity.Details = rowOne
	}

	if config.rows.rowTwo {
		rowTwo := formatString(config.rows.rowTwoText, scrobble)
		logContext.WithField("rowTwo", rowTwo).Trace("Setting row two")
		activity.State = rowTwo
	}

	if config.rows.timeElapsed {
		logContext.WithField("timestamp", scrobble.DataTimestamp).Trace("Setting timestamp")
		activity.Timestamps = &client.Timestamps{
			Start: &scrobble.DataTimestamp,
		}
	}

	var buttons []*client.Button
	if buttonsEnabled {
		profileButton := &client.Button{
			Label: config.buttons.profileButtonText,
			Url:   "https://www.last.fm/user/" + config.username,
		}

		dataLinkTitle := scrobble.DataLinkTitle
		dataLink := scrobble.DataLink
		if dataLinkTitle == "" {
			dataLinkTitle = config.buttons.previewSongButtonText
			dataLink = fmt.Sprintf("%vmusic/%v/%v", lfm.LastFmUrl, html.EscapeString(scrobble.Artist), html.EscapeString(scrobble.Name))
		}
		songButton := &client.Button{Label: dataLinkTitle, Url: dataLink}

		if config.buttons.profileButton {
			buttons = append(buttons, profileButton)
		}

		if config.buttons.songButton {
			buttons = append(buttons, songButton)
		}
	}
	activity.Buttons = buttons

	logContext.Trace("Built activity")
	return activity

}
