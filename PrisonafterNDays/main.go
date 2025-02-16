package main

import "fmt"

func main() {
	cells := []int{1, 0, 0, 1, 0, 0, 1, 0}
	fmt.Print(prisonAfterNDays(cells, 1000000000))

}

func prisonAfterNDays(cells []int, n int) []int {
	seen := make(map[string]int)
	for n > 0 {
		currentState := fmt.Sprint(cells)
		if val, ok := seen[currentState]; ok {
			n %= val - n
		} else {
			seen[currentState] = n
		}
		if n > 0 {
			n--
			cells = nextDay(cells)
		}
	}
	return cells
}

func nextDay(cells []int) []int {
	temp := []int{}
	for i := 0; i < len(cells); i++ {
		if i == 0 || i == len(cells)-1 {
			temp = append(temp, 0)
			continue
		}
		if cells[i-1] == cells[i+1] {
			temp = append(temp, 1)
		} else {
			temp = append(temp, 0)
		}
	}
	return temp

}
