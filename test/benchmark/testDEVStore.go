package main

import (
	"fmt"
	"github.com/clouddea/devs-go/modeling"
	"github.com/clouddea/devs-go/simulation"
)

func testLI() {
	model := NewDEVStoneModel("LI", 200, 40)
	genr := NewDEVStoneGenr("genr")
	env := &modeling.AbstractCoupled{}
	env.AddComponent(genr)
	env.AddComponent(model)
	env.AddCoupling(genr, "out", model, "in")

	cord := simulation.NewCoordinator(env, nil)
	root := simulation.NewRoot(cord)
	root.Simulate(0, nil)
	fmt.Printf("总事件数: %v , 应有事件数：7762 \n", N_EVENTS)
}

func testHI() {
	model := NewDEVStoneModel("HI", 200, 40)
	genr := NewDEVStoneGenr("genr")
	env := &modeling.AbstractCoupled{}
	env.AddComponent(genr)
	env.AddComponent(model)
	env.AddCoupling(genr, "out", model, "in")

	cord := simulation.NewCoordinator(env, nil)
	root := simulation.NewRoot(cord)
	root.Simulate(0, nil)
	fmt.Printf("总事件数: %v , 应有事件数：155221 \n", N_EVENTS)
}

func testHO() {
	model := NewDEVStoneModel("HO", 200, 40)
	genr1 := NewDEVStoneGenr("genr1")
	genr2 := NewDEVStoneGenr("genr2")
	env := &modeling.AbstractCoupled{}
	env.AddComponent(genr1)
	env.AddComponent(genr2)
	env.AddComponent(model)
	env.AddCoupling(genr1, "out", model, "in1")
	env.AddCoupling(genr2, "out", model, "in2")
	cord := simulation.NewCoordinator(env, nil)
	root := simulation.NewRoot(cord)
	root.Simulate(0, nil)
	fmt.Printf("总事件数: %v , 应有事件数：155221 \n", N_EVENTS)
}

func testHOmod() {
	model := NewDEVStoneModel("HOmod", 20, 40)
	genr1 := NewDEVStoneGenr("genr1")
	genr2 := NewDEVStoneGenr("genr2")
	env := &modeling.AbstractCoupled{}
	env.AddComponent(genr1)
	env.AddComponent(genr2)
	env.AddComponent(model)
	env.AddCoupling(genr1, "out", model, "in1")
	env.AddCoupling(genr2, "out", model, "in2")

	cord := simulation.NewCoordinator(env, nil)
	root := simulation.NewRoot(cord)
	root.Simulate(0, nil)
	fmt.Printf("总事件数: %v , 应有事件数：5506372 \n", N_EVENTS)
}

func main() {
	//testLI()
	//testHI()
	//testHO()
	//testHOmod()
}
