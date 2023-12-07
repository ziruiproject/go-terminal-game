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

func (player *Player) Up(yPos int) {
	player.move(0, -yPos)
}

func (player *Player) Down(yPos int) {
	player.move(0, yPos)
}

func (player *Player) Left(xPos int) {
	player.move(-xPos, 0)
}

func (player *Player) Right(xPos int) {
	player.move(xPos, 0)
}
