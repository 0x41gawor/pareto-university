# Overview

In GO functions are **first-order variables**. They cn be passed around like any other variable.

# Function params
You can do smth like this:
```go
 func sum(a , b int) int {
    return a+b
 }
```

# Named return values
```golang
func calculateSumAndDifference(a, b float64) (sum float64, diff float64) {
    sum = a + b
    diff = a - b
    return // No need to explicitly specify variable names in the return statement
}
```

# Function usages

There are 3:
- Generic usage
- Function as Type
- Function as Values / Anonymous Functions

## Generic usage
This is the most basic usage that you know.
## Function as Type
In GO function is also a type. Two functions will be of the same type if:
- They have the same number and types of arguments
- They have the same number and types of return values

So basically this two function are of the same type
```go
func sum(a, b int) int {
    ///
}
func diff(a, b int) int {
    ///
}
```

**Function type** is useful in:
- **Higher-order functions** - these are functions that can take the function as argument of return the function
- In case of defining interfaces - only the *Function type* is specified

### User defined function types
You can define your function types with `type` keyword:
```golang
type myFunc func(int, int) int
```
## Function as Values / Anonymous Functions

You can use function in Go as first-order variable. So it can be used as value.

The second named is because a function is not named and can be assigned to any name (variable).

```golang
var max = func(a, b int) int {
    if a >= b {
        return a
    }
    return b
}
```

But for what it is?

Let's see this example:
```golang
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
```

## High order function

You can have a function that take other function of given type as argument and return some other function type

Real life example:
```golang
// We have such function type
type http.HandlerFunc func(http.ResponseWriter, *http.Request) error

// And function that takes such function and returns the same
// Decorator Design Pattern is used here
func withJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tokenStr := r.Header.Get("x-jwt-token")
		token, err := ValidateJWT(tokenStr)
        if err != nil {
            // do smth
        }
        handlerFunc(w, r)
    }
}
// Note that this function returns a function that at the end calls the given function which returns what? error - so the type is good
// On the higher abstraction level such function (withJWTAuth) just add some functionalty before calling the given function - Decorator Pattern
```

# Special usages of functions

## Function closures

These are anonymous function that are created in some higher-order function. The higher-order function declares some variables which:
- can be later accessed by anonymous function returned by the higher-order function
- are stored in the variable to which return statement of higher-order function is assigned

```go
// The higher-order function here is `getSquare()`
func getSquare() func(int) int {
    count := 0                      //
    return func(x int) int {
        count++
        fmt.Printf("Square func called %d times\n", count)

        return x*x
    }
}

func main() {
    // variable square stores the `count` and can be called as function 
    var square = getSquare()
    square(1)
    square(5)
    square(3)
}               // This program will have the output "Square func called 1 times\n Square func called 2 times\n Square func called 3 times\n"
```

## Higher-order functions
Either accept as argument or return a function.

```golang
package main

import "fmt"

func main() {
    areaF := getAreaFunc()
    print(3, 4, areaF)
}

// accepts as an argument
func print(x, y int, area func(int, int) int) {
    fmt.Printf("Area is: %d\n", area(x, y))
}
// returns 
func getAreaFunc() func(int, int) int {
    return func(x, y int) int {
        return x * y
    }
}
```

## Immediately Invoked Function (IIF)

```golang
package main

import "fmt"

func main() {
    squareOf2 := func() int {
        return 2 * 2
    }()

    fmt.Println(squareOf2)
}
```

> When you don't want to expose the logic of the function either within or outside the package. For eg let's say there is a function which is setting some value. You can encapsulate all the logic of setting in an IIF function. This function won't be available for calling either outside or within the package.

## Variadic function

Function can accept a dynamic number of arguments.

```golang
package main

import "fmt"

func main() {
    fmt.Println(add(1, 2))
    fmt.Println(add(1, 2, 3))
    fmt.Println(add(1, 2, 3, 4))
}

func add(numbers ...int) int {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}
```

## Methods

Is is a function that has a receiver argument.
```golang
func (receiver receiver_type) some_func_name(arguments) return_values
```

When you attach a function to a type, then that function becomes a method for that type. <br>A receiver can be a struct or any other type. <br>The method will have access to the properties of the receiver and can call the receiver's other methods.

Diff with funcs:
- Methods cannot be used as first-order object and be passed around like variables
- There can exist different methods with the same name with a different receiver, but there cannot exist two different functions with the same name in the same package.