package main

import (
	"devs-go/examples"
	"devs-go/simulation"
	"time"
)

func main() {
	atomic := examples.NewGenerator("generator 1")
	simulator := simulation.NewSimulator(atomic, nil)
	root := simulation.NewRoot(simulator)
	root.Simulate(time.Second, true)
}
