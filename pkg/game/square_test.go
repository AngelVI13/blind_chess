package game

var out5 string = `
import { Square, Color } from "./square"


test("Square constructor", () => {
    // TODO Is this the correct way to check that this object creation succeeds
    expect(new Square(2, 3)).toEqual(new Square(2, 3));

    // Test the validity check of rank & file inputs (should be between 0-7 inclusive)
    expect(() => new Square(8, 8)).toThrow('Ivalid rank/file: f-8 r-8');
    expect(() => new Square(-1, -1)).toThrow('Ivalid rank/file: f--1 r--1');
    expect(() => new Square(-1, 5)).toThrow('Ivalid rank/file: f--1 r-5');
    expect(() => new Square(5, -1)).toThrow('Ivalid rank/file: f-5 r--1');
});

test("Square index getter", () => {
    let square = new Square(1, 1);
    expect(square.index()).toBe(9);

    square = new Square(7, 7);
    expect(square.index()).toBe(63);
});

test("Square notation getter", () => {
    let square = new Square(0, 0);
    expect(square.notation()).toBe("a1");

    square = new Square(1, 6);
    expect(square.notation()).toBe("b7");

    square = new Square(7, 7);
    expect(square.notation()).toBe("h8");
});

test("Square fromNotation constructor", () => {
    let square = Square.fromNotation("a1");
    expect(square.file).toBe(0);
    expect(square.rank).toBe(0);

    square = Square.fromNotation("d4");
    expect(square.file).toBe(3);
    expect(square.rank).toBe(3);

    square = Square.fromNotation("h7");
    expect(square.file).toBe(7);
    expect(square.rank).toBe(6);

    // Test the validity check of notation input
    // input too long (> 2 chars)
    expect(() => Square.fromNotation("a12")).toThrow('Wrong square format a12');
    // input file out of bounds
    expect(() => Square.fromNotation("z1")).toThrow('Wrong square format z1');
    // input rank not a valid number in range [0, 7]
    expect(() => Square.fromNotation("aa")).toThrow('Wrong square format aa');
});

test("Square fromIndex constructor", () => {
    let square = Square.fromIndex(63);
    expect(square).toEqual(Square.fromNotation("h8"));

    square = Square.fromIndex(5);
    expect(square).toEqual(Square.fromNotation("f1"));

    // Test the validity check of index input
    // input above upper bound
    expect(() => Square.fromIndex(64)).toThrow('Wrong index: 64. Index should be in range [0, 63]');
    // input below lower bound
    expect(() => Square.fromIndex(-1)).toThrow('Wrong index: -1. Index should be in range [0, 63]');
});

test("Square color getter", () => {
    interface SquareColor {
        notation: string;
        color: Color;
    }

    let squares: SquareColor[] = [
        // black squares
        { notation: "a1", color: Color.Black },
        { notation: "h8", color: Color.Black },
        { notation: "c1", color: Color.Black },
        { notation: "d2", color: Color.Black },
        { notation: "f6", color: Color.Black },
        { notation: "e5", color: Color.Black },

        // white squares
        { notation: "b1", color: Color.White },
        { notation: "a8", color: Color.White },
        { notation: "c2", color: Color.White },
        { notation: "d3", color: Color.White },
        { notation: "h1", color: Color.White },
        { notation: "d5", color: Color.White },
    ];

    for (const square of squares) {
        let sq = Square.fromNotation(square.notation);
        expect(sq.color()).toBe(square.color);
    }
});
`
