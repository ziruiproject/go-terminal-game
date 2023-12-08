package entity

import (
	"fmt"
)

type Player struct {
	entity Entity
	Inv    map[string]int
}

func (player Player) SayHello() {
	fmt.Println("Hello, my name is " + player.entity.Name)
}
