package entity

import "github.com/ziruiproject/go-terminal-game/tree/main/random-rpg/entity/item"

type Mob struct {
	Entity    Entity
	Equipment Equipment
}

type Equipment struct {
	MainHand *item.Weapon
	Offhand  *item.Weapon
}
