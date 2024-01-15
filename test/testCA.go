package main

import (
	"fmt"
	"github.com/clouddea/devs-go/examples"
	"github.com/clouddea/devs-go/modeling"
	"github.com/clouddea/devs-go/simulation"
	webview "github.com/webview/webview_go"
)

const nrows = 20
const ncols = 20

var states [][]bool
var atomics [][]modeling.Atomic

func genNeighborsGetNeiState(i int, j int) bool {
	if i < 0 || i >= nrows || j < 0 || j >= ncols {
		return false
	}
	return states[i][j]
}

func main() {
	states = make([][]bool, nrows)
	atomics = make([][]modeling.Atomic, nrows)
	for i := 0; i < nrows; i++ {
		states[i] = make([]bool, ncols)
		atomics[i] = make([]modeling.Atomic, ncols)
	}

	initStats := []string{
		"    .      .        ",
		"  . .      .        ",
		"   ..      .        ",
		"                    ",
		"                ... ",
		"               ...  ",
		"  ..                ",
		"  ..                ",
		"..                  ",
		"..                  ",
		"       __           ",
		"      _ _           ",
		"                    ",
		"    _ _             ",
		"    __              ",
		"                    ",
		" .                  ",
		"  ..                ",
		"..                  ",
		"  .                 ",
	}

	for i := 0; i < nrows; i++ {
		bytes := []byte(initStats[i])
		for j := 0; j < ncols; j++ {
			if bytes[j] == 46 {
				states[i][j] = true
			}
		}
	}

	coupled := &modeling.AbstractCoupled{}

	for i := 0; i < nrows; i++ {
		for j := 0; j < ncols; j++ {
			atomics[i][j] = examples.NewCA(fmt.Sprintf("%v_%v", i, j),
				states[i][j],
				[]bool{
					genNeighborsGetNeiState(i-1, j-1),
					genNeighborsGetNeiState(i-1, j),
					genNeighborsGetNeiState(i-1, j+1),
					genNeighborsGetNeiState(i, j-1),
					genNeighborsGetNeiState(i, j+1),
					genNeighborsGetNeiState(i+1, j-1),
					genNeighborsGetNeiState(i+1, j),
					genNeighborsGetNeiState(i+1, j+1),
				},
			)
			coupled.AddComponent(atomics[i][j])
			if i > 0 {
				coupled.AddCoupling(atomics[i-1][j], "6", atomics[i][j], "1")
				coupled.AddCoupling(atomics[i][j], "1", atomics[i-1][j], "6")
			}
			if j > 0 {
				coupled.AddCoupling(atomics[i][j-1], "4", atomics[i][j], "3")
				coupled.AddCoupling(atomics[i][j], "3", atomics[i][j-1], "4")
			}
			if i > 0 && j > 0 {
				coupled.AddCoupling(atomics[i-1][j-1], "7", atomics[i][j], "0")
				coupled.AddCoupling(atomics[i][j], "0", atomics[i-1][j-1], "7")
			}
			if i > 0 && j < ncols-1 {
				coupled.AddCoupling(atomics[i-1][j+1], "5", atomics[i][j], "2")
				coupled.AddCoupling(atomics[i][j], "2", atomics[i-1][j+1], "5")
			}
		}
	}

	coordinator := simulation.NewCoordinator(coupled, nil)
	root := simulation.NewRoot(coordinator)

	// // 控制台直接打印状态
	//root.Simulate(1*time.Second, func(t uint64) {
	//	fmt.Printf("time advance: %v \n", t)
	//	for i := 0; i < nrows; i++ {
	//		for j := 0; j < ncols; j++ {
	//			if atomics[i][j].(*examples.CA).State == "alive" {
	//				fmt.Print(".")
	//			} else {
	//				fmt.Print(" ")
	//			}
	//		}
	//		fmt.Println()
	//	}
	//})

	// 启动窗口
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Bind Example")
	w.SetSize(400, 400, webview.HintFixed)
	w.Bind("setup", func() {
		root.Setup()
	})
	w.Bind("step", func() [][]bool {
		root.Step()
		for i := 0; i < nrows; i++ {
			for j := 0; j < ncols; j++ {
				if atomics[i][j].(*examples.CA).State == "alive" {
					states[i][j] = true
				} else {
					states[i][j] = false
				}
			}
		}
		return states
	})

	w.SetHtml(html)
	w.Run()
}

const html = `
<!DOCTYPE html>
<html>
<head>
<style>
  html, body {
	margin: 0;
    padding: 0;
  }
  .container {
    display: flex;
    flex-wrap: wrap;
    box-sizing: border-box;
    /*border: 1px solid red;*/
    width: 400px;
    height: 400px
  }

  .container .box {
	 width: 20px;
	 height: 20px;
    box-sizing: border-box;
    border: 1px solid white;
    background-color: #FF000023
  }
</style>
</head>
<body>
<div id="container" class="container"></div>
<script>
  var container = document.getElementById("container");
  var boxes = [];
  for(let i = 0; i < 20; i++){
	  boxes[i] = []
      for(let j = 0; j < 20; j++){
		 var element = document.createElement('div');
         element.classList.add('box')
		 container.appendChild(element);
         boxes[i].push(element)
	  }
  }

  window.setup().then(() => {
     container.style.background = "#00FF0050"
  })
  window.setInterval(() => {
      window.step().then((status) => {
             //alert(status[0][0])
			for(let i = 0; i < 20; i++){
				  for(let j = 0; j < 20; j++){
                     if(status[i][j]){
                        boxes[i][j].style.background = "black"
                     } else {
                        boxes[i][j].style.background = "#FF000023"
                     }
				  }
            }
      })
  }, 1000)
</script>
</body>
</html>
`

// 下面也OK
// PS：使用 github.com/polevpn/webview， 或者使用github.com/webview/webview_go [webview官方binding] 要在windows上安装 mingw64-win32-seh-msvcrt版本
//package main
//
//import "github.com/polevpn/webview"
//
//func main() {
//	w := webview.New(500, 600, false, false)
//	defer w.Destroy()
//	w.SetTitle("Minimal webview example")
//	w.SetSize(800, 600, webview.HintNone)
//	w.Navigate("https://en.m.wikipedia.org/wiki/Main_Page")
//	w.Run()
//}
