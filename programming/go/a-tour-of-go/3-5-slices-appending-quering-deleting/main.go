package main

import "fmt"

type Vertex struct {
	x, y int
}

func printSlice(s []Vertex) {
	for i, v := range s {
		fmt.Printf("%d: %d \n", i, v)
	}
}

func remove(s []Vertex, i int) []Vertex {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	s := make([]Vertex, 0, 0)
	s = append(s, Vertex{1, 2}, Vertex{14, 5})
	s = append(s, Vertex{69, 1})
	s = append(s, Vertex{49, 95})
	s = append(s, Vertex{643, 97})
	s = append(s, Vertex{1, 26})
	printSlice(s)
	println("Removal")
	s = remove(s, 1)
	printSlice(s)
}
