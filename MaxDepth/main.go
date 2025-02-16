package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	Height int
}

func main() {
	//node6 := createNode(6, nil, nil)
	//node5 := createNode(5, nil, nil)
	node4 := createNode(7, nil, nil)
	node3 := createNode(15, nil, nil)
	node2 := createNode(20, node3, node4)
	node7 := createNode(9, nil, nil)
	node1 := createNode(3, node7, node2)
	fmt.Print(maxDepth(node1))

}
func createNode(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	} else {
// 		L_height := maxDepth(root.Left)
// 		R_height := maxDepth(root.Right)
// 		return 1 + GetMax(L_height, R_height)
// 	}
// }

// func GetMax(value1, value2 int) int {
// 	if value1 > value2 {
// 		return value1
// 	} else {
// 		return value2
// 	}
// }

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		for i := 0; i < len(queue); i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr.Left != nil {
				queue = append(queue, curr.Left)

			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)

			}

			
		}
		depth++

	}
	return depth

}
