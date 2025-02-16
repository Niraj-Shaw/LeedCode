package main

import "fmt"

func main() {
	fmt.Print(search([]int{1}, 0))

}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		//mid := start + (end-start)/2
		if target == nums[mid] {
			return mid
		} else if nums[mid] >= nums[start] {
			if target >= nums[start] && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else if nums[mid] < nums[start] {
			if target <= nums[end] && target > nums[mid] {
				start = mid + 1
			} else {
				end = mid - 1
			}

		}
	}

	return -1

}
