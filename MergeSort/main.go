// package main

// import "fmt"

// func main() {
// 	array := []int{6, 5, 12, 10, 9, 1}
// 	Sort(array)
// 	fmt.Print(array)

// }

// func Sort(array []int) {
// 	if len(array) == 0 {
// 		return
// 	}
// 	//mergeSort(array, 0, len(*array)-1)
// 	mergeSort(array)

// }

// // func mergeSort(array *[]int, low, high int) {
// // 	if low < high {
// // 		mid := (low + high) / 2
// // 		mergeSort(array, low, mid)
// // 		mergeSort(array, mid+1, high)
// // 		merge(array, low, mid, high)
// // 	}
// // }

// // func merge(array *[]int, low, mid, high int) {
// // 	arr1, arr2 := []int{}, []int{}
// // 	for i := low; i <= mid; i++ {
// // 		arr1 = append(arr1, (*array)[i])
// // 	}
// // 	for j := 0; j <= high-mid-1; j++ {
// // 		arr2 = append(arr2, (*array)[mid+1+j])
// // 	}
// // 	i, j, k := 0, 0, low
// // 	for i < len(arr1) && j < len(arr2) {
// // 		if arr1[i] <= arr2[j] {
// // 			(*array)[k] = arr1[i]
// // 			i++
// // 		} else {
// // 			(*array)[k] = arr2[j]
// // 			j++
// // 		}
// // 		k++
// // 	}
// // 	for i < len(arr1) {
// // 		(*array)[k] = arr1[i]
// // 		i++
// // 		k++
// // 	}
// // 	for j < len(arr2) {
// // 		(*array)[k] = arr2[j]
// // 		j++
// // 		k++
// // 	}
// // }
// func mergeSort(array []int) {
// 	size := len(array)
// 	if size <= 1 {
// 		return
// 	}
// 	mid := size / 2
// 	l, r := make([]int, mid), make([]int, size-mid)
// 	copy(l, array[:mid])
// 	copy(r, array[mid:])
// 	mergeSort(l)
// 	mergeSort(r)
// 	merge(array, l, r)
// }

// func merge(array, l, r []int) {
// 	i, j, k := 0, 0, 0
// 	for i < len(l) && j < len(r) {
// 		if l[i] <= r[j] {
// 			array[k] = l[i]
// 			i++
// 		} else {
// 			array[k] = r[j]
// 			j++
// 		}
// 		k++
// 	}
// 	for i < len(l) {
// 		array[k] = l[i]
// 		i++
// 		k++
// 	}
// 	for j < len(r) {
// 		array[k] = r[j]
// 		j++
// 		k++
// 	}

// }
package main

import "fmt"

func main() {
	array := []int{6, 5, 12, 10, 9, 1}
	mergeSort(&array)
	fmt.Println(array)
}

func mergeSort(array *[]int) {
	if len(*array) <= 1 {
		return
	}

	mid := len(*array) / 2
	left := (*array)[:mid]
	right := (*array)[mid:]

	mergeSort(&left)
	mergeSort(&right)

	merge(array, &left, &right)
}

func merge(array *[]int, left, right *[]int) {
	leftLen := len(*left)
	rightLen := len(*right)

	result := make([]int, leftLen+rightLen)
	i, j, k := 0, 0, 0

	for i < leftLen && j < rightLen {
		if (*left)[i] <= (*right)[j] {
			result[k] = (*left)[i]
			i++
		} else {
			result[k] = (*right)[j]
			j++
		}
		k++
	}

	for i < leftLen {
		result[k] = (*left)[i]
		i++
		k++
	}

	for j < rightLen {
		result[k] = (*right)[j]
		j++
		k++
	}

	*array = result
}
