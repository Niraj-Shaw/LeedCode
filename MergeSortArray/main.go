package main

import "fmt"

func main() {
	nums1 := []int{4, 5, 6, 0, 0, 0}
	nums2 := []int{1, 2, 8}
	m, n := 3, 3
	merge(nums1, m, nums2, n)
	fmt.Print(nums1)

}

func merge(nums1 []int, m int, nums2 []int, n int) {

	last := m + n - 1
	mp := m - 1
	np := n - 1

	for mp >= 0 && np >= 0 {

		if nums2[np] >= nums1[mp] {
			nums1[last] = nums2[np]
			np--
		} else {
			nums1[last] = nums1[mp]
			mp--
		}
		last--

	}
	for np >= 0 {
		nums1[last] = nums2[np]
		np--
		last--
	}
}
