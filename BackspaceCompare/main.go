package main

import "fmt"

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
	fmt.Println(backspaceCompare("ab##", "c#d#"))
	fmt.Println(backspaceCompare("a#c", "b"))
	fmt.Println(backspaceCompare("a##c", "#a#c"))
	fmt.Println(backspaceCompare("a##", "#bac#"))
}

func backspaceCompare(s string, t string) bool {
	s = replaceCharsWithBackspace(s)
	t = replaceCharsWithBackspace(t)

	return s == t
}

func replaceCharsWithBackspace(str string) string {
	i := 0

	for i < len(str) {
		if str[i] == '#' {
			tmp := i
			for j := tmp; j >= tmp-1 && j >= 0; j-- {
				str = str[:j] + str[j+1:]
				i--
			}
		}

		i++
	}

	return str
}
