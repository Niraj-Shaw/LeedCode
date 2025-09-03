package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 9))
}

func searchRange(nums []int, target int) []int {

	start, end := -1, -1

	index := findTarget(0, len(nums)-1, nums, target)

	if index == -1 {
		return []int{-1, -1}
	}

	start, end = index, index

	for start > 0 && nums[start-1] == target {
		start = findTarget(0, start-1, nums, target)
	}

	for end >= 0 && end < len(nums)-1 && nums[end+1] == target {
		end = findTarget(end+1, len(nums)-1, nums, target)

	}

	return []int{start, end}

}

func findTarget(start int, end int, nums []int, target int) int {
	low, high := start, end

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
