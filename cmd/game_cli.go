package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/AngelVI13/blind_chess/pkg/game"
)

func readAnswer(numAvailableOptions int) (int, error) {
	var answer string

	_, err := fmt.Scanln(&answer)
	if err != nil {
		log.Fatal(err)
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
	pieces := game.BoardPieces()
	for idx, p := range pieces {
		possibleAnswers += fmt.Sprintf("%d. %s", idx, p.Type())

		if idx < len(pieces)-1 {
			possibleAnswers += ", "
		}
	}
	question = fmt.Sprintf("%s (%s):", question, possibleAnswers)
	log.Println(question, questionPiece.Type())
}

func main() {
	g := game.New()

	g.SetupPreGame()

	log.Println("Starting position")
	for _, p := range g.BoardPieces() {
		log.Printf("%s at %s", p.Type(), p.Square().Notation())
	}

	g.StartGame()

	for {
		questionPiece, questionSquare := g.QuestionPieceAndSquare()

		printQuestion(g, questionPiece, questionSquare)
		answer, err := readAnswer(len(g.BoardPieces()))
		if err != nil {
			log.Println(err.Error())
			continue
		}

		answerPiece := g.BoardPieces()[answer]
		if answerPiece.Type() != questionPiece.Type() {
			log.Printf("Game over! (correct piece was %s)", questionPiece.Type())
			break
		}

		levelUp := g.SetNextPosition()
		if levelUp {
			levelUpPiece := g.LevelUpPiece()
			log.Printf(
				"Level up! A new %s was added to %s",
				levelUpPiece.Type(),
				levelUpPiece.Square().Notation(),
			)
		}
	}
}
