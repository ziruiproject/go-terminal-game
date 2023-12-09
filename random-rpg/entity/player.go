package entity

import (
	"fmt"
)

type Player struct {
	Entity  Entity
	Inv     [10]Item
	InvSlot int
}

func (player *Player) AddToInventory(item Item) {
	if player.InvSlot < len(player.Inv) {
		player.Inv[player.InvSlot] = item
		player.InvSlot++
	} else {
		fmt.Println("Inventory Full!")
	}
}

func (player Player) SayHello() {
	fmt.Println("Hello, my name is " + player.Entity.Name)
}
