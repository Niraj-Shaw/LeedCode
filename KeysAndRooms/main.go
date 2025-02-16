package main

import "fmt"

func main() {
	fmt.Print(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
}

func canVisitAllRooms(rooms [][]int) bool {
	queue, seen := []int{0}, map[int]bool{0: true}
	for len(queue) != 0 {
		curr := queue[0]
		keys := rooms[curr]
		for _, val := range keys {
			if !seen[val] {
				queue = append(queue, val)
				seen[val] = true
			}

		}
		queue = queue[1:]

	}
	if len(seen) == len(rooms) {
		return true
	}
	return false

}
