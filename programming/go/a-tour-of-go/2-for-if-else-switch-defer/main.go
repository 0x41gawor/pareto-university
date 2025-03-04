package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// <--------------------------------------------- F O R --------------------------------------------->
//Go has only one looping construct, the for loop.
//The basic for loop has three components separated by semicolons:
//
//the init statement: executed before the first iteration
//the condition expression: evaluated before every iteration
//the post statement: executed at the end of every iteration

// The init and post statements are optional.
func main2() {
	sum := 1
	// for ; sum < 1000; {
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// At that point you can drop the semicolons: C's while is spelled for in Go.

// <--------------------------------------------- I F --------------------------------------------->

// Go's if statements are like the expression need not be surrounded by parentheses ( ) but the braces { } are required.
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// Like for, the if statement can start with a short statement to execute before the condition.
//
// Variables declared by the statement are only in scope until the end of the if.
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// Variables declared inside an if short statement are also available inside any of the else blocks.
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

// <--------------------------------------------- S W I T C H --------------------------------------------->

// Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow.
// In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go.
//Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.
func main4() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

// default
func main5() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// Switch without a condition is the same as switch true.
//
// This construct can be a clean way to write long if-then-else chains.
func main6() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// <--------------------------------------------- D E F E R --------------------------------------------->

// A defer statement defers the execution of a function until the surrounding function returns.
//
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

// STACKING DEFERS
// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
func main7() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
