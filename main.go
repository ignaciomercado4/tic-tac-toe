package main

import (
	"fmt"
	"math/rand"
)

func printBoard(board [][]string) {
	fmt.Println("")

	for i := 0; i < len(board); i++ {
		fmt.Println(board[i][0], "|", board[i][1], "|", board[i][2])
	}
}

func getUserMovement(board [][]string) {
	fmt.Println("Current board state:")
	printBoard(board)

	var selectedTile int

	fmt.Scan(&selectedTile)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == fmt.Sprintf("%d", selectedTile) &&
				board[i][j] != "X" &&
				board[i][j] != "O" {

				board[i][j] = "X"
				return
			}
		}
	}
}

func getMachineMovement(board [][]string) {
	fmt.Println("Current board state:")
	printBoard(board)

	for {
		randomTile := rand.Intn(9) + 1

		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				if board[i][j] == fmt.Sprintf("%d", randomTile) &&
					board[i][j] != "X" &&
					board[i][j] != "O" {

					board[i][j] = "O"
					return
				}
			}
		}
	}
}

func checkWinner(board [][]string) bool {
	// Horizontal wins
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] &&
			board[i][1] == board[i][2] &&
			board[i][0] != fmt.Sprintf("%d", i*3+1) &&
			board[i][0] != fmt.Sprintf("%d", i*3+2) &&
			board[i][0] != fmt.Sprintf("%d", i*3+3) {

			fmt.Println(board[i][0], " WINS!")
			return true
		}
	}

	// Vertical wins
	for j := 0; j < 3; j++ {
		if board[0][j] == board[1][j] &&
			board[1][j] == board[2][j] &&
			board[0][j] != fmt.Sprintf("%d", j+1) &&
			board[0][j] != fmt.Sprintf("%d", j+4) &&
			board[0][j] != fmt.Sprintf("%d", j+7) {

			fmt.Println(board[0][j], " WINS!")
			return true
		}
	}

	// Diagonal wins
	if board[0][0] == board[1][1] &&
		board[1][1] == board[2][2] &&
		board[0][0] != "1" &&
		board[0][0] != "5" &&
		board[0][0] != "9" {

		fmt.Println(board[0][0], " WINS!")
		return true
	}

	if board[0][2] == board[1][1] &&
		board[1][1] == board[2][0] &&
		board[0][2] != "3" &&
		board[0][2] != "5" &&
		board[0][2] != "7" {

		fmt.Println(board[0][2], " WINS!")
		return true
	}

	return false
}

func main() {
	board := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}

	for !checkWinner(board) {
		getUserMovement(board)
		getMachineMovement(board)
	}

}
