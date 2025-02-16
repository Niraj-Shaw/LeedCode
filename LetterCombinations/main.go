package main

import "fmt"

var digitMapping = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func main() {
	result := letterCombinations("234")
	fmt.Print(result)

}

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return nil
	}
	res := []string{}
	first := string(digits[0])
	rest := digits[1:]
	letters := digitMapping[first]

	if len(rest) == 0 {
		for _, str := range letters {
			res = append(res, string(str))
		}
		return res
	}
	combinations := letterCombinations(rest)
	for _, str := range letters {
		for _, str2 := range combinations {
			res = append(res, string(str)+str2)
		}
	}
	return res

}
