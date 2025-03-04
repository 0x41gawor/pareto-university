https://golangbyexample.com/variables-in-golang-complete-guide/
# Variables

## What is variable

**variable** - name of the ***memory location***.

This ***memory location*** can store a ***value*** of some ***type***.

**Type** defines:
- size (in bytes) of this memory location
- the range in which values can be (math.domain)
- operations defined on that variable
- zero value (default value of variable)

## Declaring a variable
### Without initial value
```go
var <name> <type>
```
e.g.
```go
var a int
```
The variable will take **zero value**.
### With initial value
```go
var <name> <type> = <value>
```

e.g.
```go
var a int = 8
```

### Multiple variables
This is valid
```go
var a, b, c int
var z, x, y int = 4, 5, 6
```

### Variables of diff types
```go
var (
        aaa int
        bbb int    = 8
        ccc string = "a"
    )
```

### Type inference
If the variable has an initial value type declaration can be omitted, because GO compiler will figure out the type based on assigned value. This is called **Type inference**.
```go
var <name> = <value>
```

Check out `main.go` file

#### Short variable declaration
With `:=` both `var` keyword and type info can be omitted.

e.g.
```go
a : = 4
text := "elo"
count, counter := 0, 10
```

Good to know:
- `:=` operator is only available within a function, outside you can't use it
- variable once declared with `:=` cannot be re-declared with `:=`

In case of multiple declaration, := can also be used again for a particular variable if at least one of the variables on left hand side is new. See below example.

```go
package main

import "fmt"

func main() {
    a, b := 1, 2
    b, c := 3, 4
    fmt.Println(a, b, c)

```

## Scope of a variable

You can define variables on three levels:
- package level
- function level
- block level

And we divide variables on **local** and **global**.

### Local variable
- are defined in a scope of a *block* or *function*
- only live till the end of block/func in which they are declared. After that, they are garbage collected
- we say that variable "is local to" some block/func

### Global variable
- it is declared outside the scope of any func/block, and by the convention at the top of a file
- we say that variable "is global within" some package
- lower/upper-case of variable name denotes if it is exported or un-exported (accessible in other packages or not)
- global variable lives through the lifetime of a program


