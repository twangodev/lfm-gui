package concurrency

var REQUEST_PRESENCE_UPDATE = Opcode{
	code:   0,
	strRep: "REQUEST_PRESENCE_UPDATE",
}
var STATE_UPDATE = Opcode{
	code:   1,
	strRep: "STATE_UPDATE",
}
var CONFIG_UPDATE = Opcode{
	code:   2,
	strRep: "CONFIG_UPDATE",
}

var ELAPSED_UPDATE = Opcode{
	code:   3,
	strRep: "ELAPSED_UPDATE",
}
