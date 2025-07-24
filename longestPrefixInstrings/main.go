package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLongestPrefix([]string{""})) // Expected: ""
}

// findLongestPrefix finds the longest common prefix among a list of strings
func findLongestPrefix(lst []string) string {
	if len(lst) == 0 {
		return ""
	}

	sort.Strings(lst) // Sorting helps get lexicographically smallest & largest words

	first, last := lst[0], lst[len(lst)-1]
	minLen := min(len(first), len(last))

	for i := 0; i < minLen; i++ {
		if first[i] != last[i] {
			return first[:i]
		}
	}
	return first[:minLen]
}
