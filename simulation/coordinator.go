package simulation

import "devs-go/modeling"

type Coordinator struct {
	devs       modeling.Coupled
	processors map[modeling.Atomic]Processor
	tl         uint64
	tn         uint64
	parent     Processor
}

func NewCoordinator(devs modeling.Coupled, parent Processor) *Coordinator {
	return &Coordinator{
		devs:       devs,
		processors: make(map[modeling.Atomic]Processor),
		tl:         0,
		tn:         0,
		parent:     parent,
	}
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

func (receiver *Coordinator) GetTN() uint64 {
	return receiver.tn
}

func (receiver *Coordinator) Init(t uint64) {
	for _, v := range receiver.processors {
		v.Init(t)
	}
	receiver.tl = t
	receiver.tn = receiver.getTN()
}

func (receiver *Coordinator) Advance(count int, t uint64) {
	// IMM: 子组件中即将发生内部转移的
	// INF： Ii, 其中i 属于 IMM + {SELF}
	// 计算IMM+INF-SELF
	children := make(map[Processor]bool)
	for k, v := range receiver.processors {
		if t == v.GetTN() {
			children[v] = true
			influencors := receiver.devs.GetCoupling(k)
			for _, inf := range influencors {
				infProc := receiver.processors[inf]
				children[infProc] = true // TODO: exclude self
			}
		}
	}
	for _, child := range children {
		// 计算每个子组件对应的 influencer个数, 即影响它的组件

	}

}

//func (receiver *Coordinator) putMessage(message modeling.Message, t uint64) {
//	receiver.input.Add(message)
//	receiver.semaphore.Done()
//}
