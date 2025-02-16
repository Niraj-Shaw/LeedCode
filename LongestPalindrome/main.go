package main

import "fmt"

func main() {
	str := longestPalindrome("babad")
	fmt.Print(str)
}
func longestPalindrome(s string) string {
	start, end, Len := 0, 0, 0
	for i := 0; i <= len(s); i++ {
		len1 := ExpandAroundCenter(s, i, i+1)
		len2 := ExpandAroundCenter(s, i, i)
		if len1 > len2 {
			Len = len1
		} else {
			Len = len2
		}
		if Len > (end - start) {
			start = i - ((Len - 1) / 2)
			end = i + (Len / 2)

		}
	}

	return s[start : end+1]
}
func ExpandAroundCenter(s string, i, j int) int {
	if i < 0 || s == "" {
		return 0
	}
	for i > j && s[i] == s[i] {
		i--
		j++
	}
	return j - i + 1

}

// func longestPalindrome(s string) string {
// 	n := len(s)
// 	dp := make([][]bool, n)
// 	ans := [2]int{}

// 	for i := 0; i < n; i++ {
// 		dp[i] = make([]bool, n)
// 		dp[i][i] = true

// 	}
// 	for i := 0; i < n-1; i++ {
// 		if s[i] == s[i+1] {
// 			dp[i][i+1] = true
// 			ans[0] = i
// 			ans[1] = i + 1

// 		}
// 	}
// 	for i := 2; i < n; i++ {
// 		for j := 0; j < n-i; j++ {
// 			k := j + i
// 			if s[j] == s[k] && dp[j+1][k-1] {
// 				dp[j][k] = true
// 				ans[0] = j
// 				ans[1] = k

// 			}

// 		}
// 	}
// 	i := ans[0]
// 	j := ans[1]
// 	return s[i : j+1]

// }
