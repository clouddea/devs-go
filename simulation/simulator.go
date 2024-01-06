package simulation

import (
	"devs-go/modeling"
	"fmt"
	"time"
)

/*
*
模型是抽象的，但是仿真器和协调器是具体，这就能实现，只实现一套调度代码
*/
type Simulator struct {
	atomic modeling.Atomic
	input  modeling.Message
	tl     uint64
	tn     uint64
	parent Processor
}

func NewSimulator(atomic modeling.Atomic, parent Processor) *Simulator {
	return &Simulator{
		atomic: atomic,
		input:  *modeling.NewMessage(),
		tl:     0,
		tn:     0,
		parent: parent,
	}
}

func (receiver *Simulator) Init(t uint64) {
	receiver.atomic.Init()
	receiver.tl = t
	receiver.tn = t + receiver.atomic.Ta()
}

func (receiver *Simulator) Advance(t uint64) {
	if t != receiver.tn && receiver.input.IsEmpty() {
		return
	}
	receiver.atomic.On(t-receiver.tl, receiver.input)
	receiver.tl = t
	receiver.tn = t + receiver.atomic.Ta()
	receiver.input.Clear()
}

func (receiver *Simulator) ComputeOutput(t uint64) {
	if t == receiver.tn {
		// sent to parent
		msg := receiver.atomic.Out()
		if receiver.parent != nil {
			receiver.parent.PutMessage(msg, t)
		}
	}
}

func (receiver *Simulator) PutMessage(message modeling.Message, t uint64) {
	receiver.input.Add(message)
}

func (receiver *Simulator) GetTN() uint64 {
	return receiver.tn
}

func (receiver *Simulator) Simulate(delay time.Duration, verbose bool) {
	receiver.Init(0)
	for receiver.tl < modeling.INFINITE {
		receiver.ComputeOutput(receiver.tn)
		receiver.Advance(receiver.tn)
		if verbose {
			fmt.Printf("time advance: %v \n", receiver.tl)
		}
		time.Sleep(delay)
	}
}
