package piece

// Pawn struct
type Pawn struct {
	position Position
}

// NewPawn represents the creation of a new Pawn piece
func NewPawn(x int, y int) *Pawn {
	return &Pawn{
		position: Position{X: x, Y: y},
	}
}

// IsMoveValid represents a valid move check
func (pawn *Pawn) IsMoveValid(p2 Position) bool {

	p1 := pawn.position

	// legal move check
	if p2.X != p1.X {
		return false
	}

	if (p2.Y - p1.Y) != 1 {
		return false
	}

	return true
}
