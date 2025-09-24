package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

func canJump(nums []int) bool {
	n := len(nums) - 1
	curr := n

	for i := n - 1; i >= 0; i-- {
		jump := nums[i]

		if jump >= curr-i {
			curr = i
		}

	}

	return curr == 0

	// if len(nums) == 0 {
	// 	return false
	// }

	// canJump := make(map[int]bool)
	// n := len(nums) - 1

	// var backTrack func(i int) bool

	// backTrack = func(i int) bool {

	// 	maxJump := nums[i]

	// 	if i+maxJump >= n {
	// 		return true
	// 	}

	// 	for j := maxJump; j >= 1; j-- {
	// 		if val, ok := canJump[i+j]; !ok || val {
	// 			if backTrack(i + j) {
	// 				return true
	// 			}
	// 		}

	// 	}

	// 	canJump[i] = false

	// 	return false
	// }

	// return backTrack(0)

}
