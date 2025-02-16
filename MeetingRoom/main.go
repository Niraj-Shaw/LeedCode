package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}))
}

func minMeetingRooms(intervals [][]int) int {
	// sort.SliceStable(intervals, func(i, j int) bool {
	// 	return intervals[i][0] < intervals[j][0]
	// })

	// h := &Heap.IntHeap{intervals[0][1]}
	// Heap.InitHeap(h)
	// for _, meeting := range intervals[1:] {
	// 	if meeting[0] > (*h)[0] {
	// 		Heap.PopHeap(h)
	// 	}
	// 	Heap.PushHeap(h, meeting[1])

	// }
	// return h.Len()
	start_arr := []int{}
	end_arr := []int{}
	room_count := 0
	for _, row := range intervals {
		start_arr = append(start_arr, row[0])
		end_arr = append(end_arr, row[1])
	}
	sort.Ints(start_arr)
	sort.Ints(end_arr)
	start, end := 0, 0
	for start < len(start_arr) {
		if start_arr[start] < end_arr[end] {
			room_count++
		} else {
			end++
		}
		start++
	}
	return room_count

}
