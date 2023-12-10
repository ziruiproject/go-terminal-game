package item

type Weapon struct {
	Item       Item
	Damage     int
	Durability int
}

func WoodenSword() *Weapon {
	return &Weapon{
		Item: Item{
			Name:   "Wooden Sword",
			Amount: 1,
			Rarity: 1,
		},
		Damage:     2,
		Durability: 15,
	}
}
