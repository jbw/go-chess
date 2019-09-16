package game

import (
	"jbw.cc/auto-chess/board"
)

// Game struct
type Game struct {
	moves             []*board.Move
	boardStateHistory []*board.BoardState
	boardState        *board.BoardState
}

// NewGame creates an default game with standard
// opening position.
func NewGame() *Game {

	newBoard := initalizeBoard()
	boardState := &board.BoardState{
		Board: newBoard,
	}

	game := &Game{
		moves:             []*board.Move{},
		boardState:        boardState,
		boardStateHistory: []*board.BoardState{boardState},
	}

	return game
}

// MoveEvent represents a user move
func (game *Game) MoveEvent(move *board.Move) {
	// Check move is valid

	// Update position
	game.boardState = game.boardState.Update(move)

	// Record game events
	game.recordEvents(move)

	// Update game status
	game.checkStatus()
}

func (game *Game) recordEvents(move *board.Move) {
	game.moves = append(game.moves, move)
	game.boardStateHistory = append(game.boardStateHistory, game.boardState)
}

// checkStatus validates the game state
func (game *Game) checkStatus() {
	// Check status of the game for conclusion else
	// do nothing
}
