package main

import "fmt"

func main() {
	arr := []int{2, 7, 11, 15}
	result := TwoSum(arr, 9)
	fmt.Print(result)

}

func TwoSum(arr []int, target int) []int {
	mp := map[int]int{}
	for i := 0; i < len(arr); i++ {
		remainder := -(arr[i] - target)
		if val, ok := mp[remainder]; ok {
			return []int{i, val}
		} else {
			mp[arr[i]] = i
		}

	}

	return []int{}
}
