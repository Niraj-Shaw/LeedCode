package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 1}
	fmt.Print(DeleteAndEarn(nums))
}

func DeleteAndEarn(nums []int) int {
	sort.Ints(nums)
	nums2 := []int{}
	myMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := myMap[nums[i]]; !ok {
			myMap[nums[i]] = 1
			nums2 = append(nums2, nums[i])
		} else {
			myMap[nums[i]]++
		}

	}
	CurEarn, band1, band2 := 0, 0, nums2[0]*myMap[nums2[0]]
	for i := 1; i < len(nums2); i++ {
		CurEarn = nums2[i] * myMap[nums2[i]]
		temp := band2
		if nums2[i] == nums2[i-1]+1 {
			//checks which band has max value
			if CurEarn+band1 > band2 {
				band2 = CurEarn + band1
			}

		} else {
			//no checking, just add band 2 and swap old valuw with band1
			band2 = CurEarn + band2
		}
		band1 = temp
	}
	return band2
}
