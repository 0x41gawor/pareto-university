package main

import (
	model "a-star/internal"
	"fmt"
	"sort"
	"strconv"
	"strings"

	pq "github.com/jupp0r/go-priority-queue"
)

type Node []model.ItemState

// Converts a Node slice into a unique string representation
func (n Node) String() string {
	var sb strings.Builder
	sb.WriteString("[") // Start with an opening bracket
	for i, state := range n {
		if i > 0 {
			sb.WriteString(",") // Add a comma before each element except the first
		}
		if state == -1 {
			sb.WriteString("?") // Replace -1 with "?"
		} else {
			sb.WriteString(fmt.Sprintf("%d", state)) // Convert each state to string
		}
	}
	sb.WriteString("]") // Close with a bracket
	return sb.String()
}

// Parses a string representation of Node back into a Node slice
func parseNode(s string) Node {
	var n Node
	// Remove the brackets before splitting
	trimmed := strings.Trim(s, "[]")
	elements := strings.Split(trimmed, ",")
	for _, elem := range elements {
		if elem == "?" {
			n = append(n, -1) // Replace "?" with -1
		} else {
			// Convert string to integer
			state, err := strconv.Atoi(elem)
			if err != nil {
				fmt.Printf("Error parsing element '%s': %s\n", elem, err)
				continue
			}
			n = append(n, model.ItemState(state))
		}
	}
	return n
}

func main() {
	items := []model.Item{
		{Profit: 16, Weight: 8},
		{Profit: 8, Weight: 3},
		{Profit: 9, Weight: 5},
		{Profit: 6, Weight: 2},
	}
	b1 := model.Backpack{
		Capacity: 9,
		Items:    items,
	}
	s0 := Node{model.Undecided, model.Undecided, model.Undecided, model.Undecided}
	A := pq.New()
	A.Insert(s0.String(), oracle(s0, b1))

	for {
		fmt.Println("------")
		xInterface, err := A.Pop()
		if err != nil {
			fmt.Printf("Error while popping node %s", err.Error())
		}
		xStr, ok := xInterface.(string)
		if !ok {
			panic("Failed to assert type Node from A")
		}
		x := parseNode(xStr)
		fmt.Printf("x: %s\n", x.String())
		if isNodeTerminal(x) {
			break
		}
		y1, y2 := neighbors(x)
		fmt.Print("Neighbors: ")
		fmt.Print(y1, " -> ", oracle(y1, b1), " ; ")
		fmt.Println(y2, "->", oracle(y2, b1))
		if isValidNode(y1, b1) {
			A.Insert(y1.String(), -oracle(y1, b1))
		}
		if isValidNode(y2, b1) {
			A.Insert(y2.String(), -oracle(y2, b1))
		}
	}
}

// neighbors finds the neighbor-nodes of given node n
func neighbors(n Node) (Node, Node) {
	// Iterate to find the first Undecided item
	for i, state := range n {
		if state == model.Undecided {
			// Create copies of the original Node
			packedNode := append([]model.ItemState(nil), n...)   // Copy of original node
			excludedNode := append([]model.ItemState(nil), n...) // Copy of original node

			// Modify the copies
			packedNode[i] = model.Packed
			excludedNode[i] = model.Unpacked

			return packedNode, excludedNode
		}
	}

	// If no Undecided state is found, return empty slices (should not happen in valid input)
	return nil, nil
}

// oracle estimates the best possible profit by filling the backpack greedily based on profit-to-weight ratio.
func oracle(n Node, backpack model.Backpack) float64 {
	totalProfit := 0.0
	totalWeight := 0

	i := 0

	for {
		if n[i] == model.Undecided {
			break
		} else {
			totalProfit += float64(int(n[i])) * float64(backpack.Items[i].Profit)
			totalWeight += int(n[i]) * backpack.Items[i].Weight
		}
		i++
		if i >= len(n) {
			break
		}
	}

	if i > len(n)-1 {
		return totalProfit
	}

	type ItemWithRatio struct {
		ItemNo int
		Ratio  float64
	}

	var itemsWithRatio []ItemWithRatio

	// calculate ratio for left items
	for {
		ratio := float64(backpack.Items[i].Profit) / float64(backpack.Items[i].Weight)
		itemsWithRatio = append(itemsWithRatio, ItemWithRatio{i, ratio})
		i++
		if i > len(n)-1 {
			break
		}
	}
	// Sort items by profit-to-weight ratio in descending order
	sort.Slice(itemsWithRatio, func(i, j int) bool {
		return itemsWithRatio[i].Ratio > itemsWithRatio[j].Ratio
	})

	remainingCapacity := backpack.Capacity - totalWeight

	for _, item := range itemsWithRatio {
		if remainingCapacity >= backpack.Items[item.ItemNo].Weight {
			// Take the whole item
			totalProfit += float64(backpack.Items[item.ItemNo].Profit)
			remainingCapacity -= backpack.Items[item.ItemNo].Weight
		} else {
			// Take the fraction of the item that fits
			fraction := float64(remainingCapacity) / float64(backpack.Items[item.ItemNo].Weight)
			totalProfit += fraction * float64(backpack.Items[item.ItemNo].Profit)
			break // no more space left
		}
	}

	return totalProfit
}

// isNodeTerminal checks if a Node has any Undecided (-1) values
func isNodeTerminal(n Node) bool {
	for _, state := range n {
		if state == model.Undecided {
			return false // Found an undecided item, so it's not terminal
		}
	}
	return true // No undecided items, node is terminal
}

// isValidNode checks if the total weight of packed items does not exceed the backpack's capacity
func isValidNode(n Node, backpack model.Backpack) bool {
	totalWeight := 0

	for i, state := range n {
		if state == model.Packed { // Only count items that are packed
			totalWeight += backpack.Items[i].Weight
		}
	}

	return totalWeight <= backpack.Capacity
}
