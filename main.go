package main

import (
	"terminal-games/game"
	"terminal-games/gen"
	"terminal-games/model"
)

func main() {
	var board *model.Board = gen.GenerateMap(5, 5, 5)

	var player model.Player = model.Player{
		Position: model.Position{
			X: 2,
			Y: 4,
		},
		Head:  "^",
		Board: board,
	}

	var game game.Game = game.Game{
		Board:  board,
		Player: player,
	}

	game.Init()
	game.Start()
}
