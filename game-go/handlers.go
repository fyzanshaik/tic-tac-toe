package main

import (
	"fmt"
	"math/rand"
)

func startGame(board *[3][3]rune, indexReference *[3][3]string) rune {
	fmt.Println("This is an empty board: ")
	printBoardRune(board)
	fmt.Println("To enter your choice you need to give two values: row & column number ")
	fmt.Println("Reference: ")
	printBoardString(indexReference)

	var userChoices [3][3]int8
	var computerChoices [3][3]int8
	gameStartFlag := true
	var firstIndex, secondIndex int
	var checkFlag bool

	for gameStartFlag {
		fmt.Printf("Enter your input (row and column): ")
		fmt.Scanf("%d %d", &firstIndex, &secondIndex)

		checkFlag = checkIfChoiceExists(board, firstIndex, secondIndex)
		if checkFlag {
			fmt.Println("That choice has already been taken, please try again")
			continue
		}

		userChoices[firstIndex][secondIndex] = 1

		firstIndex, secondIndex = computerChoice(board)
		computerChoices[firstIndex][secondIndex] = 2

		updateBoard(board, userChoices, computerChoices)
		printBoardRune(board)

		hasPointBeenMade, whichPlayer := checkIfPointHasBeenMade(*board)
		if hasPointBeenMade {
			if whichPlayer == 'X' {
				fmt.Println("You have won!")
				return 'X'
			}
			fmt.Println("Computer has won!")
			return 'O'
		}

		if isBoardFull(board) {
			fmt.Println("The game is a draw!")
			return 'D'
		}
	}

	return 'C'
}

func checkIfPointHasBeenMade(board [3][3]rune) (bool, rune) {
	for i := 0; i < 3; i++ {
		if (board[i][0] == board[i][1]) && (board[i][1] == board[i][2]) && board[i][0] != ' ' {
			if board[i][0] == 'X' {
				return true, 'X'
			} else if board[i][0] == 'O' {
				return true, 'O'
			}
		}
	}

	for j := 0; j < 3; j++ {
		if (board[0][j] == board[1][j]) && (board[1][j] == board[2][j]) && board[0][j] != ' ' {
			if board[0][j] == 'X' {
				return true, 'X'
			} else if board[0][j] == 'O' {
				return true, 'O'
			}
		}
	}

	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != ' ' {
		if board[0][0] == 'X' {
			return true, 'X'
		} else if board[0][0] == 'O' {
			return true, 'O'
		}
	}

	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[0][2] != ' ' {
		if board[0][2] == 'X' {
			return true, 'X'
		} else if board[0][2] == 'O' {
			return true, 'O'
		}
	}

	return false, 'C'
}

func checkIfChoiceExists(choices *[3][3]rune, row int, col int) bool {
	if choices[row][col] == 'X' || choices[row][col] == 'O' {
		return true
	}
	return false
}

func updateBoard(board *[3][3]rune, userInput [3][3]int8, computerInput [3][3]int8) {
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

func computerChoice(choices *[3][3]rune) (first int, second int) {
	for {
		first = rand.Intn(3)
		second = rand.Intn(3)
		value := checkIfChoiceExists(choices, first, second)
		if value {
			continue
		} else {
			break
		}
	}

	return first, second
}

func printBoardRune(board *[3][3]rune) {
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

func printBoardString(board *[3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%s ", board[i][j])
		}
		fmt.Println()
	}
}

func isBoardFull(board *[3][3]rune) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == ' ' {
				return false
			}
		}
	}
	return true
}
