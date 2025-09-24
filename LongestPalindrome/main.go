package main

import "fmt"

func main() {
	str := longestPalindrome("babad")
	//str := longestPalindrome("cbbd")
	fmt.Print(str)
}

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	ans := [2]int{}

	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true

	}
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			ans[0] = i
			ans[1] = i + 1

		}
	}
	for i := 2; i < n; i++ {
		for j := 0; j < n-i; j++ {
			k := j + i
			if s[j] == s[k] && dp[j+1][k-1] {
				dp[j][k] = true
				ans[0] = j
				ans[1] = k

			}

		}
	}
	i := ans[0]
	j := ans[1]
	return s[i : j+1]

}

// func longestPalindrome(s string) string {
// 	n := len(s)
// 	left, right, maxLength := 0, 0, 0

// 	for i := 0; i < n; i++ {
// 		lenOdd := expandAroundTheCenter(s, i, i)
// 		lenEven := expandAroundTheCenter(s, i, i+1)

// 		if lenEven > lenOdd {
// 			maxLength = lenEven
// 		} else {
// 			maxLength = lenOdd
// 		}

// 		if maxLength > right-left {
// 			left = i - (maxLength-1)/2
// 			right = i + maxLength/2
// 		}
// 	}

// 	return s[left : right+1]

// }

// func expandAroundTheCenter(s string, left, right int) int {
// 	if left < 0 || s == "" {
// 		return 0
// 	}

// 	for left >= 0 && right < len(s) && s[left] == s[right] {
// 		left--
// 		right++
// 	}

// 	return right - left - 1
// }
