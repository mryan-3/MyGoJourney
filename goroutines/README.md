# Goroutines in Go

Goroutines are lightweight threads managed by the Go runtime. They allow you to execute functions concurrently in the same address space.

## Creating a Goroutine

To create a goroutine, simply call a function prefixed with the `go` keyword:

```go
go myFunction()
```

This will execute `myFunction()` asynchronously in its own goroutine.

## Communication via Channels

Goroutines can communicate with each other via channels. Channels act as pipes allowing you to send and receive values between goroutines.

Here is an example of using a channel:

```go
ch := make(chan int)

go func() {
    ch <- 42
}()

i := <-ch
fmt.Println(i)
```

This sends the integer 42 through the channel `ch`, and then receives it in the `i` variable.

## sync.WaitGroup

The `sync.WaitGroup` type allows you to wait for a collection of goroutines to finish executing.

Here is an example:

```go 
var wg sync.WaitGroup

for _, salutation := range salutations {
    wg.Add(1)

    go func(salutation string) {
        defer wg.Done()
        fmt.Println(salutation)
    }(salutation)
}

wg.Wait()
```

This concurrently prints a number of salutations, waiting for all goroutines to finish before continuing.

## Goroutine Best Practices

Here are some tips for working effectively with goroutines:

- Don't leak goroutines. Make sure to know how many you are starting and track when they exit.
- Use `sync.WaitGroup` instead of `time.Sleep()` to wait for goroutines to finish.
- Use worker pools instead of unlimited goroutines. 
- Use buffered channels to prevent deadlocks.

## Full Goroutine Example

Here is a full example demonstrating goroutines, channels and wait groups:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func processor(wg *sync.WaitGroup) {
    fmt.Println("Starting processor goroutine")

    time.Sleep(2 * time.Second)
    fmt.Println("Finished processor goroutine")

    wg.Done()
}

func main() {

    var wg sync.WaitGroup

    wg.Add(1)

    go processor(&wg)

    fmt.Println("Starting goroutines")

    wg.Wait()

    fmt.Println("All goroutines complete")
}
```

This shows asynchronously running a `processor` goroutine, waiting for it to finish using `sync.WaitGroup` before exiting the program.

Let me know if any part of goroutines and concurrency in Go needs more explanation!
