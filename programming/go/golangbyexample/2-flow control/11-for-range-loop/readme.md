# Intro
**for-range** loop is used to iterate over:
- array/slice
- string
- map
- channel

And for each type it works a little different, to better understand the characteristics of each type.

## Array / Slice
```golang
for index, value := range array/slice {
    //Do something with index and value
}
```
Both `index` or `value` are optional.

So you can use `_` in place of it.

But you can omit them both e.g.

```golang
//Without index and value. Just print array values
    fmt.Println("\nWithout Index and Value")
    i := 0
    for range letters {
        fmt.Printf("Index: %d Value: %s\n", i, letters[i])
        i++
    }
```

## String
As you remember string is a UTF-8 sequence of bytes. Each Unicode char size can vary from 1 to 4 bytes.
e.g. string `sample "a£c"`` is 4 bytes long because:
- `a` is 1 byte
- `£` is 2 bytes
- `c` is 1 byte

Therefore when we loop over with normal for as below:

```golang
sample := "a£b"
 for i := 0; i < len(sample); i++ {
    fmt.Printf("%c\n", sample[i])
 }
```

The output will be 
```sh
aÂ£b
```
because it loops byte over byte, not unicode char over unicode char

And this is where **for-range** loop comes into place :D

```golang
    sample := "a£b"

    //With index and value
    fmt.Println("Both Index and Value")
    for i, letter := range sample {
        fmt.Printf("Start Index: %d Value:%s\n", i, string(letter))
    }
```


## Map
```golang
    sample := map[string]string{
        "a": "x",
        "b": "y",
    }

    //Iterating over all keys and values
    fmt.Println("Both Key and Value")
    for k, v := range sample {
        fmt.Printf("key :%s value: %s\n", k, v)
    }
```

## Channel
**for-range** loop works differently too for a channel. For a channel, an index doesn't make any sense as the channel is similar to a pipeline where values enter from one and exit from the other end.

So in case of channel, the for-range loop will iterate over values currently present in the channel. After it has iterated over all the values currently present (if any), the for-range loop will not exit but instead wait for next value that might be pushed to the channel and it will exit only when the channel is closed

Below is the format when using for-range with channel

```golang
package main

import "fmt"

func main() {
    ch := make(chan string)
    go pushToChannel(ch)
    for val := range ch {
        fmt.Println(val)
    }
}
func pushToChannel(ch chan<- string) {
    ch <- "a"
    ch <- "b"
    ch <- "c"
    close(ch)
}
```