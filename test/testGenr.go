package main

import (
	"fmt"
	"github.com/clouddea/devs-go/examples"
	"github.com/clouddea/devs-go/simulation"
	"time"
)

func main() {
	atomic := examples.NewGenerator("generator 1")
	simulator := simulation.NewSimulator(atomic, nil)
	root := simulation.NewRoot(simulator)
	root.Simulate(1*time.Second, func(t uint64) {
		fmt.Printf("time advance: %v \n", t)
	})
}
