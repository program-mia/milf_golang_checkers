package main

import (
	"bufio"
	"checkers/models"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	board := models.SetupBoard()
	isWhitesMove := true

	for {
		if isWhitesMove {
			fmt.Println("White's turn\n")
		} else {
			fmt.Println("Black's turn\n")
		}

		board.Print()

		fmt.Println("Available moves:")

		if isWhitesMove {
			models.PrintAvailableMoves(board.GetAvailableMovesForWhite())
		} else {
			models.PrintAvailableMoves(board.GetAvailableMovesForBlack())
		}

		move := readMoveFromInput(reader)

		fmt.Println(move)
		fmt.Println("\n")

		wasMoveMade := board.MovePiece(move, isWhitesMove)

		if !wasMoveMade {
			fmt.Println("Failed to move, please try again")

			continue
		}

		isWhitesMove = !isWhitesMove
	}
}

func readMoveFromInput(reader *bufio.Reader) models.Move {
	var readerError error
	var from_x, from_y, to_x, to_y int
	result := ""

	for result == "" {
		fmt.Print("[from] [to]: ")

		result, readerError = reader.ReadString('\n')

		if readerError != nil {
			fmt.Println("There was an error when processing your input, please try again.")
		}

		result = strings.Replace(result, "\n", "", -1)

		fields := strings.Split(result, " ")

		if len(fields) != 2 || len(fields[0]) != 2 || len(fields[1]) != 2 {
			result = ""

			fmt.Println("Invalid argument, please use the correct format of [A5] [B5].")

			continue
		}

		from_x, from_y = models.GetBoardCoordinatesFromStringNotation(fields[0])
		to_x, to_y = models.GetBoardCoordinatesFromStringNotation(fields[1])

		if from_y == -1 || from_x == -1 || from_y >= 8 || to_x == -1 || to_y == -1 || to_y >= 8 {
			result = ""

			fmt.Println("Invalid argument, please use the correct format of [A5] [B5].")
		}
	}

	return models.CreateMove(from_x, from_y, to_x, to_y)
}
