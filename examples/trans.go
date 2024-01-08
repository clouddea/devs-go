package examples

import (
	"fmt"
	"github.com/clouddea/devs-go/modeling"
)

type Transmitter struct {
	modeling.AbstractAtomic
	buf modeling.Message
}

func NewTransmitter(name string) *Transmitter {
	trans := &Transmitter{}
	trans.SetName(name)
	trans.HoldIn("passivate", modeling.INFINITE)
	return trans
}

func (receiver *Transmitter) On(e uint64, message modeling.Message) {
	receiver.buf.Add(message)
	if e == receiver.Sigma && message.IsEmpty() {
		// delta int
		receiver.HoldIn("passivate", modeling.INFINITE)
	} else if e == receiver.Sigma && !message.IsEmpty() {
		// delta con
		receiver.HoldIn("active", 0) // 这一句也可以省略
	} else {
		// delta ext
		receiver.HoldIn("active", 0)
	}
}

func (receiver *Transmitter) Out() (message modeling.Message) {
	for _, cont := range receiver.buf.GetContents() {
		message.AddContent(*receiver.MakeContent("out", cont.Payload))
	}
	receiver.buf.Clear()
	fmt.Println(receiver.Name() + " out" + fmt.Sprintf("%v", message) + ", state is " + receiver.State)
	return
}
