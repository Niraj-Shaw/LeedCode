package main

import (
	"Container/heap"
	"fmt"
)

// type node struct {
// 	index    int
// 	distance float64
// }

func main() {

	fmt.Print(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))

}

func kClosest(points [][]int, k int) [][]int {
	// result := [][]int{}
	// heap := make([]node, k)
	// for i := 0; i < k; i++ {
	// 	heap[i] = node{i, getDist(points[i])}
	// }
	// buildHeap(heap)
	// for i := k; i < len(points); i++ {
	// 	dist := getDist(points[i])
	// 	if dist < heap[0].distance {
	// 		heap[0].distance = dist
	// 		heap[0].index = i
	// 		heapify(heap, 0)
	// 	}

	// }
	// for _, val := range heap {
	// 	result = append(result, points[val.index])
	// }

	// return result
	maxHeap := &MaxHeap{}

	for i := 0; i < len(points); i++ {
		dist := findDist(points[i])
		points[i] = append(points[i], dist)
		heap.Push(maxHeap, points[i])

		if maxHeap.Len() > k {
			heap.Pop(maxHeap)
		}
	}

	res := [][]int{}
	for maxHeap.Len() > 0 {
		point := heap.Pop(maxHeap).([]int)
		res = append(res, point[:2])
	}
	return res

}

type MaxHeap [][]int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i][2] > h[j][2] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func findDist(points []int) int {
	return (points[0]*points[0] + points[1]*points[1])
}

// func buildHeap(heap []node) {
// 	for i := len(heap) / 2; i >= 0; i-- {
// 		heapify(heap, i)
// 	}
// }

// func heapify(heap []node, index int) {
// 	largest := index
// 	left, right := 2*index+1, 2*index+2
// 	if left < len(heap) && heap[left].distance > heap[largest].distance {
// 		largest = left
// 	}
// 	if right < len(heap) && heap[right].distance > heap[largest].distance {
// 		largest = right
// 	}
// 	if index != largest {
// 		heap[largest], heap[index] = heap[index], heap[largest]
// 		heapify(heap, largest)
// 	}

// }

// func getDist(ordinates []int) float64 {
// 	return math.Sqrt(math.Pow(float64(ordinates[0]-0), 2) + math.Pow(float64(ordinates[1]-0), 2))
// }
