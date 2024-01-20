package main

import (
	"math/rand"
	"fmt"
	"time"
)


func main(){
    chanel := make(chan string)

    go throwStars(chanel)
    for mess := range chanel{

    fmt.Println(mess)
    }
}

func throwStars (chanel chan string ){
    rand.NewSource(time.Now().UnixNano())
    numRounds := 3
    for i:= 0; i< numRounds; i++{
        score := rand.Intn(10)
        chanel <- fmt.Sprintf("You scored: %d", score)
    }
    close(chanel)
}






















//time.sleep serves two purposes:
/*To simulate work and execution time in the goroutines. In real code, the goroutines would be doing actual work that takes time. But in these simple examples, time.Sleep gives an artificial delay to simulate the passage of time and work being done in the goroutine.
To allow goroutines time to finish before main exits. Goroutines run concurrently alongside the main thread. But when main exits, all goroutines are forced to stop. So the time.Sleep calls pause main from exiting early while goroutines are still working.
*/
