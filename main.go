package main

import "fmt"

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

func main() {
	board := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}

	getUserMovement(board)
	printBoard(board)
}
