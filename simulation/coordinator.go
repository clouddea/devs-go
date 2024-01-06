package simulation

import (
	"devs-go/modeling"
	"sync"
)

type Coordinator struct {
	devs       modeling.Coupled
	processors map[modeling.Entity]Processor
	tl         uint64
	tn         uint64
	parent     Processor
}

func NewCoordinator(devs modeling.Coupled, parent Processor) *Coordinator {
	cor := &Coordinator{
		devs:       devs,
		processors: make(map[modeling.Entity]Processor),
		tl:         0,
		tn:         0,
		parent:     parent,
	}
	// 为子组件生成对应的处理器
	for _, component := range devs.GetComponents() {
		switch component.(type) {
		case modeling.Atomic:
			cor.processors[component] = NewSimulator(component.(modeling.Atomic), cor)
		case modeling.Coupled:
			cor.processors[component] = NewCoordinator(component.(modeling.Coupled), cor)
		}
	}
	return cor
}

func (receiver *Coordinator) getTN() uint64 {
	tn := modeling.INFINITE
	for _, v := range receiver.processors {
		tn1 := v.GetTN()
		if tn1 < tn {
			tn = tn1
		}
	}
	return tn
}

func (receiver *Coordinator) sentAllMessage(callable func(item Processor)) {
	var semaphore sync.WaitGroup
	semaphore.Add(len(receiver.processors))
	for _, proc := range receiver.processors {
		go func(p Processor) {
			callable(p)
			semaphore.Done()
		}(proc)
	}
	semaphore.Wait()
}

func (receiver *Coordinator) GetTN() uint64 {
	return receiver.tn
}

func (receiver *Coordinator) Init(t uint64) {
	receiver.sentAllMessage(func(item Processor) {
		item.Init(t)
	})
	receiver.tl = t
	receiver.tn = receiver.getTN()
}

func (receiver *Coordinator) Advance(t uint64) {
	// 不需要区分IMM和INF，直接通知所有子组件发生时间推进
	receiver.sentAllMessage(func(item Processor) {
		item.Advance(t)
	})
	receiver.tl = t
	receiver.tn = receiver.getTN()
}

func (receiver *Coordinator) ComputeOutput(t uint64) {
	if t == receiver.tn {
		// sent to child
		receiver.sentAllMessage(func(item Processor) {
			item.ComputeOutput(t)
		})
	}
}

func (receiver *Coordinator) PutMessage(message modeling.Message, t uint64) {
	if message.IsEmpty() {
		return
	}
	contents := message.GetContents()
	var source string = contents[0].Source
	var sourcePort string = contents[0].SourcePort
	// 假定消息源是当前耦合模型，如果找到某个子组件名字与source一样，则说明消息源是子组件
	var entity modeling.Entity = receiver.devs
	if component, ok := receiver.devs.GetComponentMap()[source]; ok {
		entity = component
	}

	pairs := receiver.devs.GetCoupling(entity, sourcePort)
	for _, pair := range pairs {
		var processor Processor = nil
		if pair.Component.Name() == receiver.devs.Name() {
			processor = receiver.parent
		} else {
			processor = receiver.processors[pair.Component]
		}
		// 发送消息
		if processor != nil {
			// 改变消息源
			var newMessage modeling.Message
			for _, content := range message.GetContents() {
				newMessage.AddContent(modeling.Content{
					Source:     pair.Component.Name(),
					SourcePort: pair.Port,
					Payload:    content.Payload,
				})
			}
			processor.PutMessage(newMessage, t)
		}
	}
}
