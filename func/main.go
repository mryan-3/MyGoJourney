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
func itter(x []int){
    for i, v := range x {
        fmt.Println(i ,v)
    }
}

func main(){
    double(1, 2, 3, 45,4 , 56)
    greet("Reya", "Jones")
    sliced := []int{
        1, 2, 3 , 4,
    }
    itter(sliced)
}
