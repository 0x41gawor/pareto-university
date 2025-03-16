package model

import (
	"fmt"
	"strconv"
	"strings"
)

type Node []ItemState

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
func ParseNode(s string) Node {
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
			n = append(n, ItemState(state))
		}
	}
	return n
}
