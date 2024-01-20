package main

import (
	"fmt"
)


func main(){
    chanel := make(chan string, 2)
    chanel <- "First Message"
    chanel <- "Second message"
    fmt.Println(<- chanel)
    fmt.Println(<- chanel)



}






















//time.sleep serves two purposes:
/*To simulate work and execution time in the goroutines. In real code, the goroutines would be doing actual work that takes time. But in these simple examples, time.Sleep gives an artificial delay to simulate the passage of time and work being done in the goroutine.
To allow goroutines time to finish before main exits. Goroutines run concurrently alongside the main thread. But when main exits, all goroutines are forced to stop. So the time.Sleep calls pause main from exiting early while goroutines are still working.
*/
