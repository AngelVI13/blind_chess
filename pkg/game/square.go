package game

import (
	"fmt"
	"strconv"
	"strings"
)

type Color int

const (
	White Color = iota
	Black
)

var (
	Files = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	Ranks = []int{1, 2, 3, 4, 5, 6, 7, 8}
)

const (
	// 32-bit bitmask to determine the color of squares (all 1's represent black)
	ColorBitboard uint32 = 0xAA55AA55

	FileNum = 8
	RankNum = 8
)

type Square struct {
	file int
	rank int
}

func NewSquare(file, rank int) (*Square, error) {
	if file < 0 || file >= FileNum || rank < 0 || rank >= RankNum {
		return nil, fmt.Errorf("Ivalid rank/file: f-%d r-%d", file, rank)
	}
	return &Square{
		file: file,
		rank: rank,
	}, nil
}

// NewSquareFromIndex constructor for creating Square objects from 64-based index.
func NewSquareFromIndex(index int) (*Square, error) {
	if index < 0 || index >= (FileNum*RankNum) {
		return nil, fmt.Errorf("Wrong index: %d. Index should be in range [0, 63]", index)
	}

	return &Square{
		rank: index / 8,
		file: index % 8,
	}, nil
}

// NewSquareFromNotation constructor for creating Square objects
// from square notation ("a1", "d3").
func NewSquareFromNotation(notation string) (*Square, error) {
	if len(notation) != 2 {
		return nil, fmt.Errorf(
			"Wrong square notation format %s, expected format \"b7\"",
			notation,
		)
	}

	notation = strings.ToLower(notation)

	fileStr := string(notation[0])
	rankStr := string(notation[1])

	rankInt, err := strconv.Atoi(rankStr)
	if err != nil {
		return nil, fmt.Errorf(
			"Wrong square notation %s, rank is not a number",
			notation,
		)
	}

	fileIdx, fileFound := Contains(Files, fileStr)
	rankIdx, rankFound := Contains(Ranks, rankInt)

	if !fileFound || !rankFound {
		return nil, fmt.Errorf(
			"Wrong square notation %s, rank/file not found",
			notation,
		)
	}

	return &Square{
		file: fileIdx,
		rank: rankIdx,
	}, nil
}

func Contains[T comparable](s []T, e T) (int, bool) {
	for i, v := range s {
		if v == e {
			return i, true
		}
	}
	return -1, false
}

// Index 64-based index of the square
func (s *Square) Index() int {
	return (s.rank * 8) + s.file
}

// Notation Return a notation string of the square position (ex a1, d3).
func (s *Square) Notation() string {
	return fmt.Sprintf("%s%s", Files[s.file], fmt.Sprint(Ranks[s.rank]))
}

// Return the color of the square: Black or White
func (s *Square) Color() Color {
	// Convert index to 32-bit representation since the board pattern is the same
	index := s.Index()
	if index >= 32 {
		index -= 32
	}

	// Compute if square is black or white based on index
	// and its intersection to the color bitmask
	isBlack := (ColorBitboard >> index) & 1

	if isBlack != 0 {
		return Black
	}
	return White
}

/*
   // Check if this square is contained in an array of squares.
   isContained(squares: this[]): boolean {
       for (const square of squares) {
           if (square.index() === this.index()) {
               return true;
           }
       }
       return false;
   }
*/
