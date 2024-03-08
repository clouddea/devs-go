package main

import (
	"fmt"
	"github.com/clouddea/devs-go/examples"
	"github.com/clouddea/devs-go/simulation"
	"os"
)

func main() {
	coupled := examples.NewDAG("dag", 5, 8)
	processor := simulation.NewCoordinator(coupled, nil)
	root := simulation.NewRoot(processor)
	root.Simulate(0, nil)
	for i := 0; i < coupled.Depth; i++ {
		for j := 0; j < coupled.Width; j++ {
			if coupled.DAGNodes[i][j].Sum != coupled.DAGNodes[i][j].Exp {
				fmt.Println("bad verification!")
				os.Exit(-1)
			}
		}
	}
	// 生成绘图指令
	for i := 0; i < coupled.Depth; i++ {
		for j := 0; j < coupled.Width; j++ {
			fmt.Printf("%v[label=\"%v\\nval=%v\\nsum=%v\\nexp=%v\"]\n",
				coupled.DAGNodes[i][j].Name(),
				coupled.DAGNodes[i][j].Name(),
				coupled.DAGNodes[i][j].Val,
				coupled.DAGNodes[i][j].Sum,
				coupled.DAGNodes[i][j].Exp)
		}
	}

	for i := 0; i < coupled.Depth-1; i++ {
		for j := 0; j < coupled.Width; j++ {
			outs := coupled.GetCoupling(coupled.DAGNodes[i][j], "out")
			for k := 0; k < len(outs); k++ {
				fmt.Printf("%v -> %v \n", coupled.DAGNodes[i][j].Name(), outs[k].Component.Name())
			}
		}
	}

}
