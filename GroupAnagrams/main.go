package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Print(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Print(groupAnagrams([]string{"bdddddddddd", "bbbbbbbbbbc"}))
}

func groupAnagrams(strs []string) [][]string {
	// res := [][]string{}
	// mp := make(map[string][]string)
	// for _, str := range strs {
	// 	s := []rune(str)
	// 	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	// 	mp[string(s)] = append(mp[string(s)], str)
	// }
	// for _, strs := range mp {
	// 	res = append(res, strs)
	// }
	// return res

	// for _, str := range strs {
	// 	count := [26]int{}
	// 	key := ""
	// 	for _, ch := range str {
	// 		count[ch-'a']++
	// 	}
	// 	for _, val := range count {
	// 		key += "#"
	// 		key += strconv.Itoa(val)
	// 	}
	// 	mp[key] = append(mp[key], str)
	// }
	// for _, val := range mp {
	// 	res = append(res, val)
	// }
	// return res

	dp := make(map[string][]string)
	result := [][]string{}

	for _, str := range strs {
		charCount := [26]int{}

		for _, ch := range str {
			charCount[ch-'a']++
		}

		tempStr := ""

		for _, val := range charCount {
			tempStr += "#"
			tempStr += strconv.Itoa(val)
		}

		dp[tempStr] = append(dp[tempStr], str)
	}

	for _, v := range dp {
		result = append(result, v)
	}

	return result

}
