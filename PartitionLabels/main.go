package main

import (
	"fmt"
)

func main() {
	fmt.Print(partitionLabels("ababcbacadefegdehijhklij"))
}

func partitionLabels(s string) []int {
	result := []int{}
	dp := make([]int, 26)
	for i, ch := range s {
		dp[ch-'a'] = i
	}
	start, end := 0, 0
	for i, ch := range s {
		if dp[ch-'a'] > end {
			end = dp[ch-'a']
		}
		if i == end {
			result = append(result, end-start+1)
			start = end + 1
		}
	}
	return result
}
