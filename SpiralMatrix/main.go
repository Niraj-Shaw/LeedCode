package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return []int{}
	}
	top, bottom, left, right := 0, m-1, 0, n-1
	result := []int{}

	for len(result) < m*n {
		for c := left; c <= right; c++ {
			result = append(result, matrix[top][c])
		}

		for r := top + 1; r <= bottom; r++ {
			result = append(result, matrix[r][right])
		}

		if top != bottom {
			for c := right - 1; c >= left; c-- {
				result = append(result, matrix[bottom][c])
			}
		}

		if left != right {
			for r := bottom - 1; r > top; r-- {
				result = append(result, matrix[r][left])
			}
		}

		top++
		right--
		left++
		bottom--
	}

	return result
}
