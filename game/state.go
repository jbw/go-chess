package game

import (
	"jbw.cc/auto-chess/board"
)

func initalizeBoard() *board.Board {

	m := map[board.Square]board.Piece{}
	m[board.Square(0)] = board.Piece(0)
	// Create Squares

	return board.NewBoard(m)
}
