package game

type DirectionVec struct {
	file int
	rank int
}

var (
	Orthogonal = []DirectionVec{
		{file: 0, rank: -1}, // North
		{file: 0, rank: 1},  // South
		{file: -1, rank: 0}, // West
		{file: 1, rank: 0},  // East
	}

	Diagonal = []DirectionVec{
		{file: -1, rank: -1}, // Topleft
		{file: -1, rank: 1},  // Topright
		{file: 1, rank: -1},  // Bottomleft
		{file: 1, rank: 1},   // Bottomright
	}

	Combined = append(
		Orthogonal,
		Diagonal...,
	)

	Knight = []DirectionVec{
		{file: -1, rank: -2}, // TL
		{file: -2, rank: -1}, // TL2
		{file: -2, rank: 1},  // BL
		{file: -1, rank: 2},  // BL2
		{file: 1, rank: -2},  // TR
		{file: 2, rank: -1},  // TR2
		{file: 2, rank: 1},   // BR
		{file: 1, rank: 2},   // BR2
	}

	PieceAbbreviations = [...]string{"B", "R", "Q", "K", "N"}
)

type Piece interface {
	GetMoves() []Square
}

type pieceProperties struct {
	board        *Board
	square       *Square
	directions   []DirectionVec
	abbreviation string
}

type slidingPiece struct {
	pieceProperties
}

// GetMoves Get an array of squares to which the sliding piece instance can move to.
func (p *slidingPiece) GetMoves() []*Square {
	var moves []*Square

	var newSquare *Square
	for _, direction := range p.directions {
		newSquare = p.square

		for {
			newFile := newSquare.file + direction.file
			newRank := newSquare.rank + direction.rank

			var err error
			newSquare, err = NewSquare(newFile, newRank)
			if err != nil {
				// we are out of bounds -> this square doesn't exist
				// move to next direction
				break
			}

			// if square exists but its occupied -> move to next direction
			if p.board.Occupied(newSquare) {
				break
			}

			moves = append(moves, newSquare)
		}
	}

	return moves
}

type nonSlidingPiece struct {
	pieceProperties
}

// GetMoves Get an array of squares to which the non-sliding piece instance can move to.
func (p *nonSlidingPiece) GetMoves() []*Square {
	var moves []*Square

	for _, direction := range p.directions {
		newFile := p.square.file + direction.file
		newRank := p.square.rank + direction.rank

		newSquare, err := NewSquare(newFile, newRank)
		if err != nil {
			// we are out of bounds -> this square doesn't exist
			// move to next direction
			continue
		}

		// if square exists but its occupied -> move to next direction
		if p.board.Occupied(newSquare) {
			continue
		}

		moves = append(moves, newSquare)
	}

	return moves
}

var out2 string = `

export class Bishop extends SlidingPiece {
    directions = DIAGONAL;
    abbreviation = "B";
}

export class Rook extends SlidingPiece {
    directions = ORTHOGONAL;
    abbreviation = "R";
}

export class Queen extends SlidingPiece {
    directions = COMBINED;
    abbreviation = "Q";
}

export class King extends NonSlidingPiece {
    directions = COMBINED;
    abbreviation = "K";
}

export class Knight extends NonSlidingPiece {
    directions = KNIGHT;
    abbreviation = "N";
} 
`
