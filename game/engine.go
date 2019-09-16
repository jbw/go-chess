package game

import (
	"jbw.cc/auto-chess/board"
)

type engine struct{}

// Status represents the game status
func (engine) Status(position *board.BoardState) {

}

func getPossibleMove(position *board.BoardState) board.Bitboard {
	return board.Bitboard(0)
}

func calcMoves(position *board.BoardState) {}

// FileAndRankAttack returns attack positions
func FileAndRankAttack(occupied board.Bitboard, position board.Bitboard) board.Bitboard {

	// TODO add Position(), Rank(), File() to Square
	rankMask := board.Bitboard(0)
	fileMask := board.Bitboard(0)

	rank := LinearAttack(occupied, position, rankMask)
	file := LinearAttack(occupied, position, fileMask)

	return rank | file

}

// LinearAttack implements o^(o-2r) where o is occupied, r is position
// https://www.chessprogramming.org/Subtracting_a_Rook_from_a_Blocking_Piece
func LinearAttack(occupied board.Bitboard, position board.Bitboard, mask board.Bitboard) board.Bitboard {

	potentialBlockers := occupied & mask
	difference := potentialBlockers - 2*position
	changed := difference ^ occupied
	northAttacks := changed & mask

	return northAttacks
}
