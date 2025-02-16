package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, -7, 1, 8, 16, -2}
	result := threeSum(nums)
	for _, array := range result {
		for _, val := range array {
			fmt.Printf("%d ", val)
		}
		println()
	}
}

func twoSum(nums []int, result [][]int, start int, firstInt int) [][]int {

	end := len(nums) - 1
	remainder := -firstInt
	for j := start; j < len(nums); j++ {

	}
	for start < end {
		if (nums[start] + nums[end]) > remainder {
			end--
		} else if (nums[start] + nums[end]) < remainder {

			start++
		} else {
			//myarr := []int{}
			myarr := append([]int{}, firstInt, nums[start], nums[end])
			result = append(result, myarr)
			start++
			for nums[start-1] == nums[start] && start < end {
				start++
			}
		}
	}
	return result

}

func threeSum(nums []int) [][]int {

	//myMap := map[int]int{
	//		nums[0]: 0,
	//	}
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if (i != 0) && (nums[i] == nums[i-1]) {
			continue
		}
		start := i + 1

		//	first := nums[i
		result = twoSum(nums, result, start, nums[i])

	}
	return result
}

// func threeSum(nums []int) [][]int {
// 	result := [][]int{}
// 	mp := map[int]int{}
// 	sort.Ints(nums)
// 	for index, val := range nums {
// 		mp[val] = index
// 	}

// 	for i := 0; i < len(nums)-2; i++ {
// 		if i == 0 || nums[i] != nums[i-1] {
// 			res := twoSum(nums, i, mp)
// 			result = append(result, res...)
// 		}
// 		// find val + 2Sum = 0{}

// 	}
// 	return result

// }

// func twoSum(nums []int, index int, mp map[int]int) [][]int {
// 	result := [][]int{}
// 	remainder := -nums[index]
// 	var prev int

// 	for i := index + 1; i < len(nums); i++ {
// 		if i != index+1 && nums[i] == prev {
// 			continue
// 		}
// 		target := remainder - nums[i]
// 		if idx, Ok := mp[target]; Ok && idx >= i+1 {
// 			result = append(result, []int{nums[index], nums[i], target})
// 		}
// 		prev = nums[i]
// 	}
// 	return result

// }
