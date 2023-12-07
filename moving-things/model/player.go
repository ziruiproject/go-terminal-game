package model

type Position struct {
	X int
	Y int
}

type Player struct {
	Position Position
	Head     string
	Board    *Board
}

func (player *Player) move(xPos int, yPos int) {
	newX := player.Position.X + xPos
	newY := player.Position.Y + yPos

	// Apakah # && keluar map?
	if newX >= 0 && newX < len((*player.Board.Area)[0]) &&
		newY >= 0 && newY < len(*player.Board.Area) &&
		(*player.Board.Area)[newY][newX] != "#" {
		player.Position.X = newX
		player.Position.Y = newY
	}
}

func (player *Player) Up() {
	player.move(0, -1)
	player.Head = "^"
}

func (player *Player) Down() {
	player.move(0, 1)
	player.Head = "v"
}

func (player *Player) Left() {
	player.move(-1, 0)
	player.Head = "<"
}

func (player *Player) Right() {
	player.move(1, 0)
	player.Head = ">"
}
