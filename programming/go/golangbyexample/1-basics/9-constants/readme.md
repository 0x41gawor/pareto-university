# Value must be known at COMPILE TIME
Not runtime
e.g.
```golang
package main
const name = "test"
func main() {
    const a = getValue()
}
func getValue() int {
    return 1
}
```
It will produce an error. Because calling a function happens at RUNTIME not COMPILETIME.

# Typed and Untyped 
 In go constant are treated in a different way than any other language. GO has a very strong type system that doesn’t allow implicit conversion between any of the types. Even with the same numeric types no operation is allowed without explicit conversion. For eg you cannot add a int32 and int64 value. To add those either int32 has to be explicitly converted to int64 or vice versa. 

However untyped constant have the flexibility of temporary escape from the GO’s type system as we will see in this article.

## Typed constant
```golang
const a int32 = 8
```

This can be only assigned to a variable of type `int32`. Nothing special here.

## Untyped constant

This when you don't specift the type of constant.
```golang
const a = 123        //Default hidden type is int
```

ofc it has default hidden type, but the actual type (if its `int8`, `int32`, or `int64`) will be decided during runtime.

Now the question which comes to the mind is what is the use of untyped constant.  The use of untyped constant is that the type of the constant will be decided depending upon the type of variable they are being assigned to.  Sounds confusing? Let’s see with an example.

**Pi** constant value in math package is declared as below.
```golang
const Pi = 3.14159265358979323846264338327950288419716939937510582097494459
```


```golang
package main
import (
    "fmt"
    "math"
)
func main() {
    var f1 float32
    var f2 float64
    f1 = math.Pi
    f2 = math.Pi

    fmt.Printf("Type: %T Value: %v\n", math.Pi, math.Pi)
    fmt.Printf("Type: %T Value: %v\n", f1, f1)
    fmt.Printf("Type: %T Value: %v\n", f2, f2)
}
```

### Untyped const special case

```golang
type myString string // u create your own type (e.g. alias for type)

const typedConst string = "abc"
const untypedConst = "abc"

var a myString = typedConst // this is invalid :(
var b myString = untypedConst // this is valid :)
```

### Numeric Expressions
Due to untyped constants, different numeric constant can be mixed and thus this is valid:
```golang
var p = 5.2/3
```