package main

import "fmt"

func main() {
	fmt.Print(strStr("leetcode", "leeto"))

}

func strStr(haystack string, needle string) int {
	// slow, fast, i, n := 0, 0, 0, len(heystack)
	// for slow < n {
	// 	if n-slow+1 < len(needle) {
	// 		break
	// 	}

	// 	for i < len(needle) && fast < n && heystack[fast] == needle[i] {
	// 		fast++
	// 		i++
	// 	}
	// 	if i == len(needle) {
	// 		return slow
	// 	} else {
	// 		slow++
	// 		fast = slow
	// 		i = 0
	// 	}

	// }
	// return -1
	l1, l2 := len(haystack)-1, len(needle)-1
	start, end := 0, l2
	for end <= l1 {
		if haystack[start:end+1] == needle {
			return start
		} else {
			start++
			end++
		}
	}
	return -1

}
