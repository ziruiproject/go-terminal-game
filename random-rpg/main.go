package main

import (
	"fmt"

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

	player.Inv[0] = item.WoodenSword().Item
	fmt.Println(player.Inv[0])
}
