package main
import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Hello World")
	presentTime:= time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04 Monday"))
	createdDate := time.Date(2000, time.January, 20, 23, 10, 10, 39, time.UTC)
	fmt.Println(createdDate)
}