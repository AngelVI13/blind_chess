package game

import (
	"testing"
)

func FuzzNewSquare(f *testing.F) {
	f.Add(8, 8)
	f.Add(-1, -1)
	f.Add(-1, 5)
	f.Add(5, -1)

	f.Fuzz(func(t *testing.T, file, rank int) {
		_, err := NewSquare(file, rank)
		if err == nil {
			t.Errorf("no error for NewSquare(file=%d, rank=%d)", file, rank)
		}
	})
}

func FuzzSquareIndex(f *testing.F) {
	f.Add(1, 1, 9)
	f.Add(1, 2, 17)
	f.Add(2, 1, 10)
	f.Add(7, 7, 63)

	f.Fuzz(func(t *testing.T, file, rank, expIdx int) {
		square, _ := NewSquare(file, rank)
		if idx := square.Index(); idx != expIdx {
			t.Errorf(
				"wrong square index(expected %d != actual %d) for Square{file=%d, rank=%d}",
				expIdx,
				idx,
				file,
				rank,
			)
		}
	})
}

func FuzzSquareNotation(f *testing.F) {
	f.Add(0, 0, "a1")
	f.Add(1, 6, "b7")
	f.Add(7, 7, "h8")

	f.Fuzz(func(t *testing.T, file, rank int, expNotation string) {
		square, _ := NewSquare(file, rank)
		if notation := square.Notation(); notation != expNotation {
			t.Errorf(
				"wrong square notation(expected %s != actual %s) for Square{file=%d, rank=%d}",
				expNotation,
				notation,
				file,
				rank,
			)
		}
	})
}

func FuzzNewSquareFromNotation(f *testing.F) {
	f.Add("a1", 0, 0)
	f.Add("d4", 3, 3)
	f.Add("h7", 7, 6)

	f.Fuzz(func(t *testing.T, notation string, expFile, expRank int) {
		square, _ := NewSquareFromNotation(notation)
		if square.file != expFile || square.rank != expRank {
			t.Errorf(
				"wrong Square{file=%d, rank=%d} from notation(%s)",
				expFile,
				expRank,
				notation,
			)
		}
	})
}

func FuzzNewSquareFromNotationError(f *testing.F) {
	f.Add("a12")
	f.Add("z1")
	f.Add("bb")
	f.Add("1a")

	f.Fuzz(func(t *testing.T, notation string) {
		_, err := NewSquareFromNotation(notation)
		if err == nil {
			t.Errorf("no error for NewSquareFromNotation(%s)", notation)
		}
	})
}

var out5 string = `
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
