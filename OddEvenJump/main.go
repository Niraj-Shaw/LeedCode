package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(oddEvenJumps([]int{10, 13, 12, 14, 15}))
}

func oddEvenJumps(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return 1
	}

	type pair struct {
		idx int
		val int
	}

	dpOdd := make(map[int]int)
	dpEven := make(map[int]int)

	var backtrackFn func(index, jumpNum int) int
	backtrackFn = func(index, jumpNum int) int {
		if index == -1 || index == len(arr)-1 {
			return index
		}

		if jumpNum%2 == 0 {
			if val, ok := dpEven[index]; ok {
				if val == index {
					return index
				}

				return backtrackFn(val, jumpNum+1)
			}
		} else {
			if val, ok := dpOdd[index]; ok {
				if val == index {
					return index
				}

				return backtrackFn(val, jumpNum+1)
			}
		}

		var tempArr []pair
		nextIndex := -1

		if jumpNum%2 == 0 {

			for i := index + 1; i < len(arr); i++ {
				if arr[i] <= arr[index] {
					tempArr = append(tempArr, pair{
						idx: i,
						val: arr[i],
					})
				}
			}

			sort.Slice(tempArr, func(i, j int) bool {
				if tempArr[i].val == tempArr[j].val {
					return tempArr[i].idx < tempArr[j].idx
				}
				return tempArr[i].val > tempArr[j].val
			})

		} else {

			for i := index + 1; i < len(arr); i++ {
				if arr[i] >= arr[index] {
					tempArr = append(tempArr, pair{
						idx: i,
						val: arr[i],
					})
				}
			}

			sort.Slice(tempArr, func(i, j int) bool {
				if tempArr[i].val == tempArr[j].val {
					return tempArr[i].idx < tempArr[j].idx
				}
				return tempArr[i].val < tempArr[j].val
			})
		}

		if len(tempArr) != 0 {
			nextIndex = tempArr[0].idx
			if jumpNum%2 == 0 {
				dpEven[index] = nextIndex
			} else {
				dpOdd[index] = nextIndex
			}

			return backtrackFn(nextIndex, jumpNum+1)
		} else {
			if jumpNum%2 == 0 {
				dpEven[index] = index
			} else {
				dpOdd[index] = index
			}

			return index
		}

	}
	count := 0
	for i := range arr {
		endIndex := backtrackFn(i, 1)

		if endIndex == len(arr)-1 {
			count++
		}
	}
	return count

}
