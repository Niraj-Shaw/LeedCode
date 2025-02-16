package main

import (
	"fmt"
	//"sort"
)

type heap struct {
	array []int
}

func main() {

	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))

}

func findKthLargest(nums []int, k int) int {
	// sort.Ints(nums)
	// return nums[len(nums)-k]
	h := &heap{}
	h.buildHeap(nums)
	fmt.Print(h.array)
	result := 0
	for i := 0; i < k; i++ {
		result = h.pop()
	}

	return result

}

func (h *heap) buildHeap(nums []int) {

	//add it to array
	//find the chils to the same array

	//traverse an array
	for i := 0; i < len(nums); i++ {
		//find the parent
		h.array = append(h.array, nums[i])
		h.heapify(len(h.array) - 1)

	}
}

func (h *heap) heapify(index int) {
	for h.array[Parent(index)] < h.array[index] {
		h.swap(Parent(index), index)
		index = Parent(index)

	}

}

func Parent(index int) int {
	return (index - 1) / 2
}

func Left(index int) int {
	return 2*index + 1
}

func Right(index int) int {
	return 2*index + 2
}

func (h *heap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func (h *heap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := Left(index), Right(index)
	ChildToCompare := 0
	for l <= lastIndex {
		if l == lastIndex {
			ChildToCompare = l
		} else if h.array[l] > h.array[r] {
			ChildToCompare = l
		} else {
			ChildToCompare = r
		}
		if h.array[index] < h.array[ChildToCompare] {
			h.swap(index, ChildToCompare)
			index = ChildToCompare
			l, r = Left(index), Right(index)
		} else {
			return
		}

	}

}
func (h *heap) pop() int {
	if len(h.array) == 0 {
		return -1
	}
	extracted := h.array[0]
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	h.heapifyDown(0)

	return extracted
}
