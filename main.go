package main

import (
	"terminal-games/model"
)

func main() {

	var player model.Player = model.Player{
		Position: model.Position{
			X: 2,
			Y: 4,
		},
	}

	var area [][]string = [][]string{
		{"#", "#", "#", "*", "*"},
		{"*", "#", "#", "#", "*"},
		{"*", "#", "*", "#", "*"},
		{"#", "#", "*", "*", "*"},
		{"#", "*", "*", "#", "*"},
	}

	var board model.Board = model.Board{
		Area: &area,
	}

	board.DrawPlayer(player)
	board.Show()
	player.Move(1, -1)
	board.DrawPlayer(player)
	board.Show()
}
