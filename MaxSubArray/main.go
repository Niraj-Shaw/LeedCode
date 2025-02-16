package main

import "fmt"

func main() {
	nums := []int{5, 4, -1, 7, 8}
	fmt.Print(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	// sum := nums[0]
	// max := nums[0]
	// for i := 1; i < len(nums); i++ {
	// 	num := nums[i]
	// 	if num < num+sum {
	// 		sum = num + sum
	// 	} else {
	// 		sum = num

	// 	}
	// 	if sum > max {
	// 		max = sum
	// 	}

	// }

	// return max
	//leftmax, rightmax and mid+left+right
	var FindMaxArray func(left, right int) int
	FindMaxArray = func(left, right int) int {
		// if left > right {
		// 	return -2147483648
		// }
		// mid := (left + right) / 2
		// curr := 0
		// bestLeftSum := 0
		// bestRightSum := 0
		// // Iterate from the middle to the beginning.
		// for i := mid - 1; i >= left; i-- {
		// 	curr += nums[i]
		// 	if bestLeftSum < curr {
		// 		bestLeftSum = curr
		// 	}
		// }
		// // Reset curr and iterate from the middle to the end.
		// curr = 0
		// for i := mid + 1; i <= right; i++ {
		// 	curr += nums[i]
		// 	if bestRightSum < curr {
		// 		bestRightSum = curr
		// 	}
		// }
		// // The bestCombinedSum uses the middle element and the best possible sum from each half.
		// bestCombinedSum := nums[mid] + bestLeftSum + bestRightSum
		// // Find the best subarray possible from both halves.
		// leftHalf := FindMaxArray(left, mid-1)
		// rightHalf := FindMaxArray(mid+1, right)
		// // The largest of the 3 is the answer for any given input array.
		// if bestCombinedSum < leftHalf {
		// 	bestCombinedSum = leftHalf
		// }
		// if bestCombinedSum < rightHalf {
		// 	bestCombinedSum = rightHalf
		// }
		// return bestCombinedSum
		if left > right {
			return -2147483648
		}
		mid := (right + left) / 2
		leftMax, rightMax, curr := 0, 0, 0
		for i := mid - 1; i >= left; i-- {
			curr += nums[i]
			if curr > leftMax {
				leftMax = curr
			}
		}
		curr = 0
		for i := mid + 1; i <= right; i++ {
			curr += nums[i]
			if curr > rightMax {
				rightMax = curr
			}
		}
		bestCombination := leftMax + rightMax + nums[mid]
		leftSubArrayMax := FindMaxArray(left, mid-1)
		rightSubArrayMax := FindMaxArray(mid+1, right)

		if leftSubArrayMax > bestCombination {
			bestCombination = leftSubArrayMax
		}

		if rightSubArrayMax > bestCombination {
			bestCombination = rightSubArrayMax
		}

		return bestCombination

	}
	return FindMaxArray(0, len(nums)-1)

}
