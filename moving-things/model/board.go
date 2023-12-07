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
	copyState := make([][]string, len(board.InitState))
	for i := range board.InitState {
		copyState[i] = make([]string, len(board.InitState[i]))
		copy(copyState[i], board.InitState[i])
	}

	*board.Area = copyState
}

// Mencetak posisi player
func (board Board) DrawPlayer(player Player) {
	(*board.Area)[player.Position.Y][player.Position.X] = player.Head
}
