package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func lengthOfLIS(nums []int) int {
	start := 0
	index := 0
	longest := 0
	count := 0
	prev := 0

	for index < len(nums) {
		curr := nums[index]
		if curr > start && curr > prev {
			count++
		} else if curr < start {
			start = curr
			count = 0
		}

		if count > longest {
			longest = count
		}

		prev = curr
		index++
	}

	return longest
}
