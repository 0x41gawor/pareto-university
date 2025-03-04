https://golangbyexample.com/packages-modules-go-first/
# Run single go file
You can just create `main.go` copy your code and run it with `go run main.go` or `go run .`. <br> 
This way you don't need to create any packages etc.

# Packages
First line of a file is **Package Declaration**

All files in directory need to be in the same package !!!
But package name has not to be the same as directory name (its common misconception)

Package can be:
- executable - only main package
- utility package

# Modules

**Module** - collection of packages with `go.mod` at its root.
`go.mod` defines the:
- module import path (the path that you will need to enter in other .go file to import the module) e.g `github.com/0x41gawor/go-crud-api-sample/model`
- version of go with which module was created
- dependency requirements for a successful build

```go
module github.com/0x41gawor/go-crud-api-sample

go 1.20

require github.com/go-sql-driver/mysql v1.7.1

require github.com/gorilla/mux v1.8.0

require golang.org/x/crypto v0.12.0

require github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
```

***Module*** is a go support for dependency management.

In addition `go.mod` also keeps a `go.sum` file which contains the crypto-hash of bits of all project dependent modules, which validates if project dependent modules are not changed. 

## Before modules
With modules go project doesn’t necessarily have to lie in the `$GOPATH/src` folder.

But before modules all go projects had to be in `$GOPATH/src` dir.


# go install

`go install` creates the binary with the name as the last part of the ***module import path*** that contains the .go file belonging to the main package.


Within the same package, all variables, functions, constants are accessible between different .go files belonging to the package.

# Import keyword

```go
import (
    "fmt"
    "sample.com/learn/math2"
)
```

Such code means to import a package present at `math2` directory. It has not to be necessarily named `math2`.

# Exported vs Unexported Names

There are 5 identifiers that can be exported/unexported:
- Structure
- Structure's Method
- Structure's Field
- Function
- Variable

**Capitalized Identifiers** are exported and this visible in other packages.

**Non-capitalized identifiers** are unexported and will only be accessed from within the same package

# Nested Packages

To do so just in directory containing package files create subdirectory and files in there begin with other ***Package declaration***
Example is `math` and `advanced`. While importing such package just include its nesting.

# Aliasing the import

```go
package main
import (
    "fmt"
    "sample.com/learn/math"
    advc "sample.com/learn/math/advanced"
)
func main() {
    fmt.Println(math.Add(2, 1))
    fmt.Println(adc.Square(2, 1))
}
```

# Init functions

`init()` is special func that is used to initialize global variables of a package.

Each of the .go (source files) can have its own `init()` func

Whenever you import any package in the program, then on the execution of that program, init functions(if present)  in the GO source files belonging to that imported package are called first.

Quick facts:
- `init()` is optional
- `init()` does not take any args
- `init()` does not return any values

`init()` function is majorly used for the initialization of global variables that cannot be initialized using an initialization expression. For example, it requires a network call to intialize any DB client. Another example could be fetching secret keys on startup.

# Order of execution of a Go program

- program starts in the main package and goes to every imported package in every source file of main package
- it recursively goes to the source files of imported packages till it will reach a package with no further imports
- then program runs `init()` function in last level packages
- recursively it goes and finally reaches the main package 
- `init()` func in main source files is run
- `main()` func in main package is run

Package initialization is only done once even if it is imported several times.

Example:<br>
Both `a` and `b` packages have two files. `main` package imports `a` and `a` imports `b`. The order will look like this:<br>

1. Init: b1
2. Init: b2
3. Init: a1
4. Init: a2
5. Init: main
6. Main Function Executing

# Blank identifier in import
In general blank identifier can replace any unused variable (which as you know go doesn't allow).
<br> So the a blank import can be used when:
- The imported package is not being used in the current program
- But we intend to import that package so that the init function in the GO source files belonging to that package can be called and initialization of variables in that package can be done properly

So basically a blank import is used when a package is solely imported for its side effects. As an example MySQL package is used as a blank import for its side-effect of registering the MySQL driver as a database driver in the init() function of MySQL package, without importing any other functions:
```go
_ "github.com/go-sql-driver/mysql"
```

# Package naming convention
It is recommended to avoid:
- Underscore in the package name
- Camel casing or any kind of mixed caps

Name should be short but clear.
