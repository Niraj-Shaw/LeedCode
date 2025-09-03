package main

import "sort"

type AutocompleteSystem struct {
	Root          *TrieNode
	CurrentSearch string
	currentNode   *TrieNode
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	root := NewTrieNode()
	for id, sentence := range sentences {
		addToTrie(root, sentence, times[id])
	}
	return AutocompleteSystem{
		Root:        root,
		currentNode: root,
	}
}

func (this *AutocompleteSystem) Input(c byte) []string {
	result := []string{}
	node := this.currentNode

	if c == '#' {
		addToTrie(this.Root, this.CurrentSearch, 1)
		this.CurrentSearch = ""
		this.currentNode = this.Root
		return result
	}

	this.CurrentSearch += string(c)
	if _, ok := node.childeren[c]; !ok {
		this.currentNode.childeren[c] = NewTrieNode()
		this.currentNode = this.currentNode.childeren[c]
		return result
	}

	this.currentNode = this.currentNode.childeren[c]

	nodeSentences := this.currentNode.Sentences
	for key := range nodeSentences {
		result = append(result, key)
	}

	sort.Slice(result, func(i, j int) bool {
		if nodeSentences[result[i]] != nodeSentences[result[j]] {
			return nodeSentences[result[i]] > nodeSentences[result[j]]
		}
		// Secondary: Ascending ASCII order
		return result[i] < result[j]
	})

	if len(result) > 3 {
		result = result[:3]
	}

	return result

}

func addToTrie(root *TrieNode, sentence string, count int) {
	for _, letter := range sentence {
		l := byte(letter)
		if _, ok := root.childeren[l]; !ok {
			root.childeren[l] = NewTrieNode()
		}

		root = root.childeren[l]
		root.Sentences[sentence] += count
	}
}
