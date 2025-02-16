package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	arr := intersect(nums1, nums2)
	fmt.Print(arr)

}

func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	sort.Ints(nums1)
	sort.Ints(nums2)
	for p, q := 0, 0; p < len(nums1) && q < len(nums2); {
		if nums1[p] < nums2[q] {
			p++

		} else if nums1[p] > nums2[q] {
			q++
		} else if nums1[p] == nums2[q] {
			result = append(result, nums1[p])
			p++
			q++
		}
	}
	return result
}
