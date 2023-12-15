package main

import "fmt"

func main(){
    fmt.Println("Switch Statement Jones")

    role := "guest"

    switch role{
        case "guest":
            fmt.Println("Guest User")
        case "moderator":
            fmt.Println("Moderator User")
        default:
            fmt.Println("Unknown User")

    }
}
