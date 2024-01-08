package main

import (
	"github.com/clouddea/devs-go/examples"
	"github.com/clouddea/devs-go/simulation"
	"time"
)

func main() {
	atomic := examples.NewGenerator("generator 1")
	simulator := simulation.NewSimulator(atomic, nil)
	root := simulation.NewRoot(simulator)
	root.Simulate(time.Second, true)
}
