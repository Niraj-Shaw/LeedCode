package main

import "fmt"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func main() {
	//node9 := createNode(1, nil, nil)
	//node8 := createNode(4, nil, node9)
	//node7 := createNode(13, nil, nil)
	//	node6 := createNode(4, nil, nil)
	node5 := createNode(3, nil, nil)
	node4 := createNode(1, nil, nil)
	node3 := createNode(7, nil, nil)
	node2 := createNode(2, node4, node5)
	node1 := createNode(4, node2, node3)
	fmt.Print(insertIntoBST(node1, 5))

}

func createNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}

}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	newNode := TreeNode{Val: val, Left: nil, Right: nil}
	var PrevNode *TreeNode
	curr := root
	if root == nil {
		root = &newNode
	}
	// if val < root.Val {
	// 	root.Left = insertIntoBST(root.Left, val)
	// } else if val > root.Val {
	// 	root.Right = insertIntoBST(root.Right, val)

	// }
	// return root
	for curr != nil || PrevNode != nil {
		if curr != nil {
			PrevNode = curr
			if val < curr.Val {
				curr = curr.Left
			} else if val > curr.Val {
				curr = curr.Right
			}
		} else {
			if PrevNode.Val < val {
				PrevNode.Right = &newNode

			} else {
				PrevNode.Left = &newNode
			}
			PrevNode = nil
		}

	}
	return root
}
