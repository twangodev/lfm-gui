package main

var PRESENCE_UPDATE = opcode{
	code:   0,
	strRep: "PRESENCE UPDATE",
}
var STATE_UPDATE = opcode{
	code:   1,
	strRep: "STATE UPDATE",
}
var CONFIG_UPDATE = opcode{
	code:   2,
	strRep: "CONFIG UPDATE",
}

type opcode struct {
	code   int
	strRep string
}
