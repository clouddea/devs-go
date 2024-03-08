package examples

import (
	"fmt"
	"github.com/clouddea/devs-go/modeling"
	"math/rand"
)

type DAGNode struct {
	Deps int // 依赖的个数
	Val  int // 初始值
	Sum  int // 最终值
	Exp  int // 期望值
	modeling.AbstractAtomic
}

func NewDAGNode(name string, deps int, val int, exp int) *DAGNode {
	node := &DAGNode{}
	node.SetName(name)
	node.Deps = deps
	node.Val = val
	node.Sum = val
	node.Exp = exp // 期望值
	if deps == 0 {
		node.HoldIn("active", 0) // 没有依赖。则处于活跃态
	} else {
		node.HoldIn("passive", modeling.INFINITE) // 有依赖，处于等待态
	}
	return node
}

func (receiver *DAGNode) On(e uint64, message modeling.Message) {

	if e == receiver.Ta() && message.IsEmpty() {
		// delta int， 只有活跃态会发生内部转移
		receiver.HoldIn("finish", modeling.INFINITE)
	} else if e == receiver.Ta() && !message.IsEmpty() {
		// delta con
		// 不可能发生
	} else {
		// delta ext
		for _, msg := range message.Contents {
			receiver.Deps--
			receiver.Sum += msg.Payload.(int)
		}
		if receiver.Deps == 0 {
			receiver.HoldIn("active", 0)
		}
	}
}

func (receiver *DAGNode) Out() modeling.Message {
	msg := *modeling.NewMessage()
	msg.AddContent(*receiver.MakeContent("out", receiver.Sum))
	return msg
}

type DAG struct {
	Width    int
	Depth    int
	DAGNodes [][]*DAGNode
	modeling.AbstractCoupled
}

func NewDAG(name string, width int, depth int) *DAG {
	dag := &DAG{}
	dag.SetName(name)
	dag.Width = width
	dag.Depth = depth
	dag.DAGNodes = make([][]*DAGNode, depth)
	for i := 0; i < depth; i++ {
		dag.DAGNodes[i] = make([]*DAGNode, width)
	}
	for i := 0; i < depth; i++ {
		for j := 0; j < width; j++ {
			val := rand.Intn(1000)
			if i == 0 {
				dag.DAGNodes[i][j] = NewDAGNode(fmt.Sprintf("node_%v_%v", i, j), 0, val, val)
				dag.AddComponent(dag.DAGNodes[i][j])
			} else {
				deps := rand.Intn(width) + 1
				rand.Shuffle(width, func(a, b int) {
					dag.DAGNodes[i-1][a], dag.DAGNodes[i-1][b] = dag.DAGNodes[i-1][b], dag.DAGNodes[i-1][a]
				})
				exp := val
				for k := 0; k < deps; k++ {
					exp += dag.DAGNodes[i-1][k].Exp
				}
				dag.DAGNodes[i][j] = NewDAGNode(fmt.Sprintf("node_%v_%v", i, j), deps, val, exp)
				dag.AddComponent(dag.DAGNodes[i][j])
				for k := 0; k < deps; k++ {
					dag.AddCoupling(dag.DAGNodes[i-1][k], "out", dag.DAGNodes[i][j], "in")
				}
			}
		}
	}
	return dag
}
