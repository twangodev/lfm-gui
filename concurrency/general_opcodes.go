package concurrency

var PRESENCE_UPDATE = Opcode{
	code:   0,
	strRep: "PRESENCE_UPDATE",
}
var STATE_UPDATE = Opcode{
	code:   1,
	strRep: "STATE_UPDATE",
}
var CONFIG_UPDATE = Opcode{
	code:   2,
	strRep: "CONFIG_UPDATE",
}
