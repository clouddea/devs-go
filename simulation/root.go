package simulation

import (
	"devs-go/modeling"
	"fmt"
	"time"
)

type Root struct {
	processor Processor
	t         uint64
}

func NewRoot(processor Processor) *Root {
	return &Root{
		processor: processor,
		t:         0,
	}
}

func (receiver *Root) Simulate(delay time.Duration, verbose bool) {
	receiver.processor.Init(0)
	for receiver.t < modeling.INFINITE {
		tn := receiver.processor.GetTN()
		receiver.processor.ComputeOutput(tn)
		receiver.processor.Advance(tn)
		if verbose {
			fmt.Printf("time advance: %v \n", receiver.t)
		}
		receiver.t = tn
		time.Sleep(delay)
	}
}
