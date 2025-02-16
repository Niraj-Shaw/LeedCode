package main

import "fmt"

func main() {
	fmt.Print(findJudge(1, [][]int{}))

}

func findJudge(n int, trust [][]int) int {
	inward := make([]int, n+1)
	outward := make([]int, n+1)
	if len(trust) == 0 {
		return n
	}
	for _, array := range trust {
		if len(array) == 0 {
			return n
		}
		outward[array[0]]--
		inward[array[1]]++
	}
	for i := 0; i <= n; i++ {
		if inward[i] == n-1 && outward[i] == 0 {
			return i
		}
	}

	return -1
	// judges := make([]int, N)
	// for _, array := range trust {
	// 	judges[array[0]-1]--
	// 	judges[array[1]-1]++
	// }
	// for index, val := range judges {
	// 	if val == N-1 {
	// 		return index + 1
	// 	}

	//return -1

}
