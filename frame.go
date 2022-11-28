package main

import (
	"fmt"
	"lfm-gui/concurrency"
)

type frame struct {
	opcode concurrency.Opcode
	data   interface{}
}

func (f frame) String() string {
	return fmt.Sprintf("{opcode: %s, data: %v}", f.opcode, f.data)
}
