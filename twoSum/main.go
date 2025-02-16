package main

import "fmt"

func main() {
	fmt.Print(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Print(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	// mp := make(map[int]int)
	// result := make([]int, 0, 2)
	// for index, val := range nums {
	// 	remainder := target - val
	// 	if _, ok := mp[remainder]; ok {
	// 		result = append(result, mp[remainder], index)
	// 		break
	// 	} else {
	// 		mp[val] = index
	// 	}
	// }
	// return result
	mp := make(map[int]int)
	for i, val := range nums {
		if j, ok := mp[target-val]; ok {
			return []int{i, j}
		} else {
			mp[val] = i
		}
	}
	return []int{}

}
