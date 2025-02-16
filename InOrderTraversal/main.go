package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	curr := root
	for (curr != nil) || (len(stack) != 0) {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left

		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, curr.Val)
			curr = curr.Right
		}

	}
	return res
}
