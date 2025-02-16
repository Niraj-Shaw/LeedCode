package main

import "fmt"

var myMap map[rune]int

func main() {
	fmt.Print(frequencySort("2a554442f544asfasssffffasss"))

}

func frequencySort(s string) string {
	myMap = make(map[rune]int)
	result := ""
	for _, val := range s {
		myMap[val]++
	}
	myArray := make([]rune, 0, len(myMap))
	for key, _ := range myMap {
		myArray = append(myArray, key)
	}

	Buildheap(myArray)

	for i := len(myArray) - 1; i >= 0; i-- {
		for j := 0; j < myMap[myArray[0]]; j++ {
			result += string(myArray[0])
		}
		myArray[0], myArray[i] = myArray[i], myArray[0]
		heapify(myArray, 0, i-1)
	}

	return result

}

func Buildheap(nums []rune) {
	for i := len(nums) / 2; i >= 0; i-- {
		heapify(nums, i, len(nums)-1)
	}
}

func heapify(nums []rune, index int, Limit int) {
	idx := index
	left, right := 2*index+1, 2*index+2
	if left <= Limit && myMap[nums[left]] > myMap[nums[idx]] {
		idx = left

	}
	if right <= Limit && myMap[nums[right]] > myMap[nums[idx]] {
		idx = right

	}
	if idx != index {
		nums[idx], nums[index] = nums[index], nums[idx]
		heapify(nums, idx, Limit)
	}

}
