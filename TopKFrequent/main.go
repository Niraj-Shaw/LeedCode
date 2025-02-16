package main

import (
	"container/heap"
	"fmt"
)

//	type heap struct {
//		array []int
//		myMap map[int]int
//
// // }
type MyHeap []int

var myMap map[int]int

func (h MyHeap) Len() int               { return len(h) }
func (h MyHeap) Swap(i, j int)          { h[i], h[j] = h[j], h[i] }
func (h MyHeap) Less(i int, j int) bool { return myMap[h[i]] < myMap[h[j]] }
func (h *MyHeap) Push(x interface{})    { *h = append(*h, x.(int)) }
func (h *MyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	fmt.Println(topKFrequent([]int{6, 0, 1, 4, 9, 7, -3, 1, -4, -8, 4, -7, -3, 3, 2, -3, 9, 5, -4, 0}, 6))

}

func topKFrequent(nums []int, k int) []int {
	myMap = make(map[int]int)
	myHeap := &MyHeap{}
	for _, val := range nums {
		myMap[val]++
	}
	myArray := make([]int, 0, len(myMap))
	for key, _ := range myMap {
		myArray = append(myArray, key)

	}
	for i := 0; i < k; i++ {
		heap.Push(myHeap, myArray[i])
	}
	//result := []int{8, 15, 6}
	for i := k; i < len(myArray); i++ {
		heap.Push(myHeap, myArray[i])
		heap.Pop(myHeap)

	}
	// *myHeap = result
	// heap.Init(myHeap)
	// heap.Push(myHeap, 18)
	// max := heap.Pop(myHeap)
	// fmt.Printf("max : %d\n", max)
	return *myHeap

}

// func topKFrequent(nums []int, k int) []int {

// 	h := &heap{}
// 	h.myMap = make(map[int]int)

// 	for _, val := range nums {
// 		h.myMap[val]++
// 	}
// 	myArray := make([]int, 0, len(h.myMap))
// 	for key, _ := range h.myMap {
// 		myArray = append(myArray, key)
// 	}

// 	h.array = make([]int, k)
// 	for i := 0; i < k; i++ {
// 		h.array[i] = myArray[i]

// 	}
// 	h.BuildHeap()

// 	for i := k; i < len(myArray); i++ {
// 		if h.myMap[myArray[i]] > h.myMap[h.array[0]] {
// 			h.array[0] = myArray[i]
// 			h.heapify(0)

// 		}

// 	}
// 	return h.array

// }

// func (h *heap) BuildHeap() {
// 	for i := len(h.array) / 2; i >= 0; i-- {
// 		h.heapify(i)

// 	}

// }

// func (h *heap) heapify(index int) {
// 	largest := index
// 	l, r := index*2+1, index*2+2
// 	if l < len(h.array) && h.myMap[h.array[l]] < h.myMap[h.array[largest]] {
// 		largest = l

// 	}
// 	if r < len(h.array) && h.myMap[h.array[r]] < h.myMap[h.array[largest]] {
// 		largest = r
// 	}

// 	if largest != index {
// 		h.array[largest], h.array[index] = h.array[index], h.array[largest]
// 		h.heapify(largest)
// 	}
// }
