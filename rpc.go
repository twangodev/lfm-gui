package main

import (
	"github.com/hugolgst/rich-go/client"
	"github.com/hugolgst/rich-go/ipc"
	log "github.com/sirupsen/logrus"
)

var loggedIn = false

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
	getLogContext().Infoln("Logged in successfully")
}
