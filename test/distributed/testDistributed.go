package main

import (
	"github.com/clouddea/devs-go/modeling"
	"github.com/clouddea/devs-go/simulation"
	"time"
)

func main() {
	// 耦合模型节点
	genrStub := modeling.NewEntityRemote("generator1", "localhost:8081")
	transStub := modeling.NewEntityRemote("transmitter1", "localhost:8082")
	procStub := modeling.NewEntityRemote("processor1", "localhost:8083")
	coupled := &modeling.AbstractCoupled{}
	coupled.AddComponent(genrStub)
	coupled.AddComponent(transStub)
	coupled.AddComponent(procStub)
	coupled.AddCoupling(genrStub, "out", transStub, "in")
	coupled.AddCoupling(transStub, "out", procStub, "in1")
	coordinator := simulation.NewCoordinator(coupled, nil)
	root := simulation.NewRoot(coordinator)
	go root.Serve("localhost:8080")

	root.Simulate(1*time.Second, true)
}
