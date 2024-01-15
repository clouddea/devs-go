package examples

import (
	"github.com/clouddea/devs-go/modeling"
	"strconv"
)

type CA struct {
	neighbors []bool
	modeling.AbstractAtomic
}

func NewCA(name string, alive bool, neighbors []bool) *CA {
	// 初始化时需要自已的状态和邻居的状态
	//  0 1 2
	//  3 x 4
	//  5 6 7
	genr := &CA{
		neighbors: neighbors,
	}
	genr.SetName(name)
	if alive {
		genr.HoldIn("alive", 1)
	} else {
		genr.HoldIn("dead", 1)
	}
	return genr
}

func (receiver *CA) On(e uint64, message modeling.Message) {
	if e == receiver.Ta() && message.IsEmpty() {
		// delta int
		// 一般情况下邻居没有变化，因此自已也无需变化
		// 但是在初始时例外
	} else if e == receiver.Ta() && !message.IsEmpty() {
		// delta con
		// 邻居发生变化，更新一下缓存的状态

	} else {
		// delta ext
		// 不可能有单独的外部转移
	}
	receiver.HoldIn(receiver.nextState(), 1)
	for _, content := range message.GetContents() {
		index, _ := strconv.Atoi(content.SourcePort)
		receiver.neighbors[index] = !receiver.neighbors[index]
	}
}

func (receiver *CA) nextState() string {
	liveCount := 0
	for _, alived := range receiver.neighbors {
		if alived {
			liveCount += 1
		}
	}

	if receiver.State == "alive" {
		if liveCount > 3 || liveCount < 2 {
			return "dead"
		} else {
			return "alive"
		}
	} else {
		if liveCount == 3 {
			return "alive"
		} else {
			return "dead"
		}
	}
}

func (receiver *CA) Out() modeling.Message {
	msg := *modeling.NewMessage()
	if receiver.nextState() != receiver.State {
		for i := 0; i < 8; i++ {
			msg.AddContent(*receiver.MakeContent(strconv.Itoa(i), nil))
		}
	}
	return msg
}
