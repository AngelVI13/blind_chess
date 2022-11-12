package game

import (
	"testing"
)

func TestBoardOccupied(t *testing.T) {
	board := NewBoard()
	square, err := NewSquareFromNotation("a1")
	if err != nil {
		t.Error(err)
	}

	board.AddPiece(Bishop, square)

	// Iterate over all possible square indices and check that
	// none of them are occupied (except for the index on which the bishop stands)
	for i := 0; i < 64; i++ {
		sq, err := NewSquareFromIndex(i)
		if err != nil {
			t.Error(err)
		}

		if i == square.Index() {
			if !board.Occupied(sq) {
				t.Errorf("square should be occupied but isn't idx=%d", i)
			}
		} else {
			if board.Occupied(sq) {
				t.Errorf("square shouldn't be occupied but it is idx=%d", i)
			}
		}
	}
}

func TestBoardMovePiece(t *testing.T) {
	board := NewBoard()

	a1, _ := NewSquareFromNotation("a1")
	board.AddPiece(Bishop, a1)

	bishop := board.pieces[0]

	d6, _ := NewSquareFromNotation("d6")
	board.MovePiece(bishop, d6)

	if bishop.Square().Notation() != d6.Notation() {
		t.Errorf(
			"expected bishop to be on %s but is at %s",
			d6.Notation(),
			bishop.Square().Notation(),
		)
	}

}

func TestBoardSingularSquares(t *testing.T) {
	board := NewBoard()

	bishopSquare, _ := NewSquareFromNotation("a1")
	board.AddPiece(Bishop, bishopSquare)

	knightSquare, _ := NewSquareFromNotation("b1")
	board.AddPiece(Knight, knightSquare)

	squares := board.SingularSquares()

	expected := []string{"b2", "d4", "e5", "f6", "g7", "h8", "a3", "d2"}

	if len(squares) != len(expected) {
		t.Errorf(
			"expected %d singular squares but got %d instead",
			len(expected),
			len(squares),
		)
	}

	for _, expNotation := range expected {
		found := false

		for _, sq := range squares {
			if sq.Notation() == expNotation {
				found = true
			}
		}

		if !found {
			t.Errorf(
				"expected %s to be a singular square but its not",
				expNotation,
			)
		}
	}
}

func FuzzBoardPieceThatReachesSquare(f *testing.F) {
	board := NewBoard()

	bishopSquare, _ := NewSquareFromNotation("a1")
	board.AddPiece(Bishop, bishopSquare)

	knightSquare, _ := NewSquareFromNotation("b1")
	board.AddPiece(Knight, knightSquare)

	f.Add("b2", string(Bishop))
	f.Add("c3", string(Bishop))
	f.Add("d4", string(Bishop))
	f.Add("h8", string(Bishop))

	f.Add("a3", string(Knight))
	f.Add("d2", string(Knight))

	f.Add("d8", "")
	f.Add("e3", "")
	f.Add("c4", "")
	f.Add("e2", "")

	f.Fuzz(func(t *testing.T, notation, expPieceType string) {
		sq, _ := NewSquareFromNotation(notation)

		piece := board.PieceThatReachesSquare(sq)

		pieceType := ""
		if piece != nil {
			pieceType = string(piece.Type())
		}

		if pieceType != expPieceType {
			t.Errorf(
				"expected %s to be reached by %s but got '%s'",
				notation,
				string(expPieceType),
				pieceType,
			)
		}
	})
}
