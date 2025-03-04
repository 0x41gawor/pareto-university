package main

import "fmt"

type area func(int, int) float32

func areaAsRect(a int, b int) float32 {
	return float32(a * b)
}

func areaAsTri(a int, h int) float32 {
	return float32(a * h / 2)
}

func main() {
	// we have two lengths, what area based on shape builded with these will we get?
	var a = 4
	var b = 5
	var areaF area
	//
	sel := "tri" // other option is "rect"
	switch sel {
	case "tri":
		areaF = areaAsTri
	case "rect":
		areaF = areaAsRect
	}
	fmt.Printf("Area is: %f", areaF(a, b))
}
