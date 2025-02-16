package main

import "fmt"

func main() {
	fmt.Print(InsertionSort([]int{4, 2, 7, 1, 3}))

}

func InsertionSort(arr []int) []int {
	index := 1
	for index < len(arr) {
		temp := arr[index]
		i := index - 1
		for i >= 0 {
			if temp < arr[i] {
				arr[i+1] = arr[i]
				arr[i] = temp
				i--
			} else {
				break
			}

		}

		index++
	}

	return arr

}
