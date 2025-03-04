package main

import (
	"fmt"

	"example.com/m/math"
	"example.com/m/math/advanced"
)

func main() {
	fmt.Println(math.Add(2, 1))
	fmt.Println(math.Subtract(2, 1))
	fmt.Println(advanced.Square(2))
}
