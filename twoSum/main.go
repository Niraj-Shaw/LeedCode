package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
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
	///--> new version
	// mp := make(map[int]int)
	// for i, val := range nums {
	// 	if j, ok := mp[target-val]; ok {
	// 		return []int{i, j}
	// 	} else {
	// 		mp[val] = i
	// 	}
	// }
	// return []int{}
	// --> new version
	mp := make(map[int]int)

	for i, v := range nums {
		sum := target - v

		if val, ok := mp[sum]; ok {
			return []int{i, val}
		}
		mp[v] = i
	}

	return []int{}

}
