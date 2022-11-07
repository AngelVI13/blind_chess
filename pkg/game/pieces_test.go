package game

var out3 string = `
import { Knight, Bishop, King, Rook, Queen } from "./pieces"
import { Square } from "./square"
import { Board } from "./board"
 
// ----------------- KNIGHT -----------------
test("Knight moves (from a1)", () => {
    const square = Square.fromNotation("a1");
    const board = new Board();
    const knight = new Knight(square, board);

    expect(knight.square).toEqual(square);

    const moves = knight.getMoves();
    expect(moves.length).toBe(2);
    expect(Square.fromNotation("b3").isContained(moves)).toBeTruthy();
    expect(Square.fromNotation("c2").isContained(moves)).toBeTruthy();
});

test("Knight moves (from d4)", () => {
    const square = Square.fromNotation("d4");
    const board = new Board();
    const knight = new Knight(square, board);

    expect(knight.square).toEqual(square);

    const moves = knight.getMoves();
    expect(moves.length).toBe(8);

    const expectedMoves: string[] = ["b3", "b5", "c6", "c2", "e2", "e6", "f5", "f3"];
    for (const expectedMove of expectedMoves) {
        expect(Square.fromNotation(expectedMove).isContained(moves)).toBeTruthy();
    }
});

test("Knight moves (from d4) with blocker (on c2)", () => {
    const board = new Board();

    const knightSquare = Square.fromNotation("d4");
    const knight = new Knight(knightSquare, board);
    
    const bishopSquare = Square.fromNotation("c2");
    const bishop = new Bishop(bishopSquare, board);

    // add blocking piece to the board
    board.addPiece(bishop);

    const moves = knight.getMoves();
    expect(moves.length).toBe(7);

    // expect blocking piece square is no longer in moves array
    expect(bishop.square.isContained(moves)).toBeFalsy();

    const expectedMoves: string[] = ["b3", "b5", "c6", "e2", "e6", "f5", "f3"];
    for (const expectedMove of expectedMoves) {
        expect(Square.fromNotation(expectedMove).isContained(moves)).toBeTruthy();
    }
});

// ----------------- KING -----------------
test("King moves (from a1)", () => {
    const square = Square.fromNotation("a1");
    const board = new Board();
    const king = new King(square, board);

    expect(king.square).toEqual(square);

    const moves = king.getMoves();

    expect(moves.length).toBe(3);
    const expectedMoves: string[] = ["a2", "b2", "b1"];
    for (const expectedMove of expectedMoves) {
        expect(Square.fromNotation(expectedMove).isContained(moves)).toBeTruthy();
    }
});

test("King moves (from d4)", () => {
    const square = Square.fromNotation("d4");
    const board = new Board();
    const king = new King(square, board);

    expect(king.square).toEqual(square);

    const moves = king.getMoves();
    expect(moves.length).toBe(8);

    const expectedMoves: string[] = ["c3", "c4", "c5", "c3", "d5", "e3", "e4", "e5"];
    for (const expectedMove of expectedMoves) {
        expect(Square.fromNotation(expectedMove).isContained(moves)).toBeTruthy();
    }
});

test("King moves (from d4) with blocker (on c5)", () => {
    const board = new Board();

    const kingSquare = Square.fromNotation("d4");
    const king = new King(kingSquare, board);

    const bishopSquare = Square.fromNotation("c5");
    const bishop = new Bishop(bishopSquare, board);

    // add blocking piece to the board
    board.addPiece(bishop);

    const moves = king.getMoves();
    expect(moves.length).toBe(7);

    // expect blocking piece square is no longer in moves array
    expect(bishop.square.isContained(moves)).toBeFalsy();

    const expectedMoves: string[] = ["c3", "c4", "d3", "d5", "e3", "e4", "e5"];
    for (const expectedMove of expectedMoves) {
        expect(Square.fromNotation(expectedMove).isContained(moves)).toBeTruthy();
    }
});

// ----------------- BISHOP -----------------
test("Bishop moves (from a1)", () => {
    const square = Square.fromNotation("a1");
    const board = new Board();
    const bishop = new Bishop(square, board);

    expect(bishop.square).toEqual(square);

    const moves = bishop.getMoves();

    expect(moves.length).toBe(7);
    const expectedMoves: string[] = ["b2", "c3", "d4", "e5", "f6", "g7", "h8"];
    for (const expectedMove of expectedMoves) {
        expect(Square.fromNotation(expectedMove).isContained(moves)).toBeTruthy();
    }
});

test("Bishop moves (from d4)", () => {
    const square = Square.fromNotation("d4");
    const board = new Board();
    const bishop = new Bishop(square, board);

    expect(bishop.square).toEqual(square);

    const moves = bishop.getMoves();
    expect(moves.length).toBe(13);

    const expectedMoves: string[] = [
        "c3", "b2", "a1", // bottom left ray
        "e5", "f6", "g7", "h8", // top right ray
        "c5", "b6", "a7",  // top left ray
        "e3", "f2", "g1" // bottom right ray
    ];

    for (const expectedMove of expectedMoves) {
        expect(Square.fromNotation(expectedMove).isContained(moves)).toBeTruthy();
    }
});

test("Bishop moves (from d4) with blocker (on e5)", () => {
    const board = new Board();

    const bishopSquare = Square.fromNotation("d4");
    const bishop = new Bishop(bishopSquare, board);

    const knightSquare = Square.fromNotation("e5");
    const knight = new Knight(knightSquare, board);

    // add blocking piece to the board
    board.addPiece(knight);

    const moves = bishop.getMoves();
    expect(moves.length).toBe(9);

    // expect blocking piece square is no longer in moves array
    expect(knight.square.isContained(moves)).toBeFalsy();

    const expectedMoves: string[] = [
        "c3", "b2", "a1", // bottom left ray
        "c5", "b6", "a7",  // top left ray
        "e3", "f2", "g1" // bottom right ray
    ];
    for (const expectedMove of expectedMoves) {
        expect(Square.fromNotation(expectedMove).isContained(moves)).toBeTruthy();
    }
});

// ----------------- QUEEN -----------------
test("Queen moves (from a1)", () => {
    const square = Square.fromNotation("a1");
    const board = new Board();
    const queen = new Queen(square, board);

    expect(queen.square).toEqual(square);

    const moves = queen.getMoves();

    expect(moves.length).toBe(21);
    const expectedMoves: string[] = [
        "a2", "a3", "a4", "a5", "a6", "a7", "a8", // north vector 
        "b2", "c3", "d4", "e5", "f6", "g7", "h8", // north-east diagonal
        "b1", "c1", "d1", "e1", "f1", "g1", "h1", // east vector
    ];
    for (const expectedMove of expectedMoves) {
        expect(Square.fromNotation(expectedMove).isContained(moves)).toBeTruthy();
    }
});

test("Queen moves (from d4)", () => {
    const square = Square.fromNotation("d4");
    const board = new Board();
    const queen = new Queen(square, board);

    expect(queen.square).toEqual(square);

    const moves = queen.getMoves();
    expect(moves.length).toBe(27);

    const expectedMoves: string[] = [
        "a1", "b2", "c3", "e5", "f6", "g7", "h8", // north-east diagonal
        "a7", "b6", "c5", "e3", "f2", "g1", // north-west diagonal
        "d8", "d7", "d6", "d5", "d3", "d2", "d1", // north-south vector
        "a4", "b4", "c4", "e4", "f4", "g4", "h4", // west-east vector
    ];

    for (const expectedMove of expectedMoves) {
        expect(Square.fromNotation(expectedMove).isContained(moves)).toBeTruthy();
    }
});

test("Queen moves (from d4) with blockers (on c4 & e3)", () => {
    const board = new Board();

    const queenSquare = Square.fromNotation("d4");
    const queen = new Queen(queenSquare, board);

    const knightSquare1 = Square.fromNotation("c4");
    const knight1 = new Knight(knightSquare1, board);

    const knightSquare2 = Square.fromNotation("e3");
    const knight2 = new Knight(knightSquare2, board);

    // add blocking pieces to the board
    board.addPiece(knight1);
    board.addPiece(knight2);

    const moves = queen.getMoves();
    expect(moves.length).toBe(21);

    // expect blocking piece squares are no longer in move list
    expect(knight1.square.isContained(moves)).toBeFalsy();
    expect(knight2.square.isContained(moves)).toBeFalsy();

    const expectedMoves: string[] = [
        "a1", "b2", "c3", "e5", "f6", "g7", "h8", // north-east diagonal
        "a7", "b6", "c5",  // north-west diagonal
        "d8", "d7", "d6", "d5", "d3", "d2", "d1", // north-south vector
        "e4", "f4", "g4", "h4", // west-east vector
    ];
    for (const expectedMove of expectedMoves) {
        expect(Square.fromNotation(expectedMove).isContained(moves)).toBeTruthy();
    }
});

// ----------------- ROOK -----------------
test("Rook moves (from a1)", () => {
    const square = Square.fromNotation("a1");
    const board = new Board();
    const rook = new Rook(square, board);

    expect(rook.square).toEqual(square);

    const moves = rook.getMoves();

    expect(moves.length).toBe(14);
    const expectedMoves: string[] = [
        "a2", "a3", "a4", "a5", "a6", "a7", "a8", // north vector 
        "b1", "c1", "d1", "e1", "f1", "g1", "h1", // east vector
    ];
    for (const expectedMove of expectedMoves) {
        expect(Square.fromNotation(expectedMove).isContained(moves)).toBeTruthy();
    }
});

test("Rook moves (from d4)", () => {
    const square = Square.fromNotation("d4");
    const board = new Board();
    const rook = new Rook(square, board);

    expect(rook.square).toEqual(square);

    const moves = rook.getMoves();
    expect(moves.length).toBe(14);

    const expectedMoves: string[] = [
        "d8", "d7", "d6", "d5", "d3", "d2", "d1", // north-south vector
        "a4", "b4", "c4", "e4", "f4", "g4", "h4", // west-east vector
    ];

    for (const expectedMove of expectedMoves) {
        expect(Square.fromNotation(expectedMove).isContained(moves)).toBeTruthy();
    }
});

test("Rook moves (from d4) with blockers (on f4)", () => {
    const board = new Board();

    const rookSquare = Square.fromNotation("d4");
    const rook = new Rook(rookSquare, board);

    const knightSquare = Square.fromNotation("f4");
    const knight = new Knight(knightSquare, board);

    // add blocking pieces to the board
    board.addPiece(knight);

    const moves = rook.getMoves();
    expect(moves.length).toBe(11);

    // expect blocking piece squares are no longer in move list
    expect(knight.square.isContained(moves)).toBeFalsy();

    const expectedMoves: string[] = [
        "d8", "d7", "d6", "d5", "d3", "d2", "d1", // north-south vector
        "a4", "b4", "c4", "e4", // west-east vector
    ];
    for (const expectedMove of expectedMoves) {
        expect(Square.fromNotation(expectedMove).isContained(moves)).toBeTruthy();
    }
});
`
