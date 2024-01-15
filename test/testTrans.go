package main

import (
	"fmt"
	"github.com/clouddea/devs-go/examples"
	"github.com/clouddea/devs-go/modeling"
	"github.com/clouddea/devs-go/simulation"
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
	root := simulation.NewRoot(coordinator)
	root.Simulate(1*time.Second, func(t uint64) {
		fmt.Printf("time advance: %v \n", t)
	})
}
