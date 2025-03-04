# Intro

What is switch? It is just a way to prevent a if-else ladder :D

```golang
switch statement; expression {
case expression1:
     //Dosomething
case expression2:
     //Dosomething
default:
     //Dosomething
}
```

> statement - states something, e.g. it states that variable `a` is equal to `5`
> expression - express something e.g. the letter "b"

Both `statement` and `expression` are optional so we have 4 cases:

## Both present

```golang
 switch ch := "b"; ch {
    case "a":
        fmt.Println("a")
    case "b", "c":
        fmt.Println("b or c")    
    default:
        fmt.Println("No matching character")    
    }
```

variable `ch` is present only in the switch scope

## Both absent
```golang
    age := 45
    switch {
    case age < 18:
        fmt.Println("Kid")
    case age >= 18 && age < 40:
        fmt.Println("Young")
    default:
        fmt.Println("Old")
    }
```

## Only statement
```golang
switch age := 29; {
    case age < 18:
        fmt.Println("Kid")
    case age >= 18 && age < 40:
        fmt.Println("Young")
    default:
        fmt.Println("Old")
    }
```

## Only expresssion
```golang
char := "a"
    switch char {
    case "a":
        fmt.Println("a")
    case "b":
        fmt.Println("b")
    default:
        fmt.Println("No matching character")
    }
```

# Fallthrough keyword
```golang
i := 45
	switch {
	case i < 10:
		fmt.Println("i is less than 10")
	case i < 50:
		fmt.Println("i is less than 50")
		fallthrough
	case i < 100:
		fmt.Println("i is less than 100")
	}
```
Even if something matches, the switch will fall through anyway (so it can maybe match the next value).

`fallthrough` needs to be final statement in single case block

# Break
The opposite of `fallthrough` is `break`.

It will terminate the switch execution.

e.g. this code:
```golang
switch char := "b"; char {
    case "a":
        fmt.Println("a")
    case "b":
        fmt.Println("b")
        break
        fmt.Println("after b")
    default:
        fmt.Println("No matching character")
    }
```
will show `"b"` only. 

# Type switch
You can infer the type of Any using switch
```golang
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
```