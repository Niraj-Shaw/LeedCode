package main

import (
	"fmt"
	"sort"
	"strings"
)

type SourceTargetMapping struct {
	source string
	target string
}

func main() {
	fmt.Println(findReplaceString("abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}))
	fmt.Println(findReplaceString("abcd", []int{0, 2}, []string{"ab", "ec"}, []string{"eee", "ffff"}))
	fmt.Println(findReplaceString("vmokgggqzp", []int{3, 5, 1}, []string{"kg", "ggq", "mo"}, []string{"s", "so", "bfr"}))
	fmt.Println(findReplaceString("abcde", []int{2, 2, 3}, []string{"cde", "cdef", "dk"}, []string{"fe", "f", "xyz"}))
	fmt.Println(findReplaceString("abcde", []int{2, 2}, []string{"cdef", "bc"}, []string{"f", "fe"}))
	fmt.Println(findReplaceString("arpbtanyiasrgychmlmofjoaoagfantiwozheuvqrviygudcnmujoxgwlvcxysnmxxvbczdgubkavlcteeglimfjquajmrllyxmp",
		[]int{75, 97, 70, 24, 72, 27, 52, 44, 9, 36, 91, 67, 95, 33, 42, 29, 49, 18, 0, 64, 3, 86, 20, 39, 77, 56, 59, 46, 12, 93, 88, 5, 31},
		[]string{"av", "xm", "dg", "oag", "ub", "fa", "oxgw", "gu", "as", "euv", "jm", "bc", "ly", "ozh", "iy", "nt", "mu", "mo", "arp", "xxv", "bt", "fj", "fj", "qrv", "lctee", "lvc", "xysn", "dc", "gy", "rl", "qu", "anyi", "iw"},
		[]string{"ldl", "ud", "eg", "ltz", "wcq", "qu", "qxwd", "v", "aos", "ykus", "nx", "erg", "sxb", "flx", "fjy", "ld", "ylb", "uca", "ug", "vkjm", "x", "qe", "tv", "gjay", "lokks", "fnr", "wdq", "j", "dgi", "pvt", "oqg", "bmm", "d"}))
	fmt.Println(findReplaceString("abcd", []int{0, 0}, []string{"a", "b"}, []string{"b", "c"}))
}

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	if len(s) == 0 {
		return ""
	}
	if len(indices) != len(sources) {
		return ""
	}
	if len(indices) != len(targets) {
		return ""
	}

	result := []string{}
	i := 0
	mp := make(map[int][]SourceTargetMapping)
	matched := false

	for idx, indice := range indices {
		mp[indice] = append(mp[indice], SourceTargetMapping{
			source: sources[idx],
			target: targets[idx],
		})
	}

	sort.Ints(indices)

	prev := -1
	for _, index := range indices {
		if index == prev {
			continue
		}
		matched = false

		for _, record := range mp[index] {

			if i >= len(s) {
				break
			}
			if index == prev && matched {
				break
			}

			if index > i {
				for i < index {
					result = append(result, string(s[i]))
					i++
				}
			}

			srcStr := record.source

			for j := 0; i < index+len(srcStr) && i < len(s) && j < len(srcStr); i, j = i+1, j+1 {
				if s[i] != srcStr[j] {
					break
				}
			}

			if i == index+len(srcStr) {
				matched = true
			}

			if matched {
				result = append(result, record.target)
			} else {
				i = index
			}

			prev = index
		}
	}

	for i < len(s) {
		result = append(result, string(s[i]))
		i++
	}

	return strings.Join(result, "")
}
