package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {

}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Right != nil {
		root.Right.Next = helperConnect(root)
		connect(root.Right)
	}

	if root.Left != nil {
		if root.Right != nil {
			root.Left.Next = root.Right
		} else {
			root.Left.Next = helperConnect(root)
		}
		connect(root.Left)
	}

	return root
}

func helperConnect(node *Node) *Node {

	curr := node.Next

	for curr != nil {
		if curr.Left != nil {
			return curr.Left
		}

		if curr.Right != nil {
			return curr.Right
		}

		curr = curr.Next

	}

	return nil
}
