package main

import "fmt"

func main() {
	fmt.Print(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	endWordInList := false
	for _, word := range wordList {
		if word == endWord {
			endWordInList = true
		}
	}
	if !endWordInList {
		return 0
	}
	allComboDict := doPreprocess(wordList)
	queue := make([][]interface{}, 0, len(wordList))
	queue = append(queue, []interface{}{beginWord, 1})
	visted := make(map[string]bool)
	visted[beginWord] = true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		nodeWord := node[0].(string)
		nodeLevel := node[1].(int)
		for index, _ := range nodeWord {
			closestWord := nodeWord[:index] + string('*') + nodeWord[index+1:]
			for _, val := range allComboDict[closestWord] {
				if val == endWord {
					return nodeLevel + 1
				}
				_, ok := visted[val]
				if !ok {
					queue = append(queue, []interface{}{val, nodeLevel + 1})
					visted[val] = true
				}
			}

		}

	}
	return 0
}

func doPreprocess(wordList []string) map[string][]string {
	result := make(map[string][]string)
	for _, word := range wordList {
		for index, _ := range word {
			combo := word[:index] + string('*') + word[index+1:]
			result[combo] = append(result[combo], word)

		}
	}
	return result

}
