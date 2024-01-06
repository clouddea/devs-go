package main

import (
	"devs-go/examples"
	"devs-go/simulation"
)

func main() {
	rootProcStub := simulation.NewProcessorStub("localhost:8080")
	// 原子模型仿真节点
	trans := examples.NewTransmitter("transmitter1")
	simulation.NewRoot(simulation.NewSimulator(trans, rootProcStub)).Serve("localhost:8082")
}
