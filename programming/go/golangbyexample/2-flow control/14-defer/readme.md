# Intro

defer - put off (an action or event) to a later time; postpone. (odraczać)

So you put  `defer` in a function to defer cleanup activites.This cleanup activities will be done in a different function called by defer.  This different function is called at the end of the surrounding function before it returns. Below is the syntax of defer function.

```golang
defer {function_or_method_call}
```

Note that:
- Execution of a deferred function is delayed to the moment the surrounding function returns.
- Deffered function will also be executed if the enclosing function terminates with error/abruptly. e.g. in case of panic


# The need for defer
One good example of understanding the defer function is to look at the use case of writing to a file. A file that is opened for writing also must be closed.   
```golang
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    err := writeToTempFile("Some text")
    if err != nil {
        log.Fatalf(err.Error())
    }
    fmt.Printf("Write to file succesful")
}

func writeToTempFile(text string) error {
    file, err := os.Open("temp.txt")
    if err != nil {
        return err
    }
    n, err := file.WriteString("Some text")
    if err != nil {
        return err
    }
    fmt.Printf("Number of bytes written: %d", n)
    file.Close()
    return nil
}
```
In the above program, in the `writeToTempFile` function, we are opening a file and then trying to write some content to the file. After we have written the contents of the file we close the file. It is possible that during the write operation it might result into an error and function will return without closing the file. A tego byś nie chciała.

`Defer`red function helps to avoid these kinds of problems.

`Defer`red function is always executed before the surrounding function returns. 


Let's rewrite the program with `defer`.

```golang
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    err := writeToTempFile("Some text")
    if err != nil {
        log.Fatalf(err.Error())
    }
    fmt.Printf("Write to file succesful")
}

func writeToTempFile(text string) error {
    file, err := os.Open("temp.txt")
    if err != nil {
        return err
    }
    defer file.Close()

    n, err := file.WriteString("Some text")
    if err != nil {
        return err
    }
    fmt.Printf("Number of bytes written: %d", n)
    return nil
}
```

In the above program, we do defer file.Close() after opening the file. This will make sure that closing of the file is executed even if the write to the file results into an error. Defer function makes sure that the file will be closed regardless of number of return statements in the function.

# What can you put in defer?
### Custom func (not from some package)
```golang
package main
import "fmt"
func main() {
    defer test()
    fmt.Println("Executed in main")
}
func test() {
    fmt.Println("In Defer")
}
```
Output:
```
Executed in main
In Defer
```
### Inline function
```golang
package main

import "fmt"

func main() {
    defer func() { fmt.Println("In inline defer") }()
    fmt.Println("Executed")
}
```
Output:
```
Executed
In inline defer
```


# Multiple defers in one func
```golang
package main
import "fmt"
func main() {
    i := 0
    i = 1
    defer fmt.Println(i)
    i = 2
    defer fmt.Println(i)
    i = 3
    defer fmt.Println(i)
}
```
Output:
```
3
2
1
```

# Defer function and named return values
Ok, so what happens if defer change value of some variable inside 
```golang
package main
import "fmt"
func main() {
    s := test()
    fmt.Println(s)
}
func test() (size int) {
    defer func() { size = 20 }()
    size = 30
    return
}
```
Output:
```
20
```

# Defer and panic?
Panic is panic. It is really strong and finishes the program, what about defer in this case, who is stronger?
Response: Defer.


defer function will also be executed even if panic happens in a program.  When the panic is raised in a function then the execution of that function is stopped and any deferred function will be executed. In fact deferred function of all the function calls in the stack will also be executed until all the functions have returned .At that time the program will exit and it will print the panic message

So if a  defer function is present it then it will be executed and the control will be  returned back to the caller function which will again execute its defer function if present and the chain goes on until the program exists.


```golang
package main
import "fmt"
func main() {
    defer fmt.Println("Defer in main")
    panic("Panic with Defer")
    fmt.Println("After painc in f2")
}
```

Output:
```
Defer in main
panic: Panic Create

goroutine 1 [running]:
main.main()
        /Users/slohia/go/src/github.com/golang-examples/articles/tutorial/panicRecover/deferWithPanic/main.go:7 +0x95
exit status 2
```