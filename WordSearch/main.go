package main

import "fmt"

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Print(exist(board, "ABCCED"))
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	rows, cols := len(board), len(board[0])
	var backTracking func(row int, col int, index int) bool
	backTracking = func(row int, col int, index int) bool {
		if index >= len(word) {
			return true
		}
		if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != word[index] {
			return false
		}
		temp := board[row][col]
		board[row][col] = 0
		result := backTracking(row+1, col, index+1) || backTracking(row-1, col, index+1) ||
			backTracking(row, col+1, index+1) || backTracking(row, col-1, index+1)
		board[row][col] = temp
		return result
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if (board[i][j] == word[0]) && (backTracking(i, j, 0) == true) {
				return true
			}
		}
	}
	return false
}
