# Intro

In go, **arrays** behave a little differenly than in other languages. 
And also we have something called **slice** in golang which is like a reference to an array.

In this article only arrays will be studied.

# Definition
An array is contiguous collection of variables of the same type. It is ordered and this order is reflected in the memory.

Syntax
```golang
var_name := [size]type{a1, a2, a3 ... an}
```

e.g.
```golang
sample := [4]int{1, 2, 3, 4}
```

Both `size` and `a1, a2, a3, ... an` (actual elements) are optional in the array declaration

below we can see 4 valid array declarations

```golang
var ar1 = [2]int{1, 2} // both present
var ar2 = [2]int{}     // no elements initialization (so they are filled with zero-value)
var ar3 = [...]int{1, 2, 3} // no specifying the length. The symbol `...` tells the compiler to calculate the lenght based on list of elements
var ar4 = [...]int{} // same as above but in this case the calculated lenght will be `0`
```

# Accessing array elements

Same as in any other language (with `[]`)

# Arrays are value in go
An array is value type in go (the other type is reference type). <br>
It means that array variable name is not a pointer to the array. In fact it denotes the entire array.

So each time when assigning it to a new variable a new copy of array will be created.

# Iterating an array

```golang
package main

import "fmt"

func main() {
    letters := [3]string{"a", "b", "c"}
    //Using for loop
    fmt.Println("Using for loop")
    len := len(letters)
    for i := 0; i < len; i++ {
        fmt.Println(letters[i])
    }
    //Using for-range operator
    fmt.Println("\nUsing for-range loop")
    for i, letter := range letters {
        fmt.Printf("%d %s\n", i, letter)
    }
}
```

# Multidimensional arrays

Syntax:
```golang
sample := [x][y]{type}{{a11, a12 .. a1y},
                       {a21, a22 .. a2y},
                       {.. },
                       {ax1, ax2 .. axy}} 
```

# Conclusion
In GO array is some list of fixed size and of fixed values. Like the list of continents of the world. <br>
The more often used concept in every language was dynamic arrays or C++-called vectors, <br>
this is why Go introduced separate concept which is SLICE.