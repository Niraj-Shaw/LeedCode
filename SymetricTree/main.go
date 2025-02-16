package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node7 := createNode(3, nil, nil)
	node6 := createNode(4, nil, nil)
	node5 := createNode(4, nil, nil)
	node4 := createNode(3, nil, nil)
	node3 := createNode(2, node6, node7)
	node2 := createNode(2, node4, node5)
	node1 := createNode(1, node2, node3)
	fmt.Print(isSymmetric(node1))

}

func createNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}

}
func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{root, root}

	if root == nil {
		return false
	}
	for len(queue) != 0 {
		left, right := queue[0], queue[1]
		if left.Val != right.Val {
			return false
		}
		if left.Left != nil || right.Right != nil {
			if left.Left != nil && right.Right != nil {
				queue = append(queue, left.Left)
				queue = append(queue, right.Right)
			} else {
				return false
			}
		}
		if left.Right != nil || right.Left != nil {
			if left.Right != nil && right.Left != nil {
				queue = append(queue, left.Right)
				queue = append(queue, right.Left)
			} else {
				return false
			}
		}
		queue = queue[2:]

	}
	return true
}

// func isMirror(n1, n2 *TreeNode) bool {
// 	if n1 == nil && n2 == nil {
// 		return true
// 	}
// 	if n1 == nil || n2 == nil {
// 		return false
// 	}

// 	return n1.Val == n2.Val &&
// 		isMirror(n1.Left, n2.Right) &&
// 		isMirror(n1.Right, n2.Left)
// }
