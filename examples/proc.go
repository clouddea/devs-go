package examples

import (
	"devs-go/modeling"
	"fmt"
)

type Processor struct {
	modeling.AbstractAtomic
	job string
}

func NewProcessor(name string) *Processor {
	proc := &Processor{}
	proc.HoldIn("passivate", modeling.INFINITE)
	proc.SetName(name)
	return proc
}

func (receiver *Processor) On(e uint64, message modeling.Message) {
	if e == receiver.Sigma && message.IsEmpty() {
		// delta int
		receiver.HoldIn("passivate", modeling.INFINITE)
	} else if e == receiver.Sigma && !message.IsEmpty() {
		// delta con
		receiver.job = message.GetContents()[0].Payload.(string)
		receiver.HoldIn("active", 4) // 这一句也可以省略
	} else {
		// delta ext
		receiver.job = message.GetContents()[0].Payload.(string)
		receiver.HoldIn("active", 4)
	}
}

func (receiver *Processor) Out() (message modeling.Message) {
	fmt.Println(receiver.Name() + " out job " + fmt.Sprintf("%v", receiver.job) + ", state is " + receiver.State)
	return
}
