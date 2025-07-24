package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(str string) int {
	// if len(str) == 0 {
	// 	return 0
	// }

	// i := 0
	// start := 0
	// end := 0
	// max := 0
	// mp := make(map[byte]int)

	// for i < len(str) {
	// 	val, ok := mp[str[i]]
	// 	if (ok && val < start) || !ok {
	// 		mp[str[i]] = i
	// 		end++

	// 	} else {
	// 		start = val + 1
	// 	}
	// 	if max < (end - start) {
	// 		max = end - start
	// 	}
	// 	i++
	// }

	// return max
	res := 0
	left, right := 0, 0
	mp := make(map[byte]int)

	for right < len(str) {
		if val, ok := mp[str[right]]; ok && val >= left && val <= right {
			left = mp[str[right]] + 1
		}

		if res < (right - left + 1) {
			res = (right - left + 1)
		}

		mp[str[right]] = right
		right++

	}
	return res

}
