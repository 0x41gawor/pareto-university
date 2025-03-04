package main

import "fmt"

type sample struct {
	name string
}

func main() {
	var t = 123                  //Type Inferred will be int
	var u = "circle"             //Type Inferred will be string
	var v = 5.6                  //Type Inferred will be float64
	var w = true                 //Type Inferred will be bool
	var x = 'a'                  //Type Inferred will be rune
	var y = 3 + 5i               //Type Inferred will be complex128
	var z = sample{name: "test"} //Type Inferred will be main.Sample

	fmt.Printf("Type: %T Value: %v\n", t, t) // with %T in printf you can obtain the type of var
	fmt.Printf("Type: %T Value: %v\n", u, u)
	fmt.Printf("Type: %T Value: %v\n", v, v)
	fmt.Printf("Type: %T Value: %v\n", w, w)
	fmt.Printf("Type: %T Value: %v\n", x, x)
	fmt.Printf("Type: %T Value: %v\n", y, y)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
