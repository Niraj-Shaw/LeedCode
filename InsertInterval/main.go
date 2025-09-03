package main

import "fmt"

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 9}))
	fmt.Println(insert([][]int{}, []int{5, 7}))
	fmt.Println(insert([][]int{{6, 8}}, []int{1, 2}))
	fmt.Println(insert([][]int{{6, 8}}, []int{10, 12}))
	fmt.Println(insert([][]int{{1, 5}}, []int{1, 7}))
	fmt.Println(insert([][]int{{2, 4}, {5, 7}, {8, 10}, {11, 13}}, []int{3, 6}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	i := 0
	res := [][]int{}

	for i < n && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
		i++
	}
	res = append(res, newInterval)

	for i < n {
		res = append(res, intervals[i])
		i++
	}

	return res
}
