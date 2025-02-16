package Heap

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// InitHeap initializes a heap from a given IntHeap.
func InitHeap(h *IntHeap) {
	heap.Init(h)
}

// PushHeap pushes an element onto the heap.
func PushHeap(h *IntHeap, x int) {
	heap.Push(h, x)
}

// PopHeap pops the smallest element from the heap.
func PopHeap(h *IntHeap) int {
	return heap.Pop(h).(int)
}
