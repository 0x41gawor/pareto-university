package main

import (
	"a-star/internal/astar"
	"a-star/internal/model"
	"a-star/internal/pqueue"
	"fmt"
)

func main() {
	items := []model.Item{
		{Profit: 16, Weight: 8},
		{Profit: 8, Weight: 3},
		{Profit: 9, Weight: 5},
		{Profit: 6, Weight: 2},
		{Profit: 10, Weight: 3},
	}
	b1 := model.Backpack{
		Capacity: 12,
		Items:    items,
	}

	s0 := model.Node{model.Undecided, model.Undecided, model.Undecided, model.Undecided, model.Undecided}
	A := pqueue.New()
	A.Insert(s0, astar.Oracle(s0, b1))
	for {
		x := A.Pop()
		fmt.Printf("x: %s\n", x.String())
		if astar.IsNodeTerminal(x) {
			break
		}
		y1, y2 := astar.Neighbors(x)
		fmt.Print("Neighbors: ")
		fmt.Print(y1, " -> ", astar.Oracle(y1, b1), " ; ")
		fmt.Println(y2, "->", astar.Oracle(y2, b1))
		if astar.IsValidNode(y1, b1) {
			A.Insert(y1, astar.Oracle(y1, b1))
		}
		if astar.IsValidNode(y2, b1) {
			A.Insert(y2, astar.Oracle(y2, b1))
		}
	}
}
