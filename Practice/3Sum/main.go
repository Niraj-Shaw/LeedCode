package main

import (
	"fmt"
	"sort"
)

var result [][]int

func main() {
	result := get3Sum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Print(result)
}

func get3Sum(nums []int) [][]int {

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		get2Sum(nums, i, -nums[i])

	}

	return result

}
func get2Sum(nums []int, index int, target int) {
	low, high := index+1, len(nums)-1
	for low < high {
		if nums[low]+nums[high] == target {
			result = append(result, []int{nums[index], nums[low], nums[high]})
			for low < high && (nums[low] == nums[low+1]) {
				low++
			}
			for low < high && (nums[high] == nums[high-1]) {
				high--
			}
			low++
			high--
		} else if nums[low]+nums[high] < 0 {
			low++
		} else {
			high--
		}

	}

}
