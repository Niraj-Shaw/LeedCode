// Given a string, find the longest consecutive substring which consists of only digits.
// e.g.
// Input: a1b2c34de
// Output: 34

package main

import (
	"fmt"
	"unicode"
)

var mp map[string]string

func main() {

	fmt.Println(findLongestDigitSubstring("a1b2c34de"))
	fmt.Println(findLongestDigitSubstring(""))
	fmt.Println(findLongestDigitSubstring("babababa"))
	fmt.Println(findLongestDigitSubstring("1111111111"))
	fmt.Println(findLongestDigitSubstring("111111a222222"))
}

func findLongestDigitSubstring(str string) string {
	mp := make(map[string]string)

	if val, exists := mp[str]; exists{
		return val
	}

	start, end := 0, 0
	longest := ""

	for end < len(str) {
		curr := -1
		if unicode.IsDigit(rune(str[end])) {
			end++
		} else {

			end++
			start = end
		}

		curr = end - start
		if curr > len(longest) {
			longest = str[start:end]
		}

	}
	mp[str] = longest
	return longest
}

func execute(key string) string{
	if value, exist := cache.mp[key]; exist{
		return value
	}
	perform task to evaluate the value

	mp[key] = value

	return value
}

