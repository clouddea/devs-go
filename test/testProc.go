package main

import (
	"devs-go/examples"
	"devs-go/modeling"
	"devs-go/simulation"
	"time"
)

func main() {

	genr := examples.NewGenerator("generator 1")
	trans := examples.NewTransmitter("transmitter 1")
	proc := examples.NewProcessor("processor 1")
	coupled := &modeling.AbstractCoupled{}
	coupled.AddComponent(genr)
	coupled.AddComponent(trans)
	coupled.AddComponent(proc)
	coupled.AddCoupling(genr, "out", trans, "in")
	coupled.AddCoupling(trans, "out", proc, "in1")
	coordinator := simulation.NewCoordinator(coupled, nil)
	root := simulation.NewRoot(coordinator)
	root.Simulate(1*time.Second, true)
}
