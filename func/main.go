package main

import "fmt"


func greet(firstName, lastname string){
    fmt.Println("Hello", firstName, lastname)
}
func double(nums ...int){
    for _, v := range nums{
        fmt.Println(v * 2)
    }
}

func main(){
    double(1, 2, 3, 45,4 , 56)
    greet("Reya", "Jones")

}
