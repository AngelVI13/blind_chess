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

	KnightDir = []DirectionVec{
		{file: -1, rank: -2}, // TL
		{file: -2, rank: -1}, // TL2
		{file: -2, rank: 1},  // BL
		{file: -1, rank: 2},  // BL2
		{file: 1, rank: -2},  // TR
		{file: 2, rank: -1},  // TR2
		{file: 2, rank: 1},   // BR
		{file: 1, rank: 2},   // BR2
	}
)

type PieceType string

const (
	Bishop PieceType = "Bishop"
	Knight           = "Knight"
	Rook             = "Rook"
	King             = "King"
	Queen            = "Queen"
)

var PieceTypes = map[PieceType]struct{}{
	Bishop: {},
	Knight: {},
	Rook:   {},
	King:   {},
	Queen:  {},
}

type pieceProperties struct {
	board      *Board
	square     *Square
	directions []DirectionVec
	PieceType
}

func (p pieceProperties) Type() PieceType {
	return p.PieceType
}

func (p pieceProperties) Square() *Square {
	return p.square
}

func (p *pieceProperties) SetSquare(square *Square) {
	p.square = square
}

type slidingPiece struct {
	pieceProperties
}

// Moves Get an array of squares to which the sliding piece instance can move to.
func (p *slidingPiece) Moves() []*Square {
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

// Moves Get an array of squares to which the non-sliding piece instance can move to.
func (p *nonSlidingPiece) Moves() []*Square {
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

func NewBishop(board *Board, square *Square) *slidingPiece {
	return &slidingPiece{
		pieceProperties{
			directions: Diagonal,
			PieceType:  Bishop,
			board:      board,
			square:     square,
		},
	}
}

func NewRook(board *Board, square *Square) *slidingPiece {
	return &slidingPiece{
		pieceProperties{
			directions: Orthogonal,
			PieceType:  Rook,
			board:      board,
			square:     square,
		},
	}
}

func NewQueen(board *Board, square *Square) *slidingPiece {
	return &slidingPiece{
		pieceProperties{
			directions: Combined,
			PieceType:  Queen,
			board:      board,
			square:     square,
		},
	}
}

func NewKing(board *Board, square *Square) *nonSlidingPiece {
	return &nonSlidingPiece{
		pieceProperties{
			directions: Combined,
			PieceType:  King,
			board:      board,
			square:     square,
		},
	}
}

func NewKnight(board *Board, square *Square) *nonSlidingPiece {
	return &nonSlidingPiece{
		pieceProperties{
			directions: KnightDir,
			PieceType:  Knight,
			board:      board,
			square:     square,
		},
	}
}
