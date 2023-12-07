package main

import (
	"terminal-games/game"
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

	// Create a deep copy of the area
	initState := make([][]string, len(area))
	for i := range area {
		initState[i] = make([]string, len(area[i]))
		copy(initState[i], area[i])
	}

	var board model.Board = model.Board{
		Area:      &area,
		InitState: initState,
	}

	var game game.Game = game.Game{
		Board:  board,
		Player: player,
	}

	game.Init()
	game.Start()
}
