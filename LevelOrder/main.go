package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	myTree := []*TreeNode{}
	node5 := TreeNode{}
	node5.Val = 7
	node5.Left = nil
	node5.Right = nil
	node4 := TreeNode{}
	node4.Val = 15
	node4.Left = nil
	node4.Right = nil
	node3 := TreeNode{}
	node3.Val = 9
	node3.Left = nil
	node3.Right = nil
	node2 := TreeNode{}
	node2.Val = 20
	node2.Left = &node4
	node2.Right = &node5
	node1 := TreeNode{}
	node1.Val = 3
	node1.Left = &node3
	node1.Right = &node2
	myTree = append(myTree, &node1, &node2, &node3)
	fmt.Print(levelOrder(&node1))

}

func levelOrder(root *TreeNode) [][]int {
	// result := [][]int{}
	// if root == nil {
	// 	return result
	// }
	// arr := []int{}
	// queue := []*TreeNode{root}
	// cap := len(queue)
	// //if stack has element and pop it
	// for len(queue) != 0 {

	// 	for j := 0; j < cap; j++ {

	// 		curr := queue[j]

	// 		if curr.Left != nil {
	// 			queue = append(queue, curr.Left)
	// 		}
	// 		if curr.Right != nil {
	// 			queue = append(queue, curr.Right)
	// 		}

	// 		arr = append(arr, curr.Val)

	// 	}
	// 	queue = queue[cap:]
	// 	cap = len(queue)
	// 	result = append(result, arr)
	// 	arr = nil

	// }

	// return result
	//-> revise
	if root == nil {
		return nil
	}
	result := [][]int{}
	queue := []*TreeNode{root}
	cap := len(queue)
	temp := []int{}
	for len(queue) != 0 {
		//temp = temp[:0]
		for i := 0; i < cap; i++ {
			curr := queue[i]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			temp = append(temp, curr.Val)
		}
		result = append(result, temp)
		queue = queue[cap:]
		cap = len(queue)
		temp = nil
	}
	return result
	//<-

}
