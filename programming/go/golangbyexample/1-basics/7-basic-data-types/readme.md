# Intro
Data types in go can be categorized into two:
- basic types
- composite types

- Basic Types
  - Integers (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64)
  - Floats (float32, float64)
  - Complex numbers (complex32, complex64)
  - Byte
  - Rune
  - String
  - Boolean
- Composite Types
  - Collection/Aggregation or Non-Reference Types
    - Arrays
    - Structs
  - Reference Types
    - Slices
    - Maps
    - Channels
    - Pointers
    - Functions/Methods
  - Interface
    - Special case of empty interface

## Complex numbers
Two types:
- complex64
  - both real and imaginary part ar float32
- complex128
  - both are float64

#### Initialization
```go
a := complext(5, 4)
// or
a := 5 +4i
```

## Byte
`byte` is an alias for `uint8`.

Range is 0-255 so it can represent ASCII character and since go has no `char` type it can be used as such
```go
var a byte = 'a'
```

## Rune
`rune` is an alias for `uint32`

So it is integer value, and this value is meant to represent a **Unicode Code Point**.

#### What is Unicode
It is superset of ASCII charset, which assigns unique number to every character that exists.

e.g.
A pound symbol `£` is represented by Unicode Point `U+00A3 (dec 163)`

But Unicode doesn't specify how to store these code points in memory. This is where **utf-8** comes into picture.

https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/

#### UTF-8

**utf-8** saves every Unicode Point using 1,2,3 or 4 bytes.
ASCII points are stored using 1 byte. 

So this is why `rune` is alias for `int32` (because it can store any Unicode Point).

In GO every string is encoded using utf-8.

```go
rPound := '£'
fmt.Printf("Type: %s\n", reflect.TypeOf(rPound)) // will be uint32
fmt.Printf("Unicode CodePoint: %U\n", rPound) // will be U+00A3
fmt.Printf("Character: %c\n", r) // will be £ 
```

## String
`string` is read-only slice of bytes in golang

There are two ways to initialize
```go
a := "text"
a := `text`
```

string in `" "` honors the escape sequences (e.g "\n"), string in (``) does not

string is in utf-8 which uses 1 byte to save "a" or "b" byt 2 bytes to store "£".

```go
s := "ab£"
fmt.Println([]byte(s)) // prints [48 98 194 163]
```

Also when you try to print the length of the above string using `len(“ab£”)`, it will output 4 and not 3 because it contains 4 bytes. 

Also `range` loops over sequences of bytes which for each character, soo...
```go
s := "ab£"
for _, c := range s {
   fmt.Println(string(c))
} // will print a b £
```

You can use `+` on strings for concatenation

## bool

default value is `false`

operations are:
- AND
  - `&&`
- OR
  - `||`
- NOT
  - `!`

