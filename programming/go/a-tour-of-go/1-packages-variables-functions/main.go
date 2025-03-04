package main

//Every Go program is made up of packages.
//
//Programs start running in package main.
//
//This program is using the packages with import paths "fmt" and "math/rand".
//
//By convention, the package name is the same as the last element of the import path.
//For instance, the "math/rand" package comprises files that begin with the statement package rand.
import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	//In Go, a name is exported if it begins with a capital letter.
	//For example, Pi is an exported name,  which is exported from the math package.
	// When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.
	fmt.Println(math.Pi)
	fmt.Println(add(12, 5))
	show()
}

// <----------------------------------------------- F U N C T I O N S ----------------------------------------------->

// Notice that the type comes after the variable name.
func add(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func add2(x, y int) int {
	return x + y
}

// A function can return any number of results.
func swap(x, y string) (string, string) {
	return y, x
}

// N A K E D   R E T U R N S
// Go's return values may be named. If so, they are treated as variables defined at the top of the function.
// A return statement without arguments returns the named return values. This is known as a "naked" return.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// <----------------------------------------------- V A R I A B L E S ----------------------------------------------->
// The var statement declares a list of variables; as in function argument lists, the type is last
var c, python, java bool

// A var declaration can include initializers, one per variable.
//If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
var i, j int = 1, 2
var isEnabled, isDebug, title = true, false, "Szlagier"

// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
// Outside a function every statement begins with a keyword so it is impossible in such cases
func example() int {
	k := 3
	return k
}

//bool
//
//string
//
//int  int8  int16  int32  int64
//uint uint8 uint16 uint32 uint64 uintptr
//
//byte // alias for uint8
//
//rune // alias for int32
//     // represents a Unicode code point
//
//float32 float64
//
//complex64 complex128

// variable declarations may be "factored" into blocks, as with import statements
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func show() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// Variables declared without an explicit initial value are given their zero value: 0, false or ""

// Type conversions
var n int = 42
var f float64 = float64(n)
var u uint = uint(f)

// The expression T(v) converts the value v to the type T.

// Constants
// Constants cannot be declared using the := syntax.
const Pi = 3.14
