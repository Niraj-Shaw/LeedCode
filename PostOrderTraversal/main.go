package main

import (
	"fmt"
	"tree"
)

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func main() {
	myTree := []*tree.TreeNode{}
	node3 := tree.TreeNode{}
	node3.Val = 3
	node3.Left = nil
	node3.Right = nil
	node2 := tree.TreeNode{}
	node2.Val = 2
	node2.Left = &node3
	node2.Right = nil
	node1 := tree.TreeNode{}
	node1.Val = 1
	node1.Left = nil
	node1.Right = &node2
	myTree = append(myTree, &node1, &node2, &node3)
	fmt.Print(postorderTraversal(&node1))

}

/*func postorderTraversal(root *TreeNode) []int {
	t := tree{}
	t.traverse(root)
	return t.list

}

func (t *tree) traverse(node *TreeNode) {
	if node != nil {
		t.traverse(node.Left)
		t.traverse(node.Right)
		t.list = append(t.list, node.Val)
	}

}*/

func postorderTraversal(root *tree.TreeNode) []int {
	res := []int{}
	stack := []*tree.TreeNode{}
	curr := root
	var PrevNode *tree.TreeNode
	for curr != nil || len(stack) != 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left

		} else {
			curr = stack[len(stack)-1]
			if curr.Right != nil && curr.Right != PrevNode {
				curr = curr.Right
			} else {
				res = append(res, curr.Val)
				PrevNode = curr
				stack = stack[:len(stack)-1]
				curr = nil
			}

		}
	}
	return res

}
