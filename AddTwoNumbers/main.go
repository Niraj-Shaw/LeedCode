package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node3 := &ListNode{3, nil}
	node2 := &ListNode{4, node3}
	node1 := &ListNode{2, node2}

	node6 := &ListNode{4, nil}
	node5 := &ListNode{6, node6}
	node4 := &ListNode{5, node5}

	result := addTwoNumbers(node1, node4)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0
	dummy := &ListNode{}
	node := dummy

	for l1 != nil || l2 != nil || carry != 0 {
		v1, v2 := 0, 0

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		carry = sum / 10
		sum = sum % 10
		node.Next = &ListNode{Val: sum}
		node = node.Next

	}

	return dummy.Next

}
