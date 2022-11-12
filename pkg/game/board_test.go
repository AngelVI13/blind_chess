package game

import (
	"log"
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
		log.Print(sq, board)

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

var out1 string = `

test("Board getSingularSquares", () => {
    const board = new Board();

    const squareBishop = Square.fromNotation("a1");
    const bishop = new Bishop(squareBishop, board);

    const squareKnight = Square.fromNotation("b1");
    const knight = new Knight(squareKnight, board);

    board.addPiece(bishop);
    board.addPiece(knight);

    const moves = board.getSingularSquares()
    expect(moves.length).toBe(8);

    // expect square which can be reached by both pieces to not be in the array of moves
    expect(Square.fromNotation("c3").isContained(moves)).toBeFalsy();

    const expectedMoves: string[] = ["b2", "d4", "e5", "f6", "g7", "h8", "a3", "d2"];
    for (const expectedMove of expectedMoves) {
        expect(Square.fromNotation(expectedMove).isContained(moves)).toBeTruthy();
    }
});

test("Board getPieceThatReachesSquare", () => {
    const board = new Board();

    const squareBishop = Square.fromNotation("a1");
    const bishop = new Bishop(squareBishop, board);

    const squareKnight = Square.fromNotation("b1");
    const knight = new Knight(squareKnight, board);

    board.addPiece(bishop);
    board.addPiece(knight);

    expect(board.getPieceThatReachesSquare(Square.fromNotation("b2"))).toBe(bishop);
    expect(board.getPieceThatReachesSquare(Square.fromNotation("a3"))).toBe(knight);
    expect(board.getPieceThatReachesSquare(Square.fromNotation("d8"))).toBeNull();
    expect(board.getPieceThatReachesSquare(Square.fromNotation("e2"))).toBeNull();
});
`
