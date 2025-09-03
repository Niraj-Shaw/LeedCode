package main

type TrieNode struct {
	childeren map[byte]*TrieNode
	Sentences map[string]int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		childeren: make(map[byte]*TrieNode),
		Sentences: make(map[string]int),
	}
}
