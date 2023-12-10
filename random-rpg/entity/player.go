package entity

import (
	"fmt"

	"github.com/ziruiproject/go-terminal-game/tree/main/random-rpg/entity/item"
)

type Equipment struct {
	MainHand *item.Weapon
	Offhand  *item.Weapon
}

type Player struct {
	Entity    Entity
	InvSlot   int
	Inv       [10]*item.Item
	Equipment Equipment
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
	player.Equipment.MainHand = &item
}

func (player *Player) EquipToOffhand(item item.Weapon) {
	player.Equipment.Offhand = &item
}
