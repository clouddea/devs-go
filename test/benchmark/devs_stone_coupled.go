package main

import (
	"fmt"
	"github.com/clouddea/devs-go/modeling"
	"strconv"
)

func NewDEVStoneLICoupled(name string, entity modeling.Entity, atomics ...modeling.Atomic) modeling.Coupled {
	comp := &modeling.AbstractCoupled{}
	comp.SetName(name)
	// 添加组件和耦合关系
	comp.AddComponent(entity)
	comp.AddCoupling(comp, "in", entity, "in")
	comp.AddCoupling(entity, "out", comp, "out")
	for i := 0; i < len(atomics); i++ {
		comp.AddComponent(atomics[i])
		comp.AddCoupling(comp, "in", atomics[i], "in")
	}
	return comp
}

func NewDEVStoneHICoupled(name string, entity modeling.Entity, atomics ...modeling.Atomic) modeling.Coupled {
	comp := &modeling.AbstractCoupled{}
	comp.SetName(name)
	// 添加组件和耦合关系
	comp.AddComponent(entity)
	comp.AddCoupling(comp, "in", entity, "in")
	comp.AddCoupling(entity, "out", comp, "out")
	for i := 0; i < len(atomics); i++ {
		comp.AddComponent(atomics[i])
		comp.AddCoupling(comp, "in", atomics[i], "in")
	}
	// 额外增加的耦合
	for i := 0; i < len(atomics)-1; i++ {
		comp.AddCoupling(atomics[i], "out", atomics[i+1], "in")
	}
	return comp
}

func NewDEVStoneHOCoupled(name string, entity modeling.Entity, atomics ...modeling.Atomic) modeling.Coupled {
	comp := &modeling.AbstractCoupled{}
	comp.SetName(name)
	// 添加组件和耦合关系
	switch entity.(type) {
	case modeling.Atomic:
		comp.AddComponent(entity)
		comp.AddCoupling(comp, "in1", entity, "in")
		comp.AddCoupling(entity, "out", comp, "out1")
	case modeling.Coupled:
		comp.AddComponent(entity)
		comp.AddCoupling(comp, "in1", entity, "in1")
		comp.AddCoupling(entity, "out1", comp, "out1")
		comp.AddCoupling(comp, "in2", entity, "in2")
		for i := 0; i < len(atomics); i++ {
			comp.AddComponent(atomics[i])
			comp.AddCoupling(comp, "in2", atomics[i], "in")
			comp.AddCoupling(atomics[i], "out", comp, "out2")
		}
		// 额外增加的耦合
		for i := 0; i < len(atomics)-1; i++ {
			comp.AddCoupling(atomics[i], "out", atomics[i+1], "in")
		}
	}
	return comp
}

func NewDEVStoneHOmodCoupled(name string, entity modeling.Entity, w int, l int) modeling.Coupled {
	comp := &modeling.AbstractCoupled{}
	comp.SetName(name)
	// 添加组件和耦合关系
	switch entity.(type) {
	case modeling.Atomic:
		comp.AddComponent(entity)
		comp.AddCoupling(comp, "in1", entity, "in")
		comp.AddCoupling(entity, "out", comp, "out")
	case modeling.Coupled:
		comp.AddComponent(entity)
		comp.AddCoupling(comp, "in1", entity, "in1")
		comp.AddCoupling(entity, "out", comp, "out")
		// 增加原子模型
		n := w - 1
		line := make([]modeling.Atomic, n)
		mat := make([][]modeling.Atomic, n)
		for i := 0; i < n; i++ {
			mat[i] = make([]modeling.Atomic, n)
		}
		for i := 0; i < n; i++ {
			line[i] = NewDEVStoneAtomic(fmt.Sprintf("atomic_line_%v_%v", l, i))
			comp.AddComponent(line[i])
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if j >= i {
					mat[i][j] = NewDEVStoneAtomic(fmt.Sprintf("atomic_mat_%v_%v_%v", l, i, j))
					comp.AddComponent(mat[i][j])
				}
			}
		}
		// 增加耦合
		for i := 0; i < n; i++ {
			comp.AddCoupling(comp, "in2", line[i], "in")
			comp.AddCoupling(comp, "in2", mat[i][i], "in")
			comp.AddCoupling(line[i], "out", entity, "in2")
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				comp.AddCoupling(mat[0][i], "out", line[j], "in")
			}
		}
		for i := 1; i < n; i++ {
			for j := 0; j < n; j++ {
				if j >= i {
					comp.AddCoupling(mat[i][j], "out", mat[i-1][j], "in")
				}
			}
		}
	}
	return comp
}

func NewDEVStoneModel(typ string, depth int, width int) modeling.Coupled {
	if typ == "LI" {
		var model modeling.Entity = NewDEVStoneAtomic("atomic_l0")
		for i := 0; i < depth; i++ {
			if i == 0 {
				model = NewDEVStoneLICoupled("coupled_l"+strconv.Itoa(i), model)
			} else {
				atomics := make([]modeling.Atomic, 0)
				for j := 0; j < width-1; j++ {
					atomics = append(atomics, NewDEVStoneAtomic(fmt.Sprintf("atomic_l%v_i%v", i, j)))
				}
				model = NewDEVStoneLICoupled("coupled_l"+strconv.Itoa(i), model, atomics...)
			}
		}
		return model.(modeling.Coupled)
	} else if typ == "HI" {
		var model modeling.Entity = NewDEVStoneAtomic("atomic_l0")
		for i := 0; i < depth; i++ {
			if i == 0 {
				model = NewDEVStoneHICoupled("coupled_l"+strconv.Itoa(i), model)
			} else {
				atomics := make([]modeling.Atomic, 0)
				for j := 0; j < width-1; j++ {
					atomics = append(atomics, NewDEVStoneAtomic(fmt.Sprintf("atomic_l%v_i%v", i, j)))
				}
				model = NewDEVStoneHICoupled("coupled_l"+strconv.Itoa(i), model, atomics...)
			}
		}
		return model.(modeling.Coupled)
	} else if typ == "HO" {
		var model modeling.Entity = NewDEVStoneAtomic("atomic_l0")
		for i := 0; i < depth; i++ {
			if i == 0 {
				model = NewDEVStoneHOCoupled("coupled_l"+strconv.Itoa(i), model)
			} else {
				atomics := make([]modeling.Atomic, 0)
				for j := 0; j < width-1; j++ {
					atomics = append(atomics, NewDEVStoneAtomic(fmt.Sprintf("atomic_l%v_i%v", i, j)))
				}
				model = NewDEVStoneHOCoupled("coupled_l"+strconv.Itoa(i), model, atomics...)
			}
		}
		return model.(modeling.Coupled)
	} else if typ == "HOmod" {
		var model modeling.Entity = NewDEVStoneAtomic("atomic_l0")
		for i := 0; i < depth; i++ {
			model = NewDEVStoneHOmodCoupled("coupled_l"+strconv.Itoa(i), model, width, i)
		}
		return model.(modeling.Coupled)
	}
	return nil
}
