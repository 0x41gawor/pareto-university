package main

import (
	"fmt"
	"math"
)

// <--------------------------------------------- P O I N T E R S --------------------------------------------->

// Go has pointers. A pointer holds the memory address of a value.
//
// The type *T is a pointer to a T value. Its zero value is nil.

var p *int

func main1() {
	// The & operator generates a pointer to its operand.
	i := 42
	p = &i
	// The * operator denotes the pointer's underlying value.
	fmt.Println(*p) // read i through the pointer p
	*p = 21         // set i through the pointer p
	// This is known as "dereferencing" or "indirecting".
	// Unlike C, Go has no pointer arithmetic.
	main5()
}

// <--------------------------------------------- S T R U C T S --------------------------------------------->
type Vertex struct {
	X int
	Y int
}

// Struct fields can be accessed through a struct pointer.
func main2() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

// Struct fields can be accessed through a struct pointer.
//
// To access the field X of a struct when we have the struct pointer p we could write (*p).X.
//  However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

// <--------------------------------------------- A R R A Y S --------------------------------------------->
func main3() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// An array's length is part of its type, so arrays cannot be resized.
//This seems limiting, but don't worry; Go provides a convenient way of working with arrays.

// <--------------------------------------------- S L I C E S --------------------------------------------->

//An array has a fixed size.
//A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
//In practice, slices are much more common than arrays.
// The type []T is a slice with elements of type T.
func main4() {
	//A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	// a[low : high]
	// This selects a half-open range which includes the first element, but excludes the last one.
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

// --- Slices are like references to arrays ---
// A slice does not store any data, it just describes a section of an underlying array.
//
//Changing the elements of a slice modifies the corresponding elements of its underlying array.
//
//Other slices that share the same underlying array will see those changes.

func main5() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// --- Slice literals ---
// This is an array literal:
//
//[3]bool{true, true, false}
//And this creates the same array as above, then builds a slice that references it:
//
//[]bool{true, true, false}

func main6() {
	q := []int{2, 3, 5, 7, 11, 13} //this builds an array q with slice that references it
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	} // this also build array and slice but with type []struct{ i int, b bool}
	fmt.Println(s)
}

// --- Slice defaults ---
// When slicing, you may omit the high or low bounds to use their defaults instead.
//
//The default is zero for the low bound and the length of the slice for the high bound.

// For the array
//
// 	var a [10]int
//these slice expressions are equivalent:
//
//	a[0:10]
//	a[:10]
//	a[0:]
//	a[:]

func main7() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

// --- Slice length and capacity ---
// A slice has both a length and a capacity.
//
//The length of a slice is the number of elements it contains.
//
//The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
//
//The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

// --- Zero value of a slice ---
// The zero value of a slice is nil.
//
//A nil slice has a length and capacity of 0 and has no underlying array.

//  ------------ Creating a slice with make ------------
// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
// The `make` function allocates a zero-valued array and returns a slice that refers to that array:
// `b := make([]int, 0, 5)` // len(b)=0, cap(b)=5
func main112() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// ---------- Append to a slice ------------
// It is common to append new elements to a slice, and so Go provides a built-in append function.
// https://pkg.go.dev/builtin#append

// `func append(s []T, vs ...T) []T`

// The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.
//
//The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
//
//If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.
// https://go.dev/blog/slices-intro <--- more here

func main9() {
	var s []int
	printSlice2(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice2(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice2(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice2(s)
}

func printSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// <--------------------------------------------- R A N G E--------------------------------------------->

// The range form of the for loop iterates over a slice or map.
//
// When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

// You can skip the index or value by assigning to _.
// `for i, _ := range pow`
// `for _, value := range pow`

func printSlice3(s [][]uint8) {
	for i, v := range s {
		fmt.Printf("%d len=%d cap=%d %v\n", i, len(v), cap(v), v)
	}
}

// https://go.dev/tour/moretypes/18
func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy, dy)
	for y, _ := range result {
		result[y] = make([]uint8, dx, dx)
		for x, _ := range result[y] {
			result[y][x] = uint8(x * y)
		}
	}
	return result
}

func main10() {

	printSlice3(Pic(5, 5))
}

// <--------------------------------------------- M A P S --------------------------------------------->

// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
//
//The make function returns a map of the given type, initialized and ready for use

type Vertex2 struct {
	Lat, Long float64
}

var m = map[string]Vertex2{
	"Bell Labs": Vertex2{
		40.68433, -74.39967,
	},
	"Google": Vertex2{
		37.42202, -122.08408,
	},
}

var m2 = map[string]Vertex2{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main12() {
	fmt.Println(m2)
}

// Insert or update an element in map m:
//
//m[key] = elem
//Retrieve an element:
//
//elem = m[key]
//Delete an element:
//
//delete(m, key)
//Test that a key is present with a two-value assignment:
//
//elem, ok = m[key]
//If key is in m, ok is true. If not, ok is false.
//
//If key is not in the map, then elem is the zero value for the map's element type.
//

func main11() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// <--------------------------------------------- F U N C T I O N   V A L U E S --------------------------------------------->

// Functions are values too. They can be passed around just like other values.
//
//Function values may be used as function arguments and return values.

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main13() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

// <--------------------------------------------- F U N C T I O N   C L O S U R E S --------------------------------------------->

// Go functions may be closures. A closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
//
// For example, the adder function returns a closure. Each closure is bound to its own sum variable.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main14() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-i),
		)
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	now := 0
	last := now
	return func() int {
		now++
		last = now
		return last + now
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
