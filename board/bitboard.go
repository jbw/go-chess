package board

import (
	"strconv"
	"strings"
)

// Bitboard presents a chess board encoded in an unsigned 64-bit integer.
// Reference: https://www.chessprogramming.org/Efficient_Generation_of_Sliding_Piece_Attacks
type Bitboard uint64

// NewBitboard represents an emptty chess board
func NewBitboard() Bitboard {

	s := ""

	for square := 0; square < numberOfSquaresInBoard; square++ {
		s += "0"
	}

	bitboard, err := strconv.ParseUint(s, 2, 64)

	if err != nil {
		panic(err)
	}

	return Bitboard(bitboard)

}

// NewBitboardWithValue represents an emptty chess board
func NewBitboardWithValue(bitboard uint64) Bitboard {
	return Bitboard(bitboard)
}

// ToString returns a 64 character strings representing
// the board
func (bb Bitboard) ToString() string {
	s := strconv.FormatUint(uint64(bb), 2)
	return strings.Repeat("0", numberOfSquaresInBoard-len(s)) + s
}

// Draw prints a human readable bitboard
func (bb Bitboard) Draw() string {

	bbStr := "A B C D E F G H\n"
	bbStrSplits := strings.Split(bb.ToString(), "")

	for square := 0; square < len(bbStrSplits); square++ {
		if (square+1)%numberOfRowsInBoard == 0 {
			bbStr += bbStrSplits[square] + " \n"
		} else {
			bbStr += bbStrSplits[square] + " "
		}
	}

	return bbStr
}
