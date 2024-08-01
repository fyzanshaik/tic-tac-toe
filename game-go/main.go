package main

import (
	"fmt"
)

func main() {
	var userScore, computerScore int
	for {
		board := [3][3]rune{
			{' ', ' ', ' '},
			{' ', ' ', ' '},
			{' ', ' ', ' '},
		}
		indexBoard := [3][3]string{
			{"00", "01", "02"},
			{"10", "11", "12"},
			{"20", "21", "22"},
		}
		fmt.Println("Tic-Tac-Toe Game: ")
		fmt.Printf("User Score: %d, Computer Score: %d\n", userScore, computerScore)
		fmt.Print("Are you ready to play? (y/n): ")
		var input rune
		fmt.Scanf("%c\n", &input)

		if input == 'y' {
			winner := startGame(&board, &indexBoard)
			if winner == 'X' {
				userScore++
			} else if winner == 'O' {
				computerScore++
			}
		} else {
			fmt.Println("Hope you play next time :)")
			break
		}
		fmt.Println()
		fmt.Print("Would you like to play again? (y/n): ")

		var input1 rune
		fmt.Scanf("%c\n", &input1)
		if input1 != 'y' {
			fmt.Println("Thank you for playing!")
			break
		}
	}
}
