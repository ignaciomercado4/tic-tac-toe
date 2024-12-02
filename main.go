package main

import (
	"fmt"
	"math/rand"
)

func printBoard(board [][]string) {
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

func main() {
	board := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}

	getUserMovement(board)
	getMachineMovement(board)
	printBoard(board)
}
