package game

import (
	"fmt"
	"strings"
	"terminal-games/model"
)

type Game struct {
	Board  *model.Board
	Player model.Player
}

func (game Game) Start() {
	for {
		var userInput string
		fmt.Print("Enter Moves: ")
		fmt.Scanln(&userInput)

		switch strings.ToUpper(userInput) {
		case "W":
			game.changeState(game.Player.Up)
		case "S":
			game.changeState(game.Player.Down)
		case "A":
			game.changeState(game.Player.Left)
		case "D":
			game.changeState(game.Player.Right)
		default:
			break
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
