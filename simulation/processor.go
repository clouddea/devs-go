package simulation

import "devs-go/modeling"

type Processor interface {
	Init(t uint64)
	Advance(count int, t uint64)
	ComputeOutput(t uint64)
	PutMessage(message modeling.Message, t uint64)
	GetTN() uint64
}
