package main

import "fmt"

func main() {
	board := [9][9]byte{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9}}
	result := isValidSudoku(board)
	fmt.Print(result)

}

func isValidSudoku(board [9][9]byte) bool {
	rowMap := [9]map[byte]bool{}
	colMap := [9]map[byte]bool{}
	blockMap := [9]map[byte]bool{}
	for i := 0; i < 9; i++ {
		rowMap[i] = make(map[byte]bool)
		colMap[i] = make(map[byte]bool)
		blockMap[i] = make(map[byte]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				continue
			}
			if _, ok := rowMap[i][board[i][j]]; ok {
				return false
			} else {
				rowMap[i][board[i][j]] = true
			}

			if _, ok := colMap[j][board[i][j]]; ok {
				return false
			} else {
				colMap[j][board[i][j]] = true
			}

			if _, ok := blockMap[i/3*3+j/3][board[i][j]]; ok {
				return false
			} else {
				blockMap[i/3*3+j/3][board[i][j]] = true
			}
		}
	}
	return true
}
