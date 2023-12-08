package main

import "github.com/ziruiproject/go-terminal-game/tree/main/random-rpg/entity"

func main() {
	var player entity.Player = entity.Player{
		Name:   "Zirui",
		Health: 100,
		Money:  25,
	}

	player.SayHello()
}
