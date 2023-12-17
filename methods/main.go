package main

import "fmt"


type Human struct{
    gender string
    fullName string
    age int
    Friends
}
type Friends struct{
    justFriends int
    bestFriend string
}
func (h Human) printInfo(){
    fmt.Printf("My name is %v, Gender: %v, Age: %v, I have %v, Friends but my best friend is %v", h.fullName, h.gender, h.age, h.justFriends, h.bestFriend)
}

func  main(){
    fmt.Println("Methods")
    person := Human{
        gender: "Male",
        fullName: "Big Joe",
        age: 20,
        Friends: Friends{
            justFriends: 3,
            bestFriend: "Ricky Maivia",
        },
    }

    person.printInfo()

}
