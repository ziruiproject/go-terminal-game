package game

import (
	"fmt"
	"strings"
	"terminal-games/model"
)

type Game struct {
	Board  model.Board
	Player model.Player
}

func (game Game) Start() {
	for {
		var userInput string
		fmt.Print("Enter Moves: ")
		fmt.Scanln(&userInput)

		switch strings.ToUpper(userInput) {
		case "UP":
			game.changeState(game.Player.Up)
		case "DOWN":
			game.changeState(game.Player.Down)
		case "LEFT":
			game.changeState(game.Player.Left)
		case "RIGHT":
			game.changeState(game.Player.Right)
		}
	}
}

func (game *Game) changeState(changes func()) {
	changes()
	game.Board.Reset()
	game.Board.DrawPlayer(game.Player)
	game.Board.Show()
}

func (game *Game) Init() {
	game.Board.DrawPlayer(game.Player)
	game.Board.Show()
}
