package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node7 := createNode(9, nil, nil)
	node6 := createNode(6, nil, nil)
	node5 := createNode(3, nil, nil)
	node4 := createNode(1, nil, nil)
	node3 := createNode(7, node6, node7)
	node2 := createNode(2, node4, node5)
	node1 := createNode(4, node2, node3)
	fmt.Printf("%v", invertTree(node1))

}

func createNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}

}

func invertTree(root *TreeNode) *TreeNode {
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		curr := queue[0]
		if curr != nil {
			temp := curr.Right
			curr.Right = curr.Left
			curr.Left = temp
			queue = append(queue, curr.Left, curr.Right)
		}
		queue = queue[1:]

	}
	return root
}

// func invertTree(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	left := invertTree(root.Left)
// 	right := invertTree(root.Right)
// 	root.Left = right
// 	root.Right = left
// 	return root

// }
