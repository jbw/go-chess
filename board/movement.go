package board

// Compass Rose
//  noWe         nort         noEa
//          +7    +8    +9
//              \  |  /
//  west    -1 <-  0 -> +1    east
//              /  |  \
//          -9    -8    -7
//  soWe         sout         soEa

// NorthOne adds a row to the top of the bitboard
// shifting the board downwards
func (bb *Bitboard) NorthOne() *Bitboard {
	*bb = *bb >> uint(8)
	return bb
}

// SouthOne adds a row to the bottom of the bitboard
// shifting the board upwards
func (bb *Bitboard) SouthOne() *Bitboard {
	*bb = *bb << uint(8)
	return bb
}

// EastOne adds a column to the right of the bitboard
// shifting the board to the left
func (bb *Bitboard) EastOne() *Bitboard {
	*bb = (*bb << uint(1)) & notHFile
	return bb
}

// WestOne adds a column to the left of the bitboard
// shifting the board to the right
func (bb *Bitboard) WestOne() *Bitboard {
	*bb = (*bb >> uint(1)) & notAFile
	return bb
}
