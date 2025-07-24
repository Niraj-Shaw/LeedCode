package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	fmt.Print(threeSumClosest([]int{-1, 2, 1, -4}, 1))

}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	diff := math.MaxInt32
	
	for i := 0; i < len(nums) && 

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
