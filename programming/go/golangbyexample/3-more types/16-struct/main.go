package main

import "fmt"

type Employee struct {
	string
	int
	float32
}

func main() {
	emp := Employee{"Sam", 31, 2000.0}
	fmt.Println(emp)
}
