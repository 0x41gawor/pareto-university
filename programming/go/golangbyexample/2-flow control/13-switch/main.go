package main

import "fmt"

func main() {
	printType("test_string")
	printType(2)
}

func printType(t any) {
	switch v := t.(type) {
	case string:
		fmt.Println("Type: string")
	case int:
		fmt.Println("Type: int")
	default:
		fmt.Printf("Unknown Type %T", v)
	}
}
