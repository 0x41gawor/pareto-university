# Intro
Golang has only two loops:
- for loop
- for-range loop

This file covers only **for loop** loop.

## Syntax

```golang
for init_part; condition_part; post_part {

}
```

We have three parts:
- **init_part** - is executed <u>before 1st</u> iteration
- **condition_part** - is executed <u>before every</u> iteration. If condition is false the loop will exit otherwise continue
- **post_part** - is executed <u>after every</u> iteration


some points to note:
- parenthesis around these three parts is not necessary
- the **init_part** and **post_part** is optional
- the **init_part** can be any statement with:
  - short declaration,
  - function call,
  - assignment
- if **init_part** has variable declaration the scope of variable is the loop
- **post_part** can be any statement, but cannot contain initialization

## Examples
### THe simplest for loop
```golang
    for i:=0; i<5; i++ {
        fmt.Println(i)
    }
```
### Only condition for loop
```golang
    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }
```
### Not even condition for loop
```golang
    i := 0
    for {
        fmt.Println(i)
        i++
        time.Sleep(time.Second * 1)
    }
```
> So **condition_part** is not necessary then?

### Break and Continue
ofc. u can use it

### Function call and assignment in **Init_part**
```golang
package main

import "fmt"

func main() {
    i := 1
    //Function call in the init part in for loop
    for test(); i < 3; i++ {
        fmt.Println(i)
    }

    //Assignment in the init part in for loop
    for i = 2; i < 3; i++ {
        fmt.Println(i)
    }
}
func test() {
    fmt.Println("In test function")
}
```