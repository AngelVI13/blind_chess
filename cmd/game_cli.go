package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/AngelVI13/blind_chess/pkg/game"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func Score(g *game.Game) string {
	return fmt.Sprintf(
		"Level %d Score %d/%d Total %d\n",
		g.Level(),
		g.Score%game.QuestionsPerLevel,
		game.QuestionsPerLevel,
		g.Score,
	)
}

func readAnswer(numAvailableOptions int) (int, error) {
	var answer string

	_, err := fmt.Scanln(&answer)
	if err != nil {
		return -1, err
	}

	answerChoice, err := strconv.Atoi(answer)
	if err != nil {
		return -1, fmt.Errorf(
			"Answer should be an number corresponding to the piece from available options",
		)
	}

	// NOTE: answer choice here is an index
	if answerChoice < 0 || answerChoice > numAvailableOptions-1 {
		return -1, fmt.Errorf(
			"Answer should be a number between 0 and %d (inclusive)", numAvailableOptions-1,
		)
	}

	return answerChoice, nil
}

func printQuestion(game *game.Game, questionPiece game.Piece, questionSquare *game.Square) {
	question := fmt.Sprintf("Which piece can go to %s", questionSquare.Notation())
	possibleAnswers := ""
	pieces := game.PieceTypes()
	for idx, p := range pieces {
		possibleAnswers += fmt.Sprintf("%d. %s", idx, p)

		if idx < len(pieces)-1 {
			possibleAnswers += ", "
		}
	}
	question = fmt.Sprintf("%s (%s):", question, possibleAnswers)
	fmt.Println(question)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	g := game.New()

	g.SetupPreGame()

	clearScreen()
	fmt.Println("Starting position")
	for _, p := range g.BoardPieces() {
		fmt.Printf("%s at %s\n", p.Type(), p.Square().Notation())
	}

	fmt.Println("Game starts in 5 seconds")
	time.Sleep(5 * time.Second)
	clearScreen()

	g.StartGame()

	for {
		questionPiece, questionSquare := g.QuestionPieceAndSquare()

		printQuestion(g, questionPiece, questionSquare)
		answer, err := readAnswer(len(g.BoardPieces()))
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		answerPiece := g.BoardPieces()[answer]
		if answerPiece.Type() != questionPiece.Type() {
			fmt.Printf("Game over! (correct piece was %s)\n", questionPiece.Type())
			fmt.Printf("%s", Score(g))
			break
		}

		clearScreen()

		levelUp := g.SetNextPosition()
		fmt.Printf("Success! %s", Score(g))

		if levelUp {
			fmt.Printf(
				"Level up! A new %s was added to %s\n",
				g.LevelUpPiece.Type(),
				g.LevelUpPiece.Square().Notation(),
			)
		}
	}
}
