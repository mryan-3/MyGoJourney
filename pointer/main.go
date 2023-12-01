package main

import (
	"fmt"
)

func swap(x, y *int){
    temp := *x
    //Initialize a temporary variable and assign it the value pointed by *x
    *x = *y
    *y = temp
    //The value pointed to by x is set to the value pointed to by y.
    //The value pointed to by y is set to the temporary value (temp), which o    riginally held the value pointed to by x.
}

func main(){
    fmt.Println("Simple Swappning Program to pracice on pointers")
    //Initialize two variables
    a := 5
    b := 10

    fmt.Printf("Before Swap a: %d, b: %d\n", a, b)

    //Create two pointers to the variables
    ptrA := &a
    ptrB := &b

    //Call the function with pointers
    swap(ptrA, ptrB)

    fmt.Printf("After swap - a: %d, b: %d\n", a, b)

}
