package board

// A Square is one of the 64 rank and file combinations
// that make up a chess board.
type Square int8

// A Rank is the rank of a square.
type Rank int8

// A File is the file of a square.
type File int8

func getSquare(f File, r Rank) Square {
	return 0
}

// File represents the squares the file
func (square Square) File() File {
	return File(0)
}

// Rank represents the squares rank
func (square Square) Rank() Rank {
	return Rank(0)
}

// Position represent the squares position bitboard
func (square Square) Position() Bitboard {
	return Bitboard(0)
}
