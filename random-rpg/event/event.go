package event

import (
	"fmt"
	"time"
)

func tell(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(delay * time.Millisecond)
	}
	fmt.Println()
}

func Prologue() {
	text := "In the hamlet, you were but a listless lad, regarded by the townsfolk as a ne'er-do-well, a vagabond without a craft or trade to call your own."
	tell(text, 35)
}
