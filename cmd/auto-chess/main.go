package main

import (
	"fmt"

	"jbw.cc/auto-chess/board"
	"jbw.cc/auto-chess/game"
)

func main() {

	game.KingMovesBitboard()

	game.NewGame()

	// Turn into test case
	/*
		00110010
		10110111
		01000100
		00000000
		10000000
		01100010
		00011111
		00010010
	*/
	o := board.NewBitboardWithValue(3654464391579049746)
	fmt.Println(o.Draw())

	/*
		10000000
		00000000
		00000000
		00000000
		00000000
		00000000
		00010000
		00000000
	*/
	r := board.NewBitboardWithValue(9223372036854779904)
	fmt.Println(r.Draw())

	/*
		00010000
		00010000
		00010000
		00010000
		00010000
		00010000
		00010000
		00010000
	*/
	m := board.NewBitboardWithValue(1157442765409226768)
	fmt.Println(m.Draw())

	bb := game.LinearAttack(o, r, m)

	fmt.Println(bb.Draw())

	fmt.Println(game.FileAndRankAttack(o, r).Draw())

}
