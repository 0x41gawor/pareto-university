# Intro
```golang
if condition {
   //Do something
}
```

You dont need any brackets around `condition`

## Condition 
Can be combination of any statements that result in boolean value. 

Can be composed of operators: &&, ||, >, <, >=, <=, ! etc.

```golang
    a := 4
    if a > 3 && a < 6 {
        fmt.Println("a is within range")
    }
```

## If with short Statement

If statement also supports a statement before the condition. This statement will be executed before the condition. There can also be new initialized variable in the statement. Below is the format for that.

```golang
if statement; condition {
   //Do something
}
```

```golang
    if a := giveMeA(); a > 5 {
        fmt.Println("a is greater than 5")
    } else {
        fmt.Println("a is less than 5")
    }

```

## Ternary operator
There is no ternay operator in go.