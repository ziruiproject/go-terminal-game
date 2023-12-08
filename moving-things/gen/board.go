package gen

import (
	"math/rand"
	"terminal-games/moving-things/model"
	"time"
)

func GenerateMap(width, height, obstacleCount int) *model.Board {
	rand.Seed(time.Now().UnixNano())

	// Board kosong
	area := make([][]string, height)
	for i := range area {
		area[i] = make([]string, width)
		for j := range area[i] {
			area[i][j] = "*"
		}
	}

	// Random obstacle
	for count := 0; count < obstacleCount; count++ {
		x := rand.Intn(width)
		y := rand.Intn(height)
		if area[y][x] == "*" {
			area[y][x] = "#"
		} else {
			count-- // if sudah terisi, ulang
		}
	}

	// Copy board awal
	initState := CopyArea(area)

	return &model.Board{
		Area:      &area,
		InitState: initState,
	}
}

// strong copy
func CopyArea(area [][]string) [][]string {
	copyArea := make([][]string, len(area))
	for i := range area {
		copyArea[i] = make([]string, len(area[i]))
		copy(copyArea[i], area[i])
	}
	return copyArea
}
