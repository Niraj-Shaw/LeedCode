package main

import "fmt"

func main() {
	arr := [][]int{{1, 2}, {3, 4}}
	result := matrixReshape(arr, 2, 4)
	fmt.Print(result)
}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	n, m := len(mat), len(mat[0])
	if (n * m) != (r * c) {
		return mat
	}
	result := [][]int{}
	temp := 0
	tempArr := []int{}
	for _, value := range mat {
		for _, val := range value {
			if temp < c {
				tempArr = append(tempArr, val)
				temp++
			}
			if temp >= c {
				result = append(result, tempArr)
				tempArr = nil
				temp = 0
			}

		}
	}
	return result
}
