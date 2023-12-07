package model

import "fmt"

type Board struct {
	Area *[][]string
}

func (board Board) Show() {
	for _, v := range *board.Area {
		fmt.Println(v)
	}
}
