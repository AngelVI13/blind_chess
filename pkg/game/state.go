package game

import (
	"math/rand"
)

type State int

const (
	PreGame State = iota
	Countdown
	Play
	Between
	LevelUp
	GameOver
	Win
)

var (
	// Levels for every level add the corresponding piece to the board
	Levels   = []PieceType{Bishop, Knight, Rook, King, Queen}
	MaxLevel = len(Levels)
)

const QuestionsPerLevel = 10

type Game struct {
	board     *Board
	currState State
	level     int
	Score     int

	questionSquare         *Square
	pieceForQuestionSquare Piece

	LevelUpPiece Piece
}

func New() *Game {
	return &Game{
		board:                  NewBoard(),
		currState:              PreGame,
		level:                  0,
		Score:                  0,
		questionSquare:         nil,
		pieceForQuestionSquare: nil,
		LevelUpPiece:           nil,
	}
}

func (g *Game) BoardPieces() []Piece {
	return g.board.pieces
}

func (g *Game) PieceTypes() []PieceType {
	var pieceTypes []PieceType
	var seen = map[PieceType]struct{}{}

	for _, p := range g.BoardPieces() {
		pieceType := p.Type()

		if _, ok := seen[pieceType]; ok {
			continue
		}

		seen[pieceType] = struct{}{}
		pieceTypes = append(pieceTypes, pieceType)
	}

	return pieceTypes
}

func (g *Game) CheckAnswer(piece PieceType) bool {
	return piece == g.pieceForQuestionSquare.Type()
}

func (g *Game) Level() int {
	return g.level + 1 // level is used as index -> return real level number
}

// SetupPreGame Reset board and set 2 initial pieces
func (g *Game) SetupPreGame() {
	g.currState = PreGame
	g.level = 0
	g.Score = 0
	g.questionSquare = nil
	g.pieceForQuestionSquare = nil
	g.LevelUpPiece = nil

	g.board.Reset()

	// Compute initial piece positions
	idx1 := generateSquareIndex()
	idx2 := generateSquareIndex()

	for idx1 == idx2 {
		idx2 = generateSquareIndex()
	}

	knightSquare, _ := NewSquareFromIndex(idx1)
	g.board.AddPiece(Knight, knightSquare)

	bishopSquare, _ := NewSquareFromIndex(idx2)
	g.board.AddPiece(Bishop, bishopSquare)
}

// chooseSquareAndPiece Chooses a singular square and the piece that can reach it.
func (g *Game) chooseSquareAndPiece() {
	squares := g.board.SingularSquares()

	g.questionSquare = squares[rand.Intn(len(squares))]
	g.pieceForQuestionSquare = g.board.PieceThatReachesSquare(g.questionSquare)
}

func (g *Game) StartGame() {
	g.chooseSquareAndPiece()
}

func (g *Game) QuestionPieceAndSquare() (Piece, *Square) {
	return g.pieceForQuestionSquare, g.questionSquare
}

// SetNextPosition Generates the next position of the board by moving the chosen piece.
// Return true if level up is hit otherwise, false.
func (g *Game) SetNextPosition() bool {
	g.board.MovePiece(g.pieceForQuestionSquare, g.questionSquare)

	levelUp, _ := g.updateScore()
	g.chooseSquareAndPiece()

	return levelUp
}

// updateScore Updates the score after a correct answer and levels up if necessary.
func (g *Game) updateScore() (levelUp, win bool) {
	g.Score += 1

	// the +1 here is needed because the score starts from 0
	if g.Score == ((MaxLevel * QuestionsPerLevel) + 1) {
		// game over - the player won
		win = true
		return levelUp, win
	}

	if g.Score%QuestionsPerLevel == 0 {
		var idx int
		var sq *Square

		for {
			idx = generateSquareIndex()
			sq, _ = NewSquareFromIndex(idx)
			if !g.board.Occupied(sq) {
				break
			}
		}

		newPiece := Levels[g.level]
		g.board.AddPiece(newPiece, sq)
		g.LevelUpPiece = g.board.pieces[len(g.board.pieces)-1]

		g.level++
		levelUp = true
	}
	return levelUp, win
}

func generateSquareIndex() int {
	return rand.Intn(FileNum * RankNum)
}
