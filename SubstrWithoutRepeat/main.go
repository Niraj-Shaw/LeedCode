package main

import (
	"fmt"
)

func main() {
	fmt.Print(lengthOfLongestSubstring("acbabcbb"))

}

func lengthOfLongestSubstring(s string) int {
	res := 0
	left, right := 0, 0
	mp := make(map[byte]int)
	for right < len(s) {
		if val, ok := mp[s[right]]; ok && val >= left && val <= right {
			left = mp[s[right]] + 1
		}
		if (right - left + 1) > res {
			res = right - left + 1
		}
		mp[s[right]] = right
		right++
	}
	return res

}
