package main

import "fmt"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func main() {
	node9 := createNode(4, nil, nil)
	node8 := createNode(7, nil, nil)
	node7 := createNode(8, nil, nil)
	node6 := createNode(0, nil, nil)
	node5 := createNode(2, node8, node9)
	node4 := createNode(6, nil, nil)
	node3 := createNode(1, node6, node7)
	node2 := createNode(5, node4, node5)
	node1 := createNode(3, node2, node3)
	fmt.Print(lowestCommonAncestor(node1, node2, node9))
}

func createNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == p || root == q || root == nil {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	} else if left != nil {
		return nil
	}

	return right

}
