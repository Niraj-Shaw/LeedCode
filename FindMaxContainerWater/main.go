package main

import (
	"fmt"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maximumArea(height))
}

func maximumArea(height []int) int {

	left, right := 0, len(height)-1
	max := 0
	area := 0
	for left != right {
		if height[left] < height[right] {
			area = height[left] * (right - left)
			left++
		} else {
			area = height[right] * (right - left)
			right--
		}
		if area > max {
			max = area
		}
	}
	return max

	// 	left := 0
	// 	right := len(height) - 1
	// 	maxArea := 0
	// 	for left != right {
	// 		area := 0
	// 		if height[left] < height[right] {
	// 			area = height[left] * (right - left)
	// 			left++
	// 		} else {
	// 			area = height[right] * (right - left)
	// 			right--
	// 		}
	// 		if area > maxArea {
	// 			maxArea = area
	// 		}
	// 	}

	// 	return maxArea
	//
}
