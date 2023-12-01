package main

import (
	"fmt"
)

type Player struct{
    health int
}

func takeDamageFromExplosion(player *Player){
    fmt.Println("Player is taking damage from explosion\n")
    explosionDmg := 10
    player.health -= explosionDmg
}


func main(){
    player := &Player{
        health: 100,
    }
    //it references to the pointer
    fmt.Printf("Before explosion %+v\n", player)
    takeDamageFromExplosion(player)
    fmt.Printf("After explosion %+v", player)
}
