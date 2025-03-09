package astar

import (
	"a-star/internal/model"
	"slices"
	"sort"
)

// isValidNode checks if the total weight of packed items does not exceed the backpack's capacity
func IsValidNode(n model.Node, backpack model.Backpack) bool {
	totalWeight := 0

	for i, state := range n {
		if state == model.Packed { // Only count items that are packed
			totalWeight += backpack.Items[i].Weight
		}
	}

	return totalWeight <= backpack.Capacity
}

// isNodeTerminal checks if a Node has any Undecided (-1) items
func IsNodeTerminal(n model.Node) bool {
	return !slices.Contains(n, model.Undecided) // No undecided items, node is terminal
}

// neighbors finds the neighbor-nodes of given node n
func Neighbors(n model.Node) (model.Node, model.Node) {
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
func Oracle(n model.Node, backpack model.Backpack) float64 {
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
