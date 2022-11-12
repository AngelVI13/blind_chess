package game

import "fmt"

type Piece interface {
	Moves() []*Square
	Square() *Square
	Type() PieceType
}

type Board struct {
	pieces []Piece
}

func NewBoard() *Board {
	return &Board{
		pieces: make([]Piece, 0, FileNum), // Max pieces possible
	}
}

// Reset resets board (removes all pieces from the board).
func (b *Board) Reset() {
	b.pieces = make([]Piece, 0, FileNum) // Max pieces possible

}

// Occupied Checks if a given square is already occupied by a piece or not.
func (b *Board) Occupied(square *Square) bool {
	for _, piece := range b.pieces {
		if piece.Square().Index() == square.Index() {
			return true
		}
	}

	return false
}

// AddPiece Add a piece to the board
func (b *Board) AddPiece(pieceType PieceType, square *Square) error {
	var piece Piece

	switch pieceType {
	case Bishop:
		piece = NewBishop(b, square)
	case Knight:
		piece = NewKnight(b, square)
	case Rook:
		piece = NewRook(b, square)
	case King:
		piece = NewKing(b, square)
	case Queen:
		piece = NewQueen(b, square)
	default:
		return fmt.Errorf("Unknow piece type %s", pieceType)
	}

	b.pieces = append(b.pieces, piece)
	return nil
}

// SingularSquares Get a slice of squares to which only 1 piece can go.
func (b *Board) SingularSquares() []*Square {
	var allMoves = map[int]*Square{}
	var duplicateSquareIdx = map[int]struct{}{}

	// Find all duplicate squares across all piece moves
	for _, piece := range b.pieces {
		for _, pieceMove := range piece.Moves() {
			sqIdx := pieceMove.Index()

			_, duplicate := allMoves[sqIdx]

			if duplicate {
				duplicateSquareIdx[sqIdx] = struct{}{}
				continue
			}

			allMoves[sqIdx] = pieceMove
		}
	}

	var singularSquares []*Square
	for _, sq := range allMoves {
		_, duplicate := duplicateSquareIdx[sq.Index()]

		if duplicate {
			continue
		}

		singularSquares = append(singularSquares, sq)
	}

	return singularSquares
}

// PieceThatReachesSquare Get piece object that can reach the given square.
func (b *Board) PieceThatReachesSquare(square *Square) Piece {
	sqIdx := square.Index()

	for _, piece := range b.pieces {
		for _, pieceMove := range piece.Moves() {
			if sqIdx == pieceMove.Index() {
				return piece
			}
		}
	}
	return nil
}
