package main

import "fmt"

func main() {
	fmt.Print(findSmallestSetOfVertices(5, [][]int{{0, 1}, {2, 1}, {3, 1}, {1, 4}, {2, 4}}))

}

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	myMap := make(map[int]bool)
	result := []int{}
	for _, array := range edges {
		if !myMap[array[1]] {
			myMap[array[1]] = true
		}
	}
	for i := 0; i < n; i++ {
		if !myMap[i] {
			result = append(result, i)
		}
	}
	return result
	// outDegree := make([]int, n)
	// result := []int{}
	// for _, array := range edges {
	// 	outDegree[array[1]]++
	// }
	// for index, val := range outDegree {
	// 	if val == 0 {
	// 		result = append(result, index)
	// 	}
	// }
	// return result

}
