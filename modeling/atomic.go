package modeling

import (
	"fmt"
)

const (
	INFINITE uint64 = 0x7FFFFFFFFFFFFFFF
)

type Atomic interface {
	Entity
	Init()
	On(e uint64, message Message)
	Out() Message
	Ta() uint64
}

type AbstractAtomic struct {
	Sigma uint64 // 当前状态保持的时间
	State string // 当前状态
	name  string // 当前模型的名字
}

func (receiver *AbstractAtomic) Init() {
	if receiver.State == "" {
		receiver.HoldIn("passive", INFINITE)
	}
}

func (receiver AbstractAtomic) Name() string {
	return receiver.name
}

func (receiver *AbstractAtomic) SetName(name string) {
	receiver.name = name
}

func (receiver AbstractAtomic) Ta() uint64 {
	return receiver.Sigma
}

func (receiver *AbstractAtomic) On(e uint64, message Message) {
	if e == receiver.Sigma && message.IsEmpty() {
		// delta int
		fmt.Println(receiver.name + ": delta int")
	} else if e == receiver.Sigma && !message.IsEmpty() {
		// delta con
		fmt.Println(receiver.name + ": delta con")
	} else {
		// delta ext
		fmt.Println(receiver.name + ": delta ext")
	}
}

func (receiver *AbstractAtomic) Out() (message Message) {
	return *NewMessage()
}

func (receiver *AbstractAtomic) HoldIn(state string, t uint64) {
	receiver.State = state
	receiver.Sigma = t
}

func (receiver AbstractAtomic) MakeContent(outPort string, data interface{}) *Content {
	return &Content{
		Source:     receiver.Name(),
		SourcePort: outPort,
		Target:     "##NoTarget",
		TargetPort: "##NoTargetPort",
		Payload:    data,
	}
}
