package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))

}

// func findMedianSortedArrays(arr1, arr2 []int) float64 {
// 	l1, l2 := len(arr1), len(arr2)
// 	n := (l1 + l2)
// 	if n%2 == 1 {
// 		return float64(findMedian(arr1, arr2, n/2, 0, l1-1, 0, l2-1))
// 	} else {
// 		return float64(findMedian(arr1, arr2, n/2, 0, l1-1, 0, l2-1)+findMedian(arr1, arr2, n/2-1, 0, l1-1, 0, l2-1)) / 2
// 	}

// }

// func findMedian(arr1, arr2 []int, k, aleft, aRight, bLeft, bRight int) int {
// 	if aleft > aRight {
// 		return arr2[k-aleft]
// 	}
// 	if bLeft > bRight {
// 		return arr1[k-bLeft]
// 	}
// 	aIndex := (aleft + aRight) / 2
// 	bIndex := (bLeft + bRight) / 2
// 	aValue := arr1[aIndex]
// 	bValue := arr2[bIndex]

// 	index := (aIndex + bIndex)

// 	if k > index {
// 		if aValue < bValue {
// 			return findMedian(arr1, arr2, k, aIndex+1, aRight, bLeft, bRight)
// 		} else {
// 			return findMedian(arr1, arr2, k, aleft, aRight, bIndex+1, bRight)
// 		}
// 	} else {
// 		if aValue < bValue {
// 			return findMedian(arr1, arr2, k, aleft, aRight, bLeft, bIndex-1)
// 		} else {
// 			return findMedian(arr1, arr2, k, aleft, aIndex-1, bLeft, bRight)
// 		}
// 	}

// }
func findMedianSortedArrays(arr1, arr2 []int) float64 {
	if len(arr2) < len(arr1) {
		return findMedianSortedArrays(arr2, arr1)
	}
	left, right := 0, len(arr1)
	for left <= right {
		n, m := len(arr1), len(arr2)
		PartA := (left + right) / 2
		PartB := (n+m+1)/2 - PartA
		AMax, AMin := Cond(PartA == 0, math.MinInt, arr1, PartA-1), Cond(PartA == n, math.MaxInt, arr1, PartA)
		BMax, BMin := Cond(PartB == 0, math.MinInt, arr2, PartB-1), Cond(PartB == m, math.MaxInt, arr2, PartB)
		if AMax <= BMin && BMax <= AMin {
			if (m+n)%2 == 1 {
				return float64(Max(AMax, BMax))
			} else {
				return float64(Max(AMax, BMax)+Min(AMin, BMin)) / 2
			}
		} else if AMax > BMin {
			right = PartA - 1
		} else {
			left = PartA + 1
		}

	}
	return 0

}
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Cond(exp bool, def int, nums []int, index int) int {
	if exp {
		return def
	} else {
		return nums[index]
	}
}
