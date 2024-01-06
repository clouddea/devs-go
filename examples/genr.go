package examples

import (
	"devs-go/modeling"
	"fmt"
	"strconv"
)

type Generator struct {
	count int
	modeling.AbstractAtomic
}

func NewGenerator(name string) *Generator {
	genr := &Generator{}
	genr.SetName(name)
	genr.HoldIn("active", 10)
	return genr
}

func (receiver *Generator) On(e uint64, message modeling.Message) {
	if e == receiver.Ta() && message.IsEmpty() {
		// delta int
		receiver.count += 1
		fmt.Println(receiver.Name() + ": delta int")
		if receiver.State == "active" {
			receiver.HoldIn("hot active", 5)
		} else {
			receiver.HoldIn("active", 10)
		}
	} else if e == receiver.Ta() && !message.IsEmpty() {
		// delta con
		fmt.Println(receiver.Name() + ": delta con")
	} else {
		// delta ext
		fmt.Println(receiver.Name() + ": delta ext")
	}
}

func (receiver Generator) Out() modeling.Message {
	msg := *modeling.NewMessage()
	msg.AddContent(modeling.Content{})
	fmt.Println(receiver.Name() + " out count" + strconv.Itoa(receiver.count) + ", state is " + receiver.State)
	return msg
}
