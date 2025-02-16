package main

import "fmt"

func main() {
	fmt.Print(isValid("()[]{}"))

}

func isValid(s string) bool {
	myMap := map[rune]rune{'}': '{', ')': '(', ']': '['}
	stack := []rune{}

	if len(s) <= 0 {
		return false
	}
	for _, val := range s {
		if _, ok := myMap[val]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == myMap[val] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		} else {
			stack = append(stack, val)
		}

	}
	if len(stack) == 0 {
		return true
	}
	return false

}
