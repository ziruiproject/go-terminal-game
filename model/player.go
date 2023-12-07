package model

type Position struct {
	X int
	Y int
}

type Player struct {
	Position Position
}

func (player Player) Move(xPos int, yPos int) {
	player.Position.X += xPos
	player.Position.Y += yPos
}

func (player Player) SetPlayer(area [][]string) {
	area[player.Position.Y][player.Position.X] = "^"
}
