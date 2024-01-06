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
	coupled := &modeling.AbstractCoupled{}
	coupled.AddComponent(genr)
	coupled.AddComponent(trans)
	coupled.AddCoupling(genr, "out", trans, "in")
	coordinator := simulation.NewCoordinator(coupled, nil)
	coordinator.Simulate(1*time.Second, true)
}
