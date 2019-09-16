package board

const (
	numberOfRowsInBoard    = 8
	numberOfColumnsInBoard = 8
	numberOfSquaresInBoard = 64
)

// Board represents a chess board. A Bitboard is used to represent it.
// Square and Piece types are used here to present the relationship between them
// both.
// A board has squares, and squares have pieces on them. Squares also
// have a color (e.g. Black and White)
// and pieces have a type (e.g. Pawn, Rook etc) and a color (e.g. Black and White).
type Board struct {
	Name string
}

// NewBoard represents a chess board with pieces on it.
// Creates bitboard for each piece.
func NewBoard(m map[Square]Piece) *Board {

	board := &Board{}

	//bitboard := newBitboard()

	return board
}

// Update the board based on the given move.
// Updates the underlying bitboard representation
func (b *Board) Update(move *Move) {
	// Get bitboards for source and destination
	// Remove what is at destination
	// Move source to destination

}

// Draw the board for debugging
func (b *Board) Draw() string {
	return "Not implemented!"
}
