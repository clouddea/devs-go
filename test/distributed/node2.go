package main

import (
	"github.com/clouddea/devs-go/examples"
	"github.com/clouddea/devs-go/simulation"
)

func main() {
	rootProcStub := simulation.NewProcessorStub("localhost:8080")
	// 原子模型仿真节点
	trans := examples.NewTransmitter("transmitter1")
	simulation.NewRoot(simulation.NewSimulator(trans, rootProcStub)).Serve("localhost:8082")
}
