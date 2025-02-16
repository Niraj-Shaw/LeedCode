package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

/*
	type tree struct {
		list []int
	}

	func (t *tree) traverse(node *TreeNode) {
		if node != nil {
			t.list = append(t.list, node.Val)
			t.traverse(node.Left)
			t.traverse(node.Right)
		}

}

	func preorderTraversal(root *TreeNode) []int {
		t := tree{}
		t.traverse(root)
		return t.list
	}
*/
func preorderTraversal(root *TreeNode) []int {
	answer := []int{}
	stack := []*TreeNode{root}
	if root != nil {

		for len(stack) != 0 {
			currNode := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if currNode != nil {
				answer = append(answer, currNode.Val)
				if currNode.Right != nil {
					stack = append(stack, currNode.Right)
				}
				if currNode.Left != nil {
					stack = append(stack, currNode.Left)
				}
			}
		}
	}
	return answer

}
