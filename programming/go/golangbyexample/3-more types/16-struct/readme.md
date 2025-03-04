# Intro

Struct is a set of variables which can be of different types.

Struct acts as a container that has heterogeneous data types which together represents an entity.

```golang
type Employee struct {
    name   string
    age    int
    salary int
}
```
> heterogeneous - diverse in character or content

# Creating a struct variable

```golang
JohnDoe := Employee{}
```

In this case all fields will be initialized with zero values.

You can initialize it this way
```golang
emp := employee{name: "Sam", age: 31, salary: 2000}
```

or you can omit some of fields and let them take its types zero values
```golang
emp := employee{name: "Sam", salary: 2000}
```

Also you can omit field names, but in such case all values for each of the field has to be provided in sequence

```golang
emp := employee{"Sam", 0, 2000}
```

# Pointer to a struct
Like with pointers in general there are two ways to create pointers to a struct variable:
- using `new` 
- using `&`

### using `&`
```golang
emp := employee{name: "Sam", age: 31, salary: 2000}
empP := &emp
```
struct pointer can also be directly created as well
```golang
empP := &employee{name: "Sam", age: 31, salary: 2000}
```

The `employee{name: "Sam", age: 31, salary: 2000}` alone creates a variable with given values under some memory location <br>
`&` gets address of this variable

### using `new`

```golang
emp := new(employee)
emp.name = "Sam"; emp.age = 31; emp.salary=2000
```

The `new` will reserve a memory for variable of type employee and give you the address of it to variable `emp`.

The fulfillment of this variable with field you have to do manually by `.` operator

## Printing a variable

### fmt package
```golang
emp := employee{"Sam", 31, 2000}
fmt.Printf("%v", emp) // will print only values of fields
fmt.Printf("%+v", emp) // will print field names also
fmt.Printf("%#v", emp) // will print field names also
```
### econding/json package
Also you can use `encoding/json` package


```golang
package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type employee struct {
    Name   string
    Age    int
    salary int
}

func main() {
    emp := employee{Name: "Sam", Age: 31, salary: 2000}
    //Marshal
    empJSON, err := json.Marshal(emp)
    if err != nil {
        log.Fatalf(err.Error())
    }
    fmt.Printf("Marshal funnction output %s\n", string(empJSON))

    //MarshalIndent
    empJSON, err = json.MarshalIndent(emp, "", "  ")
    if err != nil {
        log.Fatalf(err.Error())
    }
    fmt.Printf("MarshalIndent funnction output %s\n", string(empJSON))
}
```

# Metadata for fields
A struct allows to add metadata to its fields. Any meta information can be stored with fields and can be used by any package for different purposes

e.g. `encoding/json` package utilizes such metadata tags

```golang
type employee struct {
    Name   string `json:"name"`
    Age    int    `json:"age"`
    Salary int    `json:"salary"`
}
```

# Anonymous Fields in struct

Anonymous field is one with no name, the type will become field name

e.g.
```golang
type employee struct {
    string
    age    int
    salary int
}
```

The anonymous field can also be accessed and assigned a value
```golang
package main

import "fmt"

type employee struct {
    string
    age    int
    salary int
}

func main() {
    emp := employee{string: "Sam", age: 31, salary: 2000}
    //Accessing a struct field
    n := emp.string
    fmt.Printf("Current name is: %s\n", n)
    //Assigning a new value
    emp.string = "John"
    fmt.Printf("New name is: %s\n", emp.string)
}
```

But I think the most cases it is used are like this:

```golang
type Employee struct {
	string
	int
	float32
}

func main() {
	emp := Employee{"Sam", 31, 2000.0}
	fmt.Println(emp)
}
```
But the types cannot be repeated, that is why I had to change last field type to `float32`.

# Nested struct
Of course type of a filed in struct can be of any type especially user defined ones especially struct.

And of course you can combine nesting with anonymity. 

# Exported and UnExported field of a struct

With upper/lower-case you can export/unexport everything/anything in go.

The scope of visibility is the same as in everywhere else:
- lower-case - visible only in the same package
- upper-case - visible everywhere

# Struct equality

The first thing to know before considering struct equality is weather if all struct fields types are comparable or not:

Which go types are comparable and which are not?
Comparable
- boolean
- numeric (int, float)
- string
- pointer
- channel
- interface types
- struct (if all field types are comparable)
- array (if the type of values of array are comparable)
Not-comparable
- slice
- map
- func

So two struct will be equal if first all their field types are comparable and all the corresponding field values are equal. Let’s see an example
```golang
package main

import "fmt"

type employee struct {
    name   string
    age    int
    salary int
}

func main() {
    emp1 := employee{name: "Sam", age: 31, salary: 2000}
    emp2 := employee{name: "Sam", age: 31, salary: 2000}
    if emp1 == emp2 {
        fmt.Println("emp1 annd emp2 are equal")
    } else {
        fmt.Println("emp1 annd emp2 are not equal")
    }
}
```

If struct contains field of not-comparable type and you will use `==` on it, the compiler will raise an error

# Structs are value types
A struct is value type in go (the other type is reference type). <br>
It means that struct variable name is not a pointer to the struct. In fact it denotes the entire struct.

So...

A new copy of struct will be created when
- a struct variable is assigned to another struct variable
- a struct variable is passed as an argument to a function (which is the 1-st case actually)

