package main

import (
	"strings"
)

func main() {
	print(mostCommonWord("a.", []string{""}))

}

func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.Trim(strings.ToLower(paragraph), " ")
	word := ""
	mp := make(map[string]int)
	for index, ch := range paragraph {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			word = word + string(ch)
		} else if (ch == ' ') || (ch == ',') || (index == len(paragraph)-1) {
			if word != "" {
				mp[word]++
				word = ""
			}
		}
	}

	count := 0
	result := ""
	for _, val := range banned {
		delete(mp, val)
	}
	for k, val := range mp {
		if val > count {
			count = val
			result = k

		}
	}
	return result

}
