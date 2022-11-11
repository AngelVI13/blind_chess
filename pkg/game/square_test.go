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

func FuzzNewSquareFromIndex(f *testing.F) {
	f.Add(63, "h8")
	f.Add(5, "f1")

	f.Fuzz(func(t *testing.T, idx int, expNotation string) {
		square, _ := NewSquareFromIndex(idx)
		if notation := square.Notation(); notation != expNotation {
			t.Errorf(
				"wrong Square{notation=%s} from notation(%s)",
				notation,
				expNotation,
			)
		}
	})
}

func FuzzNewSquareFromIndexError(f *testing.F) {
	f.Add(64)
	f.Add(-1)

	f.Fuzz(func(t *testing.T, idx int) {
		_, err := NewSquareFromIndex(idx)
		if err == nil {
			t.Errorf("no error for NewSquareFromIndex(%d)", idx)
		}
	})
}

func FuzzSquareColorWhite(f *testing.F) {
	f.Add("b1")
	f.Add("a8")
	f.Add("c2")
	f.Add("d3")
	f.Add("h1")
	f.Add("d5")

	f.Fuzz(func(t *testing.T, notation string) {
		square, _ := NewSquareFromNotation(notation)
		if color := square.Color(); color != White {
			t.Errorf(
				"wrong color %s (expected: %s) for Square{notation=%s}",
				color,
				White,
				notation,
			)
		}
	})
}

func FuzzSquareColorBlack(f *testing.F) {
	f.Add("a1")
	f.Add("h8")
	f.Add("c1")
	f.Add("d2")
	f.Add("f6")
	f.Add("e5")

	f.Fuzz(func(t *testing.T, notation string) {
		square, _ := NewSquareFromNotation(notation)
		if color := square.Color(); color != Black {
			t.Errorf(
				"wrong color %s (expected: %s) for Square{notation=%s}",
				color,
				Black,
				notation,
			)
		}
	})
}
