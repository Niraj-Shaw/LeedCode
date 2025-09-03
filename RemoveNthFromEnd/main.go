package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// node5 := ListNode{5, nil}
	// node4 := ListNode{4, &node5}
	// node3 := ListNode{3, &node4}
	node2 := ListNode{2, nil}
	node1 := ListNode{1, &node2}
	node := removeNthFromEnd(&node1, 1)

	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	// dummy := &ListNode{0, head}
	// node1 := dummy
	// node2 := dummy

	// for i := 1; i <= n+1; i++ {
	// 	node1 = node1.Next
	// }

	// for node1 != nil {
	// 	node1 = node1.Next
	// 	node2 = node2.Next
	// }

	// node2.Next = node2.Next.Next

	// return dummy.Next

	nodeCnt := 0
	dummy := &ListNode{0, head}
	node := head

	for node != nil {
		nodeCnt++
		node = node.Next
	}

	nodeCnt -= n

	node = dummy

	for nodeCnt > 0 {
		node = node.Next
		nodeCnt--
	}

	node.Next = node.Next.Next

	return dummy.Next
}
