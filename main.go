package main

import (
	"chess/board"
	"fmt"
)

func main() {
	chessboard := board.InitDefaultBoard()

	fmt.Println(chessboard)
}
