package main

import "fmt"

func main() {
	n5 := &TreeNode{5, nil, nil}
	n4 := &TreeNode{4, nil, nil}
	n3 := &TreeNode{3, n4, n5}
	n2 := &TreeNode{2, nil, nil}
	n1 := &TreeNode{1, n2, n3}
	cd := Constructor()
	fmt.Print(cd.serialize(n1))
}
