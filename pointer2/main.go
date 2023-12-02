package main

import "fmt"

//Define a struct with two integer fields
type Point struct{
    X, Y int
}

func swapPoints(p1, p2 *Point){
    temp := *p1
    *p1 = *p2
    *p2 = temp
}

func main(){
    fmt.Println("A simple program to swap the values of two structs")
    point1 := Point{X: 40, Y: 50}
    point2 := Point{X: 70, Y: 100}
    fmt.Println("Before swap: X: %d, Y: %d", point1.X, point1.Y)
    fmt.Println("Before swap: X: %d ,Y: %d", point2.X, point2.Y)
    swapPoints(&point1, &point2)
    fmt.Println("the value of the struct 1 after swapping is as follows, X:%d, Y:%d", point1.X, point1.Y)
    fmt.Println("the value of the struct 2 after swapping is as follows: X = %d, Y - %d", point2.X, point2.Y)


}
