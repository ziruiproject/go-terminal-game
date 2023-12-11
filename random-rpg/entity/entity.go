package entity

type Entity struct {
	Name    string
	Health  int
	Defense int
	Attack  int
}

func (entity *Entity) GotAttacked(enemy Mob) {

	if enemy.Equipment.MainHand != nil {
		entity.Health -= enemy.Equipment.MainHand.Damage
	}
	entity.Health = entity.Health - (enemy.Entity.Attack + entity.Defense)
}

func (entity *Entity) CheckHealth() bool {
	return entity.Health <= 0
}
