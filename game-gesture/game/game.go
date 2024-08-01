package game

import (
	"fmt"
	"math/rand"
	"tictactoe-mediapipe/scripts"
)

func StartGame(board *[3][3]rune, indexReference *[3][3]string) rune {
	fmt.Println("This is an empty board: ")
	PrintBoardRune(board)
	fmt.Println("To enter your choice, raise your hand when prompted.")

	var userChoices, computerChoices [3][3]int8

	for {
		// Get row input from Python script
		row, err := scripts.GetPythonGestureInput("Give row input")
		if err != nil || row < 0 || row > 2 {
			fmt.Println("Invalid row input. Please try again.")
			continue
		}

		// Get column input from Python script
		col, err := scripts.GetPythonGestureInput("Give column input")
		if err != nil || col < 0 || col > 2 {
			fmt.Println("Invalid column input. Please try again.")
			continue
		}

		if CheckIfChoiceExists(board, row, col) {
			fmt.Println("That choice has already been taken, please try again")
			continue
		}

		userChoices[row][col] = 1
		UpdateBoard(board, userChoices, computerChoices)
		PrintBoardRune(board)

		if IsGameOver(board) {
			return CheckWinner(board)
		}

		compRow, compCol := ComputerChoice(board)
		computerChoices[compRow][compCol] = 2
		UpdateBoard(board, userChoices, computerChoices)
		PrintBoardRune(board)

		if IsGameOver(board) {
			return CheckWinner(board)
		}
	}
}
func CheckIfChoiceExists(board *[3][3]rune, row int, col int) bool {
	return board[row][col] != ' '
}

func UpdateBoard(board *[3][3]rune, userInput [3][3]int8, computerInput [3][3]int8) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == ' ' {
				if userInput[i][j] == 1 {
					board[i][j] = 'X'
				} else if computerInput[i][j] == 2 {
					board[i][j] = 'O'
				}
			}
		}
	}
}

func IsGameOver(board *[3][3]rune) bool {
	return CheckWinner(board) != 'C' || IsBoardFull(board)
}

func CheckWinner(board *[3][3]rune) rune {
	if winner, _ := CheckIfPointHasBeenMade(*board); winner != 'C' {
		return winner
	}
	return 'C'
}

func CheckIfPointHasBeenMade(board [3][3]rune) (rune, bool) {
	lines := [][3]rune{
		{board[0][0], board[0][1], board[0][2]},
		{board[1][0], board[1][1], board[1][2]},
		{board[2][0], board[2][1], board[2][2]},
		{board[0][0], board[1][0], board[2][0]},
		{board[0][1], board[1][1], board[2][1]},
		{board[0][2], board[1][2], board[2][2]},
		{board[0][0], board[1][1], board[2][2]},
		{board[0][2], board[1][1], board[2][0]},
	}

	for _, line := range lines {
		if line[0] != ' ' && line[0] == line[1] && line[1] == line[2] {
			return line[0], true
		}
	}
	return 'C', false
}

func ComputerChoice(board *[3][3]rune) (int, int) {
	for {
		row := rand.Intn(3)
		col := rand.Intn(3)
		if !CheckIfChoiceExists(board, row, col) {
			return row, col
		}
	}
}

func PrintBoardRune(board *[3][3]rune) {
	fmt.Println()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf(" %c ", board[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		if i < 2 {
			fmt.Println("\n-----------")
		}
	}
	fmt.Println()
}

func IsBoardFull(board *[3][3]rune) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == ' ' {
				return false
			}
		}
	}
	return true
}
