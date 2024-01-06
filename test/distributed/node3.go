package main

import (
	"devs-go/examples"
	"devs-go/simulation"
)

func main() {
	rootProcStub := simulation.NewProcessorStub("localhost:8080")
	// 原子模型仿真节点
	proc := examples.NewProcessor("processor1")
	simulation.NewRoot(simulation.NewSimulator(proc, rootProcStub)).Serve("localhost:8083")
}
