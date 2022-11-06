package main

import (
	lfm "github.com/lastfm-discordrpc/lfm-api"
	log "github.com/sirupsen/logrus"
)

func updatePresenceReceiver() {
	currentScrobble, err := lfm.GetActiveScrobble(config.username)
	logContext := log.WithFields(log.Fields{
		"scrobble": currentScrobble,
		"username": config.username,
	})
	if err != nil {
		logContext.Warnln("Failed to get active scrobble:", err)
		return
	}

	if currentScrobble.Active {

	}

}
