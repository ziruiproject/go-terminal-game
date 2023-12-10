package main

import (
	"github.com/ziruiproject/go-terminal-game/tree/main/random-rpg/entity"
	"github.com/ziruiproject/go-terminal-game/tree/main/random-rpg/entity/item"
)

func main() {
	var player entity.Player = entity.Player{
		Entity: entity.Entity{
			Name:    "Zirui",
			Health:  100,
			Defense: 10,
			Attack:  7,
		},
	}

	woodenSword := item.WoodenSword()
	player.AddToInventory(woodenSword.Item)
	player.ShowInventory()
	player.EquipToHand(*woodenSword)
}
