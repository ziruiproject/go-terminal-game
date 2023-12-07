package model

type Position struct {
	X int
	Y int
}

type Player struct {
	Position Position
}

func (player *Player) move(xPos int, yPos int) {
	player.Position.X += xPos
	player.Position.Y += yPos
}

func (player *Player) Up() {
	player.move(0, -1)
}

func (player *Player) Down() {
	player.move(0, 1)
}

func (player *Player) Left() {
	player.move(-1, 0)
}

func (player *Player) Right() {
	player.move(1, 0)
}
