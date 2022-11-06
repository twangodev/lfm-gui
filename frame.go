package main

import "lfm-gui/concurrency"

type frame struct {
	opcode concurrency.Opcode
	data   interface{}
}
