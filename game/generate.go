package game

import (
	"fmt"

	"jbw.cc/auto-chess/board"
)

// KingMovesBitboard represents 64 bitboards squares where
// the King can move to from a given square on an empty board.
func KingMovesBitboard() {

	boardColumns := 8

	// Start at this position to traverse all the
	// positions on the board for the King piece (B,7).
	// A B C D E F G H
	// 1 1 1 0 0 0 0 0
	// 1 0 1 0 0 0 0 0
	// 1 1 1 0 0 0 0 0
	// 0 0 0 0 0 0 0 0
	// 0 0 0 0 0 0 0 0
	// 0 0 0 0 0 0 0 0
	// 0 0 0 0 0 0 0 0
	// 0 0 0 0 0 0 0 0
	bitboard := board.NewBitboardWithValue(16186183351374184448)
	fmt.Println(bitboard.Draw())

	for row := 1; row < 5; row++ {
		for column := 1; column < boardColumns-2; column++ {

			bitboard.NorthOne()
			fmt.Println(bitboard.Draw())
		}

		bitboard.WestOne()
		fmt.Println(bitboard.Draw())

		for column := 1; column < boardColumns-2; column++ {

			bitboard.SouthOne()
			fmt.Println(bitboard.Draw())

		}

		bitboard.WestOne()
		fmt.Println(bitboard.Draw())
	}
}
