package main

import log "github.com/sirupsen/logrus"

func changeState() error {

	config.state = !config.state
	log.Trace("State changed to: ", config.state)
	return nil
}

func kill() error {
	config.state = false
	return nil
}
