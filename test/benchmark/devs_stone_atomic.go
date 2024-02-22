package main

import (
	"github.com/clouddea/devs-go/modeling"
	"sync"
)

type DEVStoneAtomic struct {
	modeling.AbstractAtomic
	list modeling.Message
}

func NewDEVStoneAtomic(name string) *DEVStoneAtomic {
	atomic := &DEVStoneAtomic{}
	atomic.SetName(name)
	atomic.HoldIn("passive", modeling.INFINITE)
	return atomic
}

func (receiver *DEVStoneAtomic) On(e uint64, message modeling.Message) {
	if e == receiver.Ta() && message.IsEmpty() {
		// delta int
		receiver.HoldIn("passive", modeling.INFINITE)
	} else if e == receiver.Ta() && !message.IsEmpty() {
		// delta con
		receiver.list.Add(message)
		receiver.HoldIn("active", 0)
	} else {
		// delta ext
		receiver.list.Add(message)
		receiver.HoldIn("active", 0)
	}
}

var N_EVENTS int = 0
var mutex sync.Mutex

func (receiver *DEVStoneAtomic) Out() modeling.Message {
	msg := *modeling.NewMessage()
	//for i := 0; i < len(receiver.list.Contents); i++ {
	//	msg.AddContent(*receiver.MakeContent("out", "data"))
	//}
	// 为了节省内存，只变成1个
	msg.AddContent(*receiver.MakeContent("out", "data"))
	//fmt.Printf("%v out count: %v \n", receiver.Name(), len(receiver.list.Contents))
	//mutex.Lock()
	//N_EVENTS += 1
	//mutex.Unlock()

	receiver.list.Clear()
	return msg
}

type DEVStoneGenr struct {
	modeling.AbstractAtomic
}

func NewDEVStoneGenr(name string) *DEVStoneGenr {
	atomic := &DEVStoneGenr{}
	atomic.SetName(name)
	atomic.HoldIn("active", 0)
	return atomic
}

func (receiver *DEVStoneGenr) On(e uint64, message modeling.Message) {
	receiver.HoldIn("passive", modeling.INFINITE)
}

func (receiver *DEVStoneGenr) Out() modeling.Message {
	msg := *modeling.NewMessage()
	msg.AddContent(*receiver.MakeContent("out", "data"))
	return msg
}
