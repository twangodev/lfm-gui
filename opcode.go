package main

var PRESENCE_UPDATE = opcode{
	code:   0,
	strRep: "PRESENCE_UPDATE",
}
var STATE_UPDATE = opcode{
	code:   1,
	strRep: "STATE_UPDATE",
}

type opcode struct {
	code   int
	strRep string
}

func (o opcode) String() string {
	return o.strRep
}
