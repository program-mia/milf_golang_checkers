package models

import (
	"fmt"
	"strconv"
)

type Pawn struct {
	is_white bool
	is_queen bool
}

type Square struct {
	pawn *Pawn
}

type Board struct {
	board [8][8]*Square
}

type Move struct {
	from_x int
	from_y int
	to_x   int
	to_y   int
}

func SetupBoard() *Board {
	var board Board

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			var square Square

			if (y == 0 || y == 1 || y == 2) && (x+y)%2 == 1 {
				square.pawn = &Pawn{
					is_white: false,
					is_queen: false,
				}
			}

			if (y == 6 || y == 7 || y == 5) && (x+(y-4))%2 == 1 {
				square.pawn = &Pawn{
					is_white: true,
					is_queen: false,
				}
			}

			board.board[y][x] = &square
		}
	}

	return &board
}

func (board *Board) Print() {
	for y := 0; y < 8; y++ {
		fmt.Print(y + 1)
		fmt.Print(". ")

		for x := 0; x < 8; x++ {
			squareInterior := "_"

			if board.board[y][x].pawn != nil {
				if board.board[y][x].pawn.is_white {
					squareInterior = "W"
				} else {
					squareInterior = "B"
				}
			}

			fmt.Print("[" + squareInterior + "]")
		}

		fmt.Println("")
	}

	fmt.Print("   ")

	for i := 0; i < 8; i++ {
		fmt.Print(" ")
		fmt.Print(GetIndexAsLetter(i))
		fmt.Print(" ")
	}

	fmt.Println("")
}

func GetIndexAsLetter(index int) string {
	return [8]string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
	}[index]
}

func GetIndexFromLetter(letter string) int {
	letters := [8]string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
	}

	for i, v := range letters {
		if letter == v {
			return i
		}
	}

	return -1
}

func GetBoardCoordinatesFromStringNotation(coordinates string) (int, int) {
	if len(coordinates) != 2 {
		return -1, -1
	}

	x := GetIndexFromLetter((string)(coordinates[0]))
	y, _ := strconv.Atoi((string)(coordinates[1]))

	y = y - 1

	return x, y
}

func (board *Board) GetAvailableMovesForWhite() map[int]Move {
	return board.getAvailableMovesFor(true)
}

func (board *Board) GetAvailableMovesForBlack() map[int]Move {
	return board.getAvailableMovesFor(false)
}

func (board *Board) getAvailableMovesFor(isWhite bool) map[int]Move {
	availableMoves := make(map[int]Move)

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			pawn := board.board[y][x].pawn

			if pawn == nil {
				continue
			}

			if pawn.is_white != isWhite {
				continue
			}

			possibleMoves := [2]int{-1, 1}

			allDirections := [2]int{-1, 1}
			var directions []int

			if pawn.is_queen {
				directions = allDirections
			} else if pawn.is_white {
				directions = allDirections[0:1]
			} else {
				directions = allDirections[1:2]
			}

			for _, val := range possibleMoves {
				for _, direction := range directions {
					if !board.doesFieldExist(x+val, y+direction) {
						continue
					}

					if board.isFieldEmpty(x+val, y+direction) {
						availableMoves[len(availableMoves)] = Move{
							from_x: x,
							from_y: y,
							to_x:   x + val,
							to_y:   y + direction,
						}

						continue
					}

					if !board.doesFieldHasPawnOfGivenColor(x+val, y+direction, !isWhite) {
						continue
					}

					if !board.doesFieldExist(x+(val*2), y+(direction*2)) {
						continue
					}

					if board.isFieldEmpty(x+(val*2), y+(direction*2)) {
						availableMoves[len(availableMoves)] = Move{
							from_x: x,
							from_y: y,
							to_x:   x + (val * 2),
							to_y:   y + (direction * 2),
						}
					}
				}
			}
		}
	}

	return availableMoves
}

func (board *Board) doesFieldExist(x int, y int) bool {
	return x >= 0 && x < 8 && y >= 0 && y < 8
}

func (board *Board) isFieldEmpty(x int, y int) bool {
	square := board.board[y][x]

	return square.pawn == nil
}

func (board *Board) doesFieldHasPawnOfGivenColor(x int, y int, isWhite bool) bool {
	square := board.board[y][x]

	if square.pawn == nil {
		return false
	}

	return square.pawn.is_white == isWhite
}

func PrintAvailableMoves(moves map[int]Move) {
	for index, val := range moves {
		fmt.Printf(
			"%d: %s%d - %s%d\n",
			index+1,
			GetIndexAsLetter(val.from_x),
			val.from_y+1,
			GetIndexAsLetter(val.to_x),
			val.to_y+1,
		)
	}
}

func (board *Board) MovePiece(move Move, isWhiteMove bool) bool {
	canMove := false
	var arrayOfMoves map[int]Move

	if isWhiteMove {
		arrayOfMoves = board.GetAvailableMovesForWhite()
	} else {
		arrayOfMoves = board.GetAvailableMovesForBlack()
	}

	canMove = isMoveInArray(&arrayOfMoves, &move)

	if !canMove {
		return false
	}

	board.board[move.to_y][move.to_x].pawn = board.board[move.from_y][move.from_x].pawn
	board.board[move.from_y][move.from_x].pawn = nil

	var xDiff int = move.from_x - move.to_x

	if xDiff < 0 {
		xDiff *= -1
	}

	if xDiff == 2 {
		board.board[(move.from_y+move.to_y)/2][(move.from_x+move.to_x)/2].pawn = nil
	}

	return true
}

func isMoveInArray(moves *map[int]Move, move *Move) bool {
	for _, val := range *moves {
		if val == *move {
			return true
		}
	}

	return false
}

func CreateMove(from_x, from_y, to_x, to_y int) Move {
	return Move{
		from_x: from_x,
		from_y: from_y,
		to_x:   to_x,
		to_y:   to_y,
	}
}
