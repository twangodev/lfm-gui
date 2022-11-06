package concurrency

type Opcode struct {
	code   int
	strRep string
}

func (o Opcode) String() string {
	return o.strRep
}
