package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Print(longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
	//fmt.Print(longestIncreasingPath([][]int{{7, 7, 5}, {2, 4, 6}, {8, 2, 0}}))
	fmt.Print(longestIncreasingPath([][]int{{1}}))
}

func longestIncreasingPath(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])

	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
	}

	var backtrack func(i, j, prev int) int

	backtrack = func(i, j, prev int) int {
		if i < 0 || i >= n || j < 0 || j >= m {
			return 0
		}

		if matrix[i][j] == -1 || matrix[i][j] <= prev {
			return 0
		}

		if cache[i][j] != 0 {
			return cache[i][j]
		}

		curr := matrix[i][j]

		cache[i][j] = 1 + max(backtrack(i+1, j, curr),
			backtrack(i-1, j, curr),
			backtrack(i, j+1, curr),
			backtrack(i, j-1, curr))

		return cache[i][j]

	}

	longest := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			path := backtrack(i, j, math.MinInt64)
			longest = max(longest, path)
		}
	}
	return longest
}
