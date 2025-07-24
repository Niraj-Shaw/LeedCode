package main

import (
	"sort"
	"strings"
)

func main() {
	reorderLogFiles([]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"})
}

func reorderLogFiles(logs []string) []string {
	letterLogs := []string{}
	digitLogs := []string{}

	for _, v := range logs {
		Index := strings.Index(v, " ")
		if Index+1 < len(v) && (v[Index+1] >= 'a' && v[Index+1] <= 'z') {
			letterLogs = append(letterLogs, v)
		} else {
			digitLogs = append(digitLogs, v)
		}
	}

	sort.Slice(letterLogs, func(i, j int) bool {
		iIndex := strings.Index(letterLogs[i], " ")
		jIndex := strings.Index(letterLogs[j], " ")

		istr := letterLogs[i][iIndex+1:]
		jstr := letterLogs[j][jIndex+1:]

		if istr == jstr {
			return letterLogs[i] < letterLogs[j]
		}
		return istr < jstr
	})

	letterLogs = append(letterLogs, digitLogs...)

	return letterLogs

}
