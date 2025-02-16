package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	// mp := make(map[rune]int)
	// mp_t := make(map[rune]int)
	// min := math.MaxInt64
	// left_min := 0
	// right_min := 0

	// for _, value := range t {
	// 	mp[value]++
	// }

	// left, right := 0, 0
	// n := 0

	// for right < len(s) {
	// 	char := rune(s[right])
	// 	if val, ok := mp[char]; ok {
	// 		mp_t[char]++
	// 		if mp_t[char] <= val {
	// 			n++
	// 		}
	// 	}

	// 	for n == len(t) {
	// 		if right-left+1 < min {
	// 			min = right - left + 1
	// 			left_min = left
	// 			right_min = right
	// 		}

	// 		left_char := rune(s[left])
	// 		if _, ok := mp[left_char]; ok {
	// 			if mp_t[left_char] == mp[left_char] {
	// 				n--
	// 			}
	// 			mp_t[left_char]--
	// 		}
	// 		left++
	// 	}
	// 	right++
	// }

	// if min == math.MaxInt64 {
	// 	return ""
	// }
	// return s[left_min : right_min+1]
	mp, mp_t := make(map[string]int), make(map[string]int)
	left, right := 0, 0
	min := math.MaxInt64
	left_min, right_min := 0, 0
	n := 0
	for _, val := range t {
		mp[string(val)]++
	}
	for right < len(s) {
		char := string(s[right])
		if _, ok := mp[char]; ok {
			mp_t[char]++
			if mp_t[char] <= mp[char] {
				n++
			}
		}
		for n == len(t) {
			if right-left+1 < min {
				min = right - left + 1
				left_min = left
				right_min = right
			}
			left_char := string(s[left])
			if _, ok := mp_t[left_char]; ok {
				if mp[left_char] == mp_t[left_char] {
					n--
				}
				mp_t[left_char]--
			}
			left++
		}
		right++
	}
	if min == math.MaxInt64 {
		return ""
	}
	return s[left_min : right_min+1]
}
