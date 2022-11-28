package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/skratchdot/open-golang/open"
)

func generateLogContext(src string) *log.Entry {
	return log.WithField("src", src)
}

func ooState(boolean bool) string {
	if boolean {
		return "ON"
	}
	return "OFF"
}

func ooBoolean(state string) bool {
	if state == "ON" {
		return true
	}
	return false
}

func ynState(boolean bool) string {
	if boolean {
		return "YES"
	}
	return "NO"
}

func openBrowser(url string) {
	logContext := log.WithField("url", url)
	logContext.Info("Opening browser")
	err := open.Start(url)
	if err != nil {
		logContext.WithField("error", err).Error("Failed to open browser")
	}
}
