package game

var out1 string = `
import { Board } from "./board"
import { Bishop, Knight } from "./pieces"
import { Square } from "./square"

test("Board isOccupied", () => {
    // NOTE: this also checks Square.isContained functionality
    const board = new Board();

    const square = Square.fromNotation("a1");
    const bishop = new Bishop(square, board);

    board.addPiece(bishop);
    
    // Iterate over all possible square indices and check that
    // none of them are occupied (except for the index on which the bishop stands)
    for (let i = 0; i < 64; i++ ) {
        if (i !== square.index()) {
            expect(board.isOccupied(Square.fromIndex(i))).toBeFalsy();
        } else {
            expect(board.isOccupied(Square.fromIndex(i))).toBeTruthy();
        }
    }
});

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
