package main

import "terminal-games/model"

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

	player.SetPlayer(area)

}
