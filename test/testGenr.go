package main

import (
	"devs-go/examples"
	"devs-go/simulation"
	"time"
)

func main() {
	atomic := examples.NewGenerator("generator 1")
	simulator := simulation.NewSimulator(atomic, nil)
	simulator.Simulate(time.Second, true)
}
