package main

import "fmt"

type Set struct {
	i int
	j int
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Print(numIslands(grid))

}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	queue := make([]Set, 0)
	n, m := len(grid), len(grid[0])
	count := 0
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			point := Set{r, c}
			if grid[r][c] == '1' {
				queue = append(queue, point)
				grid[r][c] = '0'
				count++

				for len(queue) != 0 {
					node := queue[0]
					queue = queue[1:]
					row := node.i
					col := node.j
					if (row < n) && (col-1 >= 0) && (grid[row][col-1] == '1') {
						queue = append(queue, Set{row, col - 1})
						grid[row][col-1] = '0'
					}
					if (row < n) && (col+1 < m) && (grid[row][col+1] == '1') {
						queue = append(queue, Set{row, col + 1})
						grid[row][col+1] = '0'
					}
					if (row-1 >= 0) && (col < m) && (grid[row-1][col] == '1') {
						queue = append(queue, Set{row - 1, col})
						grid[row-1][col] = '0'
					}
					if (row+1 < n) && (col < m) && (grid[row+1][col] == '1') {
						queue = append(queue, Set{row + 1, col})
						grid[row+1][col] = '0'
					}
				}
			}

		}
	}
	return count
}
