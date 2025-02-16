package main

import "fmt"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func main() {
	// node9 := createNode(1, nil, nil)
	// node8 := createNode(4, nil, node9)
	// node7 := createNode(13, nil, nil)
	node6 := createNode(7, nil, nil)
	node5 := createNode(4, nil, nil)
	node4 := createNode(2, nil, nil)
	node3 := createNode(6, nil, node6)
	node2 := createNode(3, node4, node5)
	node1 := createNode(5, node2, node3)
	fmt.Print(findTarget(node1, 9))

}

func createNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}

}

func findTarget(root *TreeNode, k int) bool {
	myMap := make(map[int]bool)
	// 	// stack := []*TreeNode{root}
	// 	// for len(stack) != 0 {
	// 	// 	root = stack[len(stack)-1]
	// 	// 	stack = stack[:len(stack)-1]

	// 	// 	if _, Ok := mymap[k-root.Val]; Ok {
	// 	// 		return true

	// 	// 	} else {
	// 	// 		mymap[root.Val] = true
	// 	// 	}

	// 	// 	if root.Right != nil {
	// 	// 		stack = append(stack, root.Right)
	// 	// 	}
	// 	// 	if root.Left != nil {
	// 	// 		stack = append(stack, root.Left)
	// 	//	}
	// 	//return false

	// 	//}
	return find(root, k, myMap)

}

func find(root *TreeNode, k int, myMap map[int]bool) bool {
	if root == nil {
		return false
	}
	if myMap[k-root.Val] {
		return true
	}
	myMap[root.Val] = true

	return find(root.Left, k, myMap) || find(root.Right, k, myMap)

}
