package main

import (
	"github.com/gen2brain/iup-go/iup"
	log "github.com/sirupsen/logrus"
	"lfm-gui/concurrency"
	"strconv"
)

var channel = make(chan frame)

func sendStartSignal() {
	channel <- frame{opcode: concurrency.STATE_UPDATE, data: true}
}

func sendKillSignal() {
	channel <- frame{opcode: concurrency.STATE_UPDATE, data: false}
}

func updatePresence(force bool) {
	channel <- frame{opcode: concurrency.REQUEST_PRESENCE_UPDATE, data: force}
}

func updateElapsed() {
	channel <- frame{opcode: concurrency.ELAPSED_UPDATE}
}

func sendConfigBoolFrame(o concurrency.Opcode, state bool) {
	channel <- frame{
		opcode: concurrency.CONFIG_UPDATE,
		data: frame{
			opcode: o,
			data:   state,
		},
	}
}

func updateConfigBoolCallback(opcode concurrency.Opcode, ref *bool,
	dependentHandles ...boolDependencyHandle,
) iup.ValueChangedFunc {
	return func(ih iup.Ihandle) int {
		state := ooBoolean(ih.GetAttribute("VALUE"))
		logContext := log.WithFields(log.Fields{
			"opcode":    opcode,
			"new_state": state,
			"original":  *ref,
		})
		logContext.Info("Updating boolean configuration")
		*ref = state
		sendConfigBoolFrame(opcode, state)

		for _, dependent := range dependentHandles {
			logContext.WithFields(log.Fields{
				"dependent": dependent.configOpcode,
				"parent":    opcode,
			}).
				Info("Updating boolean configuration dependent")
			*dependent.ref = state
			iup.SetAttribute(iup.GetHandle(dependent.handle), "ACTIVE", ynState(state))
			sendConfigBoolFrame(dependent.configOpcode, state)
		}

		return iup.DEFAULT
	}
}

func updateConfigIntCallback(opcode concurrency.Opcode, ref *int, min int, originalValue *int) iup.ValueChangedFunc {
	return func(ih iup.Ihandle) int {
		strVal := ih.GetAttribute("VALUE")
		val, err := strconv.Atoi(strVal)

		logContext := log.WithFields(log.Fields{
			"opcode":    opcode,
			"new_value": val,
			"original":  *ref,
			"min":       min,
		})

		if err != nil {
			logContext.Warn("Invalid value received, revert to original")
			val = *originalValue
		}

		if val <= min {
			logContext.Warn("Value too low, revert to minimum specified")
			val = min
		}

		ih.SetAttribute("VALUE", strconv.Itoa(val))
		if val == *originalValue {
			logContext.Trace("Value unchanged, no update required")
			return iup.DEFAULT
		}

		*ref = val
		channel <- frame{
			opcode: concurrency.CONFIG_UPDATE,
			data: frame{
				opcode: opcode,
				data:   val,
			},
		}
		return iup.DEFAULT
	}
}

func handleMsg(msg frame) {
	logContext := generateLogContext("channelReceiver").WithFields(log.Fields{
		"opcode": msg.opcode,
		"data":   msg.data,
	})
	logContext.Trace("Received frame")
	switch msg.opcode {
	case concurrency.STATE_UPDATE:
		if msg.data.(bool) {
			go updatePresence(true)
		} else {
			logout()
		}
	case concurrency.REQUEST_PRESENCE_UPDATE:
		go updatePresenceReceiver(msg.data.(bool))
	case concurrency.CONFIG_UPDATE:
		switch msg.data.(frame).opcode {
		default:
			go updatePresence(true)
		}
	case concurrency.ELAPSED_UPDATE:
		go elapsedTimeReceiver()
	default:
		logContext.Error("Unknown opcode received")
	}
}
