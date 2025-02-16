package main

import "fmt"

func main() {
	nums := []int{1, 0, 1, 1, 0, 1}
	fmt.Print(FindMaxConsecutiveOnes(nums))

}

func FindMaxConsecutiveOnes(nums []int) int {
	//loop
	LastZeroIndex := 0
	max := 0
	zeroCount := 0
	start := 0
	end := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
			if zeroCount == 1 && i < len(nums) {
				LastZeroIndex = i
			}
		}

		if zeroCount == 2 {
			start = LastZeroIndex
			end = i
			LastZeroIndex = i
			zeroCount--

		} else {
			end++
		}
		if (end - start) > max {
			max = end - start
		}
		// if(zeroCount > 1){
		//   zeroCount = 0;
		//   count = 0;
		// }

	}
	return max
}
