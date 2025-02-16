package main

import "fmt"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func main() {
	node9 := createNode(9, nil, nil)
	node8 := createNode(5, nil, nil)
	node7 := createNode(3, nil, nil)
	node6 := createNode(7, nil, nil)
	node5 := createNode(4, node7, node8)
	node4 := createNode(0, nil, nil)
	node3 := createNode(8, node6, node9)
	node2 := createNode(2, node4, node5)
	node1 := createNode(6, node2, node3)
	fmt.Print(lowestCommonAncestor(node1, node2, node3))

}

func createNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// if p.Val < root.Val && q.Val < root.Val {
	// 	return lowestCommonAncestor(root.Left, p, q)
	// } else if p.Val > root.Val && q.Val > root.Val {
	// 	return lowestCommonAncestor(root.Right, p, q)

	// } else {
	// 	return root
	// }
	curr := root
	for curr != nil {

		if curr.Val < p.Val && curr.Val < q.Val {
			curr = curr.Right
		} else if curr.Val > p.Val && curr.Val > q.Val {
			curr = curr.Left
		} else {
			break
		}
	}

	if curr != q && curr != p {
		return curr
	} else if curr == p {
		return p
	} else {
		return q
	}

}
