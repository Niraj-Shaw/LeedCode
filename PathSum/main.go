package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node9 := createNode(1, nil, nil)
	node8 := createNode(4, nil, node9)
	node7 := createNode(13, nil, nil)
	node6 := createNode(2, nil, nil)
	node5 := createNode(7, nil, nil)
	node4 := createNode(11, node5, node6)
	node3 := createNode(8, node7, node8)
	node2 := createNode(4, node4, nil)
	node1 := createNode(5, node2, node3)
	fmt.Print(hasPathSum(node1, 26))

}

func createNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}

}

func hasPathSum(root *TreeNode, targetSum int) bool {
	//return IfHasPathSum(root, targetSum, 0)

	stack := []*TreeNode{}
	var PrevNode *TreeNode
	curr := root
	Sum := 0

	for len(stack) != 0 || curr != nil {
		if curr != nil {
			Sum += curr.Val
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			if curr.Left == nil && curr.Right == nil {
				if Sum == targetSum {
					return true
				}
			}
			if curr.Right != nil && curr.Right != PrevNode {
				curr = curr.Right
			} else {
				Sum -= curr.Val
				PrevNode = curr
				stack = stack[:len(stack)-1]
				curr = nil
			}
		}
	}
	return false

}

// func IfHasPathSum(root *TreeNode, targetSum int, Sum int) bool {
// 	if root == nil {
// 		return false
// 	}
// 	Sum += root.Val
// 	if root.Left == nil && root.Right == nil {
// 		return Sum == targetSum

// 	}

// 	return IfHasPathSum(root.Left, targetSum, Sum) || IfHasPathSum(root.Right, targetSum, Sum)

// }
