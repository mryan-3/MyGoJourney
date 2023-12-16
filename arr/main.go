package main

import (
	"fmt"
//	"strings"
)

func main(){
    /*
    arr1 := [3]string{"John", "Jane", "Joe"}
    for i := 0; i < len(arr1); i++{
        fmt.Println(arr1[i])
        fmt.Printf("%q\n", strings.Split(arr1[i], ""))
    }
    fmt.Println(len(arr1))
    */
    fmt.Println("Arrays")
    arr1 := [5]int { 1, 3 , 65, 6456, 53}
    fmt.Println(arr1)
    arr2 := [3]string{ "Dragon", "Swiss", "Mox"  }
    fmt.Println(arr2)

}
