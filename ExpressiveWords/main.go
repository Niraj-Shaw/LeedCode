package main

import (
	"fmt"
)

func main() {
	fmt.Println(expressiveWords("heeellooo", []string{"hello", "hi", "helo"}))
	fmt.Println(expressiveWords("zzzzzyyyyy", []string{"zzyy", "zy", "zyy"}))
	fmt.Println(expressiveWords("lee", []string{"le"}))
	fmt.Println(expressiveWords("heeellooo", []string{"helloooworld"}))
	fmt.Println(expressiveWords("heeellooo", []string{"hellooooworld"}))
	fmt.Println(expressiveWords("nnnnnnzzzznnnnnnpppppfffff", []string{"nzznpf", "nnzznnpff", "nznpff", "nnznnpf", "nnznpff", "nzznppf", "nzznpff", "nnzznnppf"}))
	fmt.Println(expressiveWords("abcd", []string{"abc"}))
	fmt.Println(expressiveWords("aaa", []string{"aaaa"}))
}

func expressiveWords(s string, words []string) int {
	result := 0

	for _, word := range words {
		lenS, lenW := 0, 0
		cntS, cntW := 0, 0
		curr := s[0]

		for lenS < len(s) || lenW < len(word) {
			for lenS < len(s) && s[lenS] == curr {
				lenS++
				cntS++
			}

			for lenW < len(word) && word[lenW] == curr {
				lenW++
				cntW++
			}

			if cntS == 0 || cntW == 0 || cntW > cntS || (cntS != cntW && cntS <= 2) {
				break
			}

			if lenS == len(s) && lenW == len(word) {
				result++
			}

			cntS, cntW = 0, 0

			if lenS < len(s) {
				curr = s[lenS]
			}
		}
	}
	return result
}
