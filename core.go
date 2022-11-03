package main

import (
	"github.com/gen2brain/iup-go/iup"
	log "github.com/sirupsen/logrus"
)

func changeState() {
	if config.state {
		sendKillSignal()
	} else {
		sendStartSignal()
	}

	config.state = !config.state
	log.Trace("Updating control frame button states")
	iup.GetHandle("controlButton").SetAttribute("TITLE", enabledDisabled(config.state))
	iup.GetHandle("forceUpdateButton").SetAttribute("ACTIVE", ooState(config.state))
	iup.GetHandle("forceReconnectButton").SetAttribute("ACTIVE", ooState(config.state))
}

func reconnect() {
	sendKillSignal()
	sendStartSignal()
}
