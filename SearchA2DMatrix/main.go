package main

import "fmt"

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	fmt.Print(searchMatrix(matrix, 3))
}

func searchMatrix(matrix [][]int, target int) bool {
	r, c := len(matrix), len(matrix[0])
	if r == 0 || c == 0 {
		return false
	}
	left, right := 0, r*c-1
	for left <= right {
		mid := (left + right) / 2
		i := mid / c
		j := mid % c
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			left = mid + 1
		} else {

			right = mid - 1
		}
	}
	return false

}
