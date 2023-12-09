package main

import (
	"fmt"

	"github.com/ziruiproject/go-terminal-game/tree/main/random-rpg/entity"
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

	player.Inv[0] = entity.Item{
		Name:   "Wooden Sword",
		Amount: 1,
		Rarity: 1,
	}
	fmt.Println(player.Inv[0])
}
