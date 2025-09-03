package main

import "container/heap"

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	//nums []int
	small *MaxHeap
	large *MinHeap
}

func Constructor() MedianFinder {
	// return MedianFinder{
	// 	nums: make([]int, 0),
	// }
	s := &MaxHeap{}
	l := &MinHeap{}
	heap.Init(s)
	heap.Init(l)
	return MedianFinder{s, l}
}
func (this *MedianFinder) addNum(num int) {
	//this.nums = append(this.nums, num)
	heap.Push(this.small, num)

	heap.Push(this.large, heap.Pop(this.small))

	if len(*this.large) > len(*this.small) {
		x := heap.Pop(this.large)
		heap.Push(this.small, x)
	}
}

func (this *MedianFinder) findMedian() float64 {
	// sort.Ints(this.nums)
	// length := len(this.nums)
	// if length%2 == 0 {
	// 	id := length / 2
	// 	return float64(this.nums[id]+this.nums[id-1]) / 2.0
	// } else {
	// 	return float64(this.nums[length/2])
	// }
	if (len(*this.large)+len(*this.small))%2 == 0 {
		return float64((*this.small)[0]+(*this.large)[0]) / 2.0
	} else {
		return float64((*this.small)[0])
	}
}

// package main

// import (
// 	"container/heap"
// 	//"fmt"
// )

// // MaxHeap for smaller half
// type MaxHeap []int

// func (h MaxHeap) Len() int           { return len(h) }
// func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Reversed for max-heap
// func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *MaxHeap) Push(x interface{}) {
// 	*h = append(*h, x.(int))
// }

// func (h *MaxHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

// // MinHeap for larger half
// type MinHeap []int

// func (h MinHeap) Len() int           { return len(h) }
// func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // Standard for min-heap
// func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *MinHeap) Push(x interface{}) {
// 	*h = append(*h, x.(int))
// }

// func (h *MinHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

// // MedianFinder struct
// type MedianFinder struct {
// 	small *MaxHeap
// 	large *MinHeap
// }

// // Constructor to initialize the MedianFinder
// func Constructor() MedianFinder {
// 	small := &MaxHeap{}
// 	large := &MinHeap{}
// 	heap.Init(small)
// 	heap.Init(large)
// 	return MedianFinder{small, large}
// }

// // AddNum inserts a number into the data structure
// func (mf *MedianFinder) AddNum(num int) {
// 	heap.Push(mf.small, num) // Add to max-heap first

// 	// Balance: ensure all elements in min-heap are greater
// 	heap.Push(mf.large, heap.Pop(mf.small).(int))

// 	// If min-heap has more elements, move one to max-heap
// 	if mf.large.Len() > mf.small.Len() {
// 		heap.Push(mf.small, heap.Pop(mf.large).(int))
// 	}
// }

// // FindMedian calculates the median
// func (mf *MedianFinder) FindMedian() float64 {
// 	if mf.small.Len() > mf.large.Len() {
// 		return float64((*mf.small)[0])
// 	}
// 	return float64((*mf.small)[0]+(*mf.large)[0]) / 2.0
// }

// // Testing the implementation
// func main() {
// 	mf := Constructor()
// 	mf.AddNum(1)
// 	mf.AddNum(2)
// 	fmt.Println(mf.FindMedian()) // Output: 1.5

// 	mf.AddNum(3)
// 	fmt.Println(mf.FindMedian()) // Output: 2
// }
