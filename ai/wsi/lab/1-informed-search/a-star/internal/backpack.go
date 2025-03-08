package internal

import "fmt"

// Item represents single item in backpack
type Item struct {
	Profit int
	Weight int
}

// Backpack represents backpack, which can have its items and has specified capacity
type Backpack struct {
	Capacity int
	Items    []Item
}

// ItemState represents the state of an item in the backpack
type ItemState int

// String method to customize how ItemState is printed
func (s ItemState) String() string {
	switch s {
	case Undecided:
		return "?"
	case Packed:
		return "1"
	case Unpacked:
		return "0"
	default:
		return fmt.Sprintf("%d", int(s)) // Fallback for unexpected values
	}
}

const (
	Undecided ItemState = -1 // Item hasn't been processed yet
	Packed    ItemState = 1  // Item is included in the backpack
	Unpacked  ItemState = 0  // Item is not included in the backpack
)
