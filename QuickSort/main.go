package main

import "fmt"

func main() {
	fmt.Print(quickSort([]int{0, 5, 2, 1, 6, 3}))

}

func quickSort(array []int) []int {
	doQuickSort(array, 0, len(array)-1)
	return array

}

func doQuickSort(array []int, start, end int) {
	if end-start <= 0 {
		return
	}
	pivotIndex := doPartion(array, start, end)
	doQuickSort(array, start, pivotIndex-1)
	doQuickSort(array, pivotIndex+1, end)
}

func doPartion(array []int, start, end int) int {
	pivotIndex, pivot := end, array[end]
	end--
	for start < end {
		for array[start] < pivot {
			start++
		}
		for end >= 0 && array[end] > pivot {
			end--
		}
		if start < end {
			array[start], array[end] = array[end], array[start]
			start++
		}

	}
	if array[start] > pivot {
		array[start], array[pivotIndex] = array[pivotIndex], array[start]
	}
	return start

}
