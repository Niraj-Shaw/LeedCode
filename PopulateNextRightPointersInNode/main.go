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
	// if root == nil{
	//     return nil
	// }

	// q := []*Node{root}

	// for len(q) != 0{
	//     n := len(q)

	//     for i := 0; i < n; i++{
	//         node := q[i]

	//         if node.Left != nil{
	//             q = append(q, node.Left)
	//         }

	//         if node.Right != nil{
	//             q = append(q, node.Right)
	//         }

	//          if i + 1 < n{
	//             node.Next = q[i+1]
	//         }

	//     }

	//     // for i := 0; i < n; i++{
	//     //     node := q[i]

	//     //     if i + 1 < n{
	//     //         node.Next = q[i+1]
	//     //     }

	//     // }

	//     q = q[n:]

	if root == nil {
		return root
	}

	if root.Left != nil {
		root.Left.Next = root.Right
	}

	if root.Right != nil && root.Next != nil {
		root.Right.Next = root.Next.Left
	}

	connect(root.Left)
	connect(root.Right)

	return root
}
