package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Node5 := &TreeNode{5, nil, nil}
	// Node4 := &TreeNode{4, nil, nil}
	// Node3 := &TreeNode{3, nil, nil}
	Node2 := &TreeNode{2, nil, nil}
	Node1 := &TreeNode{1, Node2, nil}
	fmt.Print(diameterOfBinaryTree(Node1))
}

func diameterOfBinaryTree(root *TreeNode) int {
	var diameter int
	dfs(root, &diameter)
	return diameter

}

func dfs(node *TreeNode, diameter *int) int {
	if node == nil {
		return 0
	}
	left_diameter := dfs(node.Left, diameter)
	right_diameter := dfs(node.Right, diameter)
	*diameter = max(*diameter, left_diameter+right_diameter)
	return max(left_diameter, right_diameter) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
