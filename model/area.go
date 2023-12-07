package model

import "fmt"

type Board struct {
	Area      *[][]string
	InitState [][]string
}

func (board Board) Show() {
	for _, v := range *board.Area {
		fmt.Println(v)
	}
	fmt.Println()
}

func (board *Board) Reset() {
	*board.Area = board.InitState
}

// Mencetak posisi player
func (board Board) DrawPlayer(player Player) {
	(*board.Area)[player.Position.Y][player.Position.X] = "^"
}
