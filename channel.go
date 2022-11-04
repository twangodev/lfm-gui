package main

import log "github.com/sirupsen/logrus"

var channel = make(chan frame)

func sendStartSignal() {
	log.Trace("Sending start signal")
	channel <- frame{opcode: STATE_UPDATE, data: true}
}

func sendKillSignal() {
	log.Trace("Sending kill signal")
	channel <- frame{opcode: STATE_UPDATE, data: false}
}

func updatePresence() {
	log.Trace("Sending presence update signal")
	channel <- frame{opcode: PRESENCE_UPDATE, data: nil}
}

func handleMsg(msg frame) {

}
