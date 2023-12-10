package main

import (
	"fmt"
	"strconv"

	"github.com/ziruiproject/go-terminal-game/tree/main/random-rpg/entity"
)

func main() {

	player := entity.Player{
		Mob: entity.Mob{
			Entity: entity.Entity{
				Name:    "Zirui",
				Health:  100,
				Defense: 5,
				Attack:  4,
			},
			Equipment: entity.Equipment{},
		},
	}

	enemy := entity.Mob{
		Entity: entity.Entity{
			Name:    "Goblin",
			Health:  25,
			Defense: 2,
			Attack:  3,
		},
		Equipment: entity.Equipment{},
	}

	for {
		enemy.Entity.GotAttacked(player.Mob)
		player.Mob.Entity.GotAttacked(enemy)
		fmt.Println(player.Mob.Entity.Name + ": " + strconv.Itoa(player.Mob.Entity.Health))
		fmt.Println(enemy.Entity.Name + ": " + strconv.Itoa(enemy.Entity.Health))
		if enemy.Entity.CheckHealth() || player.Mob.Entity.CheckHealth() {
			fmt.Println("He's Dead")
			break
		}
	}

}
