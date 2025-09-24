package main

import "fmt"

// type Trie struct {
// 	root *Node
// }

// type Node struct {
// 	next map[rune]*Node
// }

// // Insert
// func (t *Trie) Insert(w string) {
// 	currNode := t.root
// 	for _, ch := range w {
// 		if _, ok := currNode.next[ch]; !ok {
// 			currNode.next[ch] = &Node{next: make(map[rune]*Node)}
// 		}
// 		currNode = currNode.next[ch]
// 	}
// 	currNode.next['*'] = nil

// }

// // Search
// func (t *Trie) Search(w string) bool {
// 	currNode := t.root
// 	for _, ch := range w {
// 		if _, ok := currNode.next[ch]; !ok {
// 			return false
// 		}
// 		currNode = currNode.next[ch]

// 	}
// 	if _, ok := currNode.next['*']; ok {
// 		return true
// 	}
// 	return false
// }

// // Init

// func InitTrie() *Trie {
// 	return &Trie{root: &Node{next: make(map[rune]*Node)}}
// }

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	next map[rune]*TrieNode
}

func InitTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			next: make(map[rune]*TrieNode),
		},
	}
}

func (t *Trie) Insert(str string) {
	r := t.root

	for _, ch := range str {
		if _, ok := r.next[ch]; !ok {
			r.next[ch] = &TrieNode{next: make(map[rune]*TrieNode)}
		}

		r = r.next[ch]
	}
	r.next['*'] = nil
}

func main() {
	Tr := InitTrie()
	Tr.Insert("bat")
	fmt.Println(Tr)
	fmt.Println(Tr.Search("batter"))
	fmt.Println(Tr.Search("bat"))

}
