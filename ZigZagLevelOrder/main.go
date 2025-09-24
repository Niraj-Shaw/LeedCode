package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node5 := &TreeNode{7, nil, nil}
	node4 := &TreeNode{15, nil, nil}
	node3 := &TreeNode{20, node4, node5}
	node2 := &TreeNode{9, nil, nil}
	node1 := &TreeNode{3, node2, node3}
	//node1 := &TreeNode{1, nil, nil}
	fmt.Printf("%v", zigzagLevelOrder(node1))

}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	nodeList := []*TreeNode{root}
	values := []int{}
	isLeft := true
	levelCount := len(nodeList)
	for len(nodeList) > 0 {
		for i := 0; i < levelCount; i++ {
			currNode := nodeList[i]
			if isLeft {
				values = append(values, currNode.Val)
			} else {
				values = append([]int{currNode.Val}, values...)
			}
			if currNode.Left != nil {
				nodeList = append(nodeList, currNode.Left)
			}
			if currNode.Right != nil {
				nodeList = append(nodeList, currNode.Right)
			}
		}
		if values != nil {
			result = append(result, values)
		}
		values = nil
		nodeList = nodeList[levelCount:]
		levelCount = len(nodeList)
		isLeft = !isLeft
	}

	return result

}
