package item

import "github.com/ziruiproject/go-terminal-game/tree/main/random-rpg/entity"

type Weapon struct {
	Item       entity.Item
	Damage     int
	Durability int
}

func (weapon *Weapon) WoodenSword() *Weapon {
	return &Weapon{
		Item: entity.Item{
			Name:   "Wooden Sword",
			Amount: 1,
			Rarity: 1,
		},
		Damage:     2,
		Durability: 15,
	}
}
