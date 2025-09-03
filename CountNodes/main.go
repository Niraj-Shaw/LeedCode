package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node6 := &TreeNode{6, nil, nil}
	node5 := &TreeNode{5, nil, nil}
	node4 := &TreeNode{4, nil, nil}
	node3 := &TreeNode{3, node6, nil}
	node2 := &TreeNode{2, node4, node5}
	node1 := &TreeNode{1, node2, node3}
	fmt.Print(countNodes(node1))

}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
