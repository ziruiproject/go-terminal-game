package entity

import (
	"fmt"

	"github.com/ziruiproject/go-terminal-game/tree/main/random-rpg/entity/item"
)

type Player struct {
	Mob     Mob
	InvSlot int
	Inv     [10]*item.Item
}

func (player *Player) AddToInventory(item item.Item) {
	if player.InvSlot < len(player.Inv) {
		player.Inv[player.InvSlot] = &item
		player.InvSlot++
	} else {
		fmt.Println("Inventory Full!")
	}
}

func (player *Player) ShowInventory() {
	for _, v := range player.Inv {
		fmt.Print(v)
	}
}

func (player *Player) EquipToHand(item item.Weapon) {
	player.Mob.Equipment.MainHand = &item
}

func (player *Player) EquipToOffhand(item item.Weapon) {
	player.Mob.Equipment.Offhand = &item
}
