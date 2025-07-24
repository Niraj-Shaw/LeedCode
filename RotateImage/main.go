package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	rotateImage(matrix)
	fmt.Print(matrix)
}

func rotateImage(matrix [][]int) {
	// r, c := len(matrix), len(matrix[0])
	// for i := 0; i < r; i++ {
	// 	for j := i; j < c; j++ {
	// 		matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
	// 	}
	// }
	// for i := 0; i < r; i++ {
	// 	for j, z := 0, c-1; j < z; j, z = j+1, z-1 {
	// 		matrix[i][j], matrix[i][z] = matrix[i][z], matrix[i][j]
	// 	}
	// }

	///
	// n := len(matrix)
	// for i := 0; i < n/2; i++ {
	// 	for j := 0; j < (n+1)/2; j++ {
	// 		temp := matrix[n-j-1][i]
	// 		matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
	// 		matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
	// 		matrix[j][n-i-1] = matrix[i][j]
	// 		matrix[i][j] = temp

	// 	}
	// }

	//

	r, c := len(matrix), len(matrix[0])

	for i := 0; i < r; i++{
		for j := i; j< c; j++{
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < r; i++{
		for j, z := 0, c - 1; j < z; j, z = j+1, z-1{
			matrix[i][j], matrix[i][z] = matrix[i][z], matrix[i][j]
		}
	}

}
