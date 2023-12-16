package main

import "fmt"

func main(){
    fmt.Println("Maps City")
    userInfo := map[string]string{
        "huxn": "Full Stack dev",
        "john": "Frontend dev",
        "alex": "Backend dev",
    }
    userInfo["Kumar"] = "Mobile dev"

    delete(userInfo, "huxn")
    fmt.Println(userInfo)
    randNums := map[int]int{
        0: 10,
        1: 20,
        2: 30,
        3: 40,
    }
    fmt.Println(randNums)
}
