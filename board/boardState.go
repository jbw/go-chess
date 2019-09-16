package board

// BoardState represents the state of the game
type BoardState struct {
	Board *Board
}

// Update returns a new position from a given move.
func (position *BoardState) Update(move *Move) *BoardState {

	// Update the board with new move
	position.Board.Update(move)

	return &BoardState{
		Board: position.Board,
	}
}
