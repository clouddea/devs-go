package main

import (
	"devs-go/examples"
	"devs-go/simulation"
)

func main() {
	rootProcStub := simulation.NewProcessorStub("localhost:8080")
	// 原子模型仿真节点
	genr := examples.NewGenerator("generator1")
	simulation.NewRoot(simulation.NewSimulator(genr, rootProcStub)).Serve("localhost:8081")
}
