# Intro

Pointer is a variable that holds a memory address of other variable.

```golang
var ex *T
```

The **zero-value** of pointer is `nil`.

## Initialization
There are two ways:
- the `new` keyword, which is more like a allocating a placeholder for a variable of some type
- using `&` syntax, which is getting the address of already existing variable

### New
```golang
a := new(int) // a is a memory address of memory location which can be filled with int
```

### &
```golang
b := 5
a := &b // a is a memory address of b
```

## Dereferencing a pointer
To dereference (wyłuskać) a pointer means to obtain the variable which resides under the memory location which pointer variable actually holds.

```golang
a := new(int) // reservation of memory location for int
*a = 5 // putting a 5 into that memory location
fmt.Println(a) // prints memory location e.g 0xFFDSFD
fmt.Println(*a) // prints 5
```

## Pointer to pointer
synax is hierarchical

## Pointer arithmetic
There is no pointer arithmetic in Go.