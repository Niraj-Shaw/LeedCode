package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(binarySearch([]int{1, 9, 6, 4, 3, 5, 10, 8}, 8))
}

func binarySearch(nums []int, num int) bool {
	sort.Ints(nums)

	lo := 0
	hi := len(nums) - 1

	for lo <= hi {
		mid := (hi + lo) / 2

		if nums[mid] == num {
			return true
		} else if num > nums[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return false

}
