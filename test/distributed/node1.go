package main

import (
	"github.com/clouddea/devs-go/examples"
	"github.com/clouddea/devs-go/simulation"
)

func main() {
	rootProcStub := simulation.NewProcessorStub("localhost:8080")
	// 原子模型仿真节点
	genr := examples.NewGenerator("generator1")
	simulation.NewRoot(simulation.NewSimulator(genr, rootProcStub)).Serve("localhost:8081")
}
