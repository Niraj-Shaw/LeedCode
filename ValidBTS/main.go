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
	node5 := createNode(6, nil, nil)
	node4 := createNode(3, nil, nil)
	node3 := createNode(4, node4, node5)
	node2 := createNode(1, nil, nil)
	node1 := createNode(5, node2, node3)
	fmt.Print(isValidBST(node1))

}

func createNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}

}
func isValidBST(root *TreeNode) bool {

	// stack := []*TreeNode{}
	// curr := root
	// var prevNode *TreeNode
	// for curr != nil || len(stack) != 0 {
	// 	if curr != nil {
	// 		stack = append(stack, curr)
	// 		curr = curr.Left
	// 	} else {
	// 		curr = stack[len(stack)-1]
	// 		if prevNode != nil && curr.Val <= prevNode.Val {
	// 			return false
	// 		}
	// 		prevNode = curr
	// 		stack = stack[:len(stack)-1]
	// 		curr = curr.Right
	// 	}

	// }
	// return true
	return helperIsValidBST(root, nil, nil)

}

func helperIsValidBST(root, low, high *TreeNode) bool {
	if root == nil {
		return true
	}
	if low != nil && root.Val <= low.Val || high != nil && root.Val >= high.Val {
		return false
	}
	return helperIsValidBST(root.Left, low, root) && helperIsValidBST(root.Right, root, high)
}
