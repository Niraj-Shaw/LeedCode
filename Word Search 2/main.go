package main

import "fmt"

func main() {
	fmt.Print(findWords([][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
		[]string{"oath", "pea", "eat", "rain"}))
	// fmt.Print(findWords([][]byte{{'a'}, {'a'}},
	// 	[]string{"aaa"}))

}

type Node struct {
	children [26]*Node
	isEnd    bool
}
type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	return &Trie{root: &Node{}}
}

func (t *Trie) Insert(w string) {
	curNode := t.root
	for i := 0; i < len(w); i++ {
		index := w[i] - 'a'
		if curNode.children[index] == nil {
			curNode.children[index] = &Node{}
		}
		curNode = curNode.children[index]
	}
	curNode.isEnd = true
}

func findWords(board [][]byte, words []string) []string {
	res := []string{}
	// mp := make(map[string]bool)
	// for _, word := range words {
	// 	mp[word] = true
	// }
	// var backTrack func(row, col int, list string)
	// backTrack = func(row, col int, curWord string) {
	// 	if len(curWord) == 10 {
	// 		curWord = ""
	// 		return
	// 	}
	// 	if row < 0 || row > len(board)-1 || col < 0 || col > len(board[0])-1 {
	// 		curWord = ""
	// 		return
	// 	}
	// 	if board[row][col] == 0 {
	// 		return
	// 	}

	// 	curWord += string(board[row][col])
	// 	if _, ok := mp[curWord]; ok {
	// 		res = append(res, curWord)
	// 		delete(mp, curWord)
	// 	}
	// 	temp := board[row][col]
	// 	board[row][col] = 0
	// 	backTrack(row+1, col, curWord)
	// 	backTrack(row-1, col, curWord)
	// 	backTrack(row, col+1, curWord)
	// 	backTrack(row, col-1, curWord)
	// 	board[row][col] = temp

	// }
	// for i := 0; i < len(board); i++ {
	// 	curWord := ""
	// 	for j := 0; j < len(board[0]); j++ {
	// 		backTrack(i, j, curWord)
	// 	}
	// }
	Tr := InitTrie()
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, word := range words {
		Tr.Insert(word)
	}
	var backTrack func(i, j int, parentTrie *Node, word string)
	backTrack = func(i, j int, parentTrie *Node, word string) {
		if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 {
			return
		}
		if board[i][j] == 0 {
			return
		}
		ch := board[i][j]
		word = word + string(ch)
		if parentTrie.children[ch-'a'] == nil {
			return
		}
		if parentTrie.children[ch-'a'].isEnd {
			res = append(res, word)
			parentTrie.children[ch-'a'].isEnd = false
			parentTrie.children[ch-'a'] = nil
			parentTrie = nil
			return
		}
		temp := board[i][j]
		board[i][j] = 0
		for _, dir := range dirs {
			backTrack(i+dir[0], j+dir[1], parentTrie.children[ch-'a'], word)
		}
		board[i][j] = temp

	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			backTrack(i, j, Tr.root, "")
		}
	}
	return res

}
