package entity

import (
	"fmt"
)

type Player struct {
	Entity Entity
	Inv    [10]Item
}

func (player Player) SayHello() {
	fmt.Println("Hello, my name is " + player.Entity.Name)
}
