package main

import "fmt"

func main() {

	fmt.Print(generateParenthesis(3))

}

func generateParenthesis(n int) []string {
	result := []string{}

	backTrackParenthsis(&result, "", 0, 0, n)
	return result

}

// func backTrackParenthsis(result *[]string, str string, left, right, n int) {
// 	if n == 0 {
// 		return
// 	}
// 	if len(str) == 2*n {
// 		*result = append(*result, str)
// 	}
// 	if left < n {
// 		str += "("
// 		backTrackParenthsis(result, str, left+1, right, n)
// 		str = str[:len(str)-1]
// 	}
// 	if right < left {
// 		str += ")"
// 		backTrackParenthsis(result, str, left, right+1, n)
// 		str = str[1:]
// 	}

// }

func backTrackParenthsis(result *[]string, str string, left, right, n int) {

	if n == 0 {
		return
	}

	if len(str) == 2*n {
		*result = append(*result, str)
	}

	if left < n {
		str += "("
		backTrackParenthsis(result, str, left+1, right, n)
		str = str[:len(str)-1]
	}

	if right < left {
		str += ")"
		backTrackParenthsis(result, str, left, right+1, n)
		str = str[:len(str)-1]
	}
}
