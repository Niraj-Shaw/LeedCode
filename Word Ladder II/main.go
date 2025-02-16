package main

import "fmt"

func main() {
	fmt.Print(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	// ladders := [][]string{}
	// endWordFound := false
	// for _, word := range wordList {
	// 	if word == endWord {
	// 		endWordFound = true
	// 	}
	// }
	// if !endWordFound {
	// 	return ladders
	// }
	// endWordFound = false
	// visited := make(map[string]bool)
	// comb := make(map[string][]string)
	// for _, word := range wordList {
	// 	for i := range word {
	// 		temp := word[:i] + "*" + word[i+1:]
	// 		comb[temp] = append(comb[temp], word)
	// 	}
	// }
	// q := [][]string{{beginWord}}
	// for len(q) > 0 && !endWordFound {
	// 	n := len(q)
	// 	processed_List := []string{}
	// 	for n > 0 {
	// 		currList := q[0]
	// 		q = q[1:]
	// 		currWord := currList[len(currList)-1]
	// 		processed_List = append(processed_List, currWord)
	// 		for i := range currWord {
	// 			temp := currWord[:i] + "*" + currWord[i+1:]
	// 			if _, ok := comb[temp]; ok {
	// 				neighbours := comb[temp]
	// 				for _, neighbour := range neighbours {
	// 					if neighbour == endWord {
	// 						endWordFound = true
	// 					}
	// 					if !visited[neighbour] && neighbour != currWord {
	// 						lst := []string{}
	// 						lst = append(lst, currList...)
	// 						lst = append(lst, neighbour)
	// 						q = append(q, lst)
	// 					}
	// 				}
	// 			}
	// 		}
	// 		n--
	// 	}
	// 	for _, pWord := range processed_List {
	// 		visited[pWord] = true
	// 	}

	// }
	// for _, ladder := range q {
	// 	if ladder[len(ladder)-1] == endWord {
	// 		ladders = append(ladders, ladder)
	// 	}
	// }

	// return ladders
	ladders := [][]string{}
	wordLevels := make(map[string]int)
	q := []string{beginWord}
	wordLevels[beginWord] = 0
	for len(q) > 0 {
		currWord := q[0]
		q = q[1:]
		for _, word := range wordList {
			if _, haveLevel := wordLevels[word]; !haveLevel && diffChar(currWord, word) == 1 {
				q = append(q, word)
				wordLevels[word] = wordLevels[currWord] + 1
			}
		}

	}
	if _, exists := wordLevels[endWord]; !exists {
		return ladders
	}
	// var dfs func(curWord string, curList []string)
	// dfs = func(curWord string, curList []string) {
	// 	curList = append(curList, curWord)
	// 	if wordLevels[curWord] == 1 {
	// 		curList = append(curList, beginWord)
	// 		temp := curList
	// 		for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
	// 			temp[i], temp[j] = temp[j], temp[i]
	// 		}
	// 		ladders = append(ladders, temp)
	// 		curList = curList[:len(curList)-2]
	// 		return
	// 	}
	// 	for _, word := range wordList {
	// 		if _, exist := wordLevels[word]; exist && wordLevels[word]+1 == wordLevels[curWord] && diffChar(curWord, word) == 1 {
	// 			dfs(word, curList)
	// 		}

	// 	}
	// 	curList = curList[:len(curList)-1]

	// }
	// dfs(endWord, []string{})
	var dfs func(curWord string, curList []string)
	dfs = func(curWord string, curList []string) {
		curList = append(curList, curWord)
		if wordLevels[curWord] == 1 {
			curList = append(curList, beginWord)
			temp := curList
			for i, j := 0, len(curList)-1; i < j; i, j = i+1, j-1 {
				temp[i], temp[j] = temp[j], temp[i]
			}
			ladders = append(ladders, temp)
			curList = curList[:len(curList)-2]
			return
		} else {
			for _, word := range wordList {
				if _, exist := wordLevels[word]; exist && wordLevels[curWord] == wordLevels[word]+1 && diffChar(curWord, word) == 1 {
					dfs(word, curList)
				}
			}
			curList = curList[:len(curList)-1]
		}
	}
	dfs(endWord, []string{})

	return ladders
}

func diffChar(str1, str2 string) int {
	diff := 0
	for i, _ := range str1 {
		if str1[i] != str2[i] {
			diff += 1
		}
	}
	return diff

}
