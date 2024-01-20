package main

import (
	"fmt"
	"time"
)


func main(){
    now := time.Now()
    defer func (){
        fmt.Println(time.Since(now))
    }()

    smokeSignal := make(chan bool)

    evilNinja := "Tommy"

    go attack(evilNinja, smokeSignal)

    <- smokeSignal
}

func attack(target string, smokeSignal chan bool){
    time.Sleep(1 * time.Second)
    fmt.Println("Throwing stars at", target)
    smokeSignal <- true
}






















//time.sleep serves two purposes:
/*To simulate work and execution time in the goroutines. In real code, the goroutines would be doing actual work that takes time. But in these simple examples, time.Sleep gives an artificial delay to simulate the passage of time and work being done in the goroutine.
To allow goroutines time to finish before main exits. Goroutines run concurrently alongside the main thread. But when main exits, all goroutines are forced to stop. So the time.Sleep calls pause main from exiting early while goroutines are still working.
*/
