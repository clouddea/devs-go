package main

import (
	"fmt"
	"github.com/clouddea/devs-go/modeling"
	"github.com/clouddea/devs-go/simulation"
)

func testLI(w, d int) {
	model := NewDEVStoneModel("LI", d, w)
	genr := NewDEVStoneGenr("genr")
	env := &modeling.AbstractCoupled{}
	env.AddComponent(genr)
	env.AddComponent(model)
	env.AddCoupling(genr, "out", model, "in")

	cord := simulation.NewCoordinator(env, nil)
	root := simulation.NewRoot(cord)
	root.Simulate(0, nil)

	shouldEvents := (w-1)*(d-1) + 1
	fmt.Printf("总事件数: %v , 应有事件数：%v \n", N_EVENTS, shouldEvents)
}

func testHI(w, d int) {
	model := NewDEVStoneModel("HI", d, w)
	genr := NewDEVStoneGenr("genr")
	env := &modeling.AbstractCoupled{}
	env.AddComponent(genr)
	env.AddComponent(model)
	env.AddCoupling(genr, "out", model, "in")

	cord := simulation.NewCoordinator(env, nil)
	root := simulation.NewRoot(cord)
	root.Simulate(0, nil)
	shouldEvents := 1 + (d-1)*(w-1)*w/2
	fmt.Printf("总事件数: %v , 应有事件数：%v \n", N_EVENTS, shouldEvents)
}

func testHO(w, d int) {
	model := NewDEVStoneModel("HO", d, w)
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
	shouldEvents := 1 + (d-1)*(w-1)*w/2
	fmt.Printf("总事件数: %v , 应有事件数：%v \n", N_EVENTS, shouldEvents)
}

func testHOmod(w, d int) {
	model := NewDEVStoneModel("HOmod", d, w)
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

	shouldEvents := 1
	for i := 1; i < d; i++ {
		shouldEvents += (1+(i-1)*(w-1))*(w-1)*w/2 + (w-1)*(w+(i-1)*(w-1))
	}
	fmt.Printf("总事件数: %v , 应有事件数：%v \n", N_EVENTS, shouldEvents)
}

func main() {
	//testLI(200, 200)
	//N_EVENTS = 0
	//testLI(200, 40)
	//N_EVENTS = 0
	//testLI(40, 200)
	//N_EVENTS = 0
	//testHI(200, 200)
	//N_EVENTS = 0
	//testHI(200, 40)
	//N_EVENTS = 0
	//testHI(40, 200)
	//N_EVENTS = 0
	//testHO(200, 200)
	//N_EVENTS = 0
	//testHO(200, 40)
	//N_EVENTS = 0
	//testHO(40, 200)
	//N_EVENTS = 0
	testHOmod(20, 20)
	N_EVENTS = 0
	testHOmod(20, 4)
	N_EVENTS = 0
	testHOmod(4, 20)
	N_EVENTS = 0
	testHOmod(4, 4)
	N_EVENTS = 0
}
