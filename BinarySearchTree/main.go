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
	//node6 := createNode(2, nil, nil)
	node5 := createNode(3, nil, nil)
	node4 := createNode(1, nil, nil)
	node3 := createNode(7, nil, nil)
	node2 := createNode(2, node4, node5)
	node1 := createNode(4, node2, node3)
	fmt.Print(searchBST(node1, 8))

}

func createNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}

}
func searchBST(root *TreeNode, val int) *TreeNode {
	// if root == nil {
	// 	return nil
	// }
	// if val == root.Val {
	// 	return root
	// }

	// if val < root.Val {
	// 	return searchBST(root.Left, val)

	// } else {
	// 	return searchBST(root.Right, val)
	// }
	curr := root
	for curr != nil {
		if curr.Val == val {
			return curr
		}
		if curr.Left != nil && curr.Val > val {
			curr = curr.Left
		} else if curr.Right != nil && curr.Val < val {
			curr = curr.Right
		} else {
			curr = nil
		}
	}
	return nil

}
