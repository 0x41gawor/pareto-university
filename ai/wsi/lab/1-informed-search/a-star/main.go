package main

import (
	bp "a-star/internal"
	"sort"
)

type Node []bp.ItemState

func main() {
	s0 := Node{bp.Undecided, bp.Undecided, bp.Undecided, bp.Undecided}

	// Inicjalizacja przedmiotów
	items := []bp.Item{
		{Profit: 16, Weight: 8},
		{Profit: 8, Weight: 3},
		{Profit: 9, Weight: 5},
		{Profit: 6, Weight: 2},
	}

	// Tworzymy plecak (opcjonalnie z określoną pojemnością, np. 10)
	b1 := bp.Backpack{
		Capacity: 9, // Można dostosować pojemność według potrzeb
		Items:    items,
	}

	println(oracle(s0, b1))
}

// neighbors finds the neighbor-nodes of given node n
func neighbors(n Node) (Node, Node) {
	// Iterate to find the first Undecided item
	for i, state := range n {
		if state == bp.Undecided {
			// Create copies of the original Node
			packedNode := append([]bp.ItemState(nil), n...)   // Copy of original node
			excludedNode := append([]bp.ItemState(nil), n...) // Copy of original node

			// Modify the copies
			packedNode[i] = bp.Packed
			excludedNode[i] = bp.Unpacked

			return packedNode, excludedNode
		}
	}

	// If no Undecided state is found, return empty slices (should not happen in valid input)
	return nil, nil
}

// oracle estimates the best possible profit by filling the backpack greedily based on profit-to-weight ratio.
func oracle(n Node, backpack bp.Backpack) float64 {
	type ItemRatio struct {
		Item  bp.Item
		Ratio float64
	}

	var itemsWithRatio []ItemRatio

	// Compute profit-to-weight ratio and store copies of items
	for i, item := range backpack.Items {
		if n[i] == bp.Unpacked { // Ignore items that are already Packed
			continue
		}
		ratio := float64(item.Profit) / float64(item.Weight)
		itemsWithRatio = append(itemsWithRatio, ItemRatio{Item: item, Ratio: ratio})
	}

	// Sort items by profit-to-weight ratio in descending order
	sort.Slice(itemsWithRatio, func(i, j int) bool {
		return itemsWithRatio[i].Ratio > itemsWithRatio[j].Ratio
	})

	// Fill the backpack greedily
	totalProfit := 0.0
	remainingCapacity := float64(backpack.Capacity)

	for _, item := range itemsWithRatio {
		if remainingCapacity >= float64(item.Item.Weight) {
			// Take the whole item
			totalProfit += float64(item.Item.Profit)
			remainingCapacity -= float64(item.Item.Weight)
		} else {
			// Take the fraction of the item that fits
			fraction := remainingCapacity / float64(item.Item.Weight)
			totalProfit += fraction * float64(item.Item.Profit)
			break // No more space left
		}
	}

	return totalProfit
}
