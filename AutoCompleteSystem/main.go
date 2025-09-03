package main

import "fmt"

func main() {
	autocompleteSystem := Constructor([]string{"i love you", "island", "iroman", "i love leetcode"},
		[]int{5, 3, 2, 2})
	fmt.Print(autocompleteSystem.Input('i'))
	fmt.Print(autocompleteSystem.Input(' '))
	fmt.Print(autocompleteSystem.Input('a'))
	fmt.Print(autocompleteSystem.Input('#'))
	fmt.Print(autocompleteSystem.Input('i'))
	fmt.Print(autocompleteSystem.Input(' '))
	fmt.Print(autocompleteSystem.Input('a'))
	fmt.Print(autocompleteSystem.Input('#'))
	fmt.Print(autocompleteSystem.Input('i'))
	fmt.Print(autocompleteSystem.Input(' '))
	fmt.Print(autocompleteSystem.Input('a'))
	fmt.Print(autocompleteSystem.Input('#'))
}
