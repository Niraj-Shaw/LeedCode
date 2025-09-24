package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{8, node4}
	node2 := &ListNode{1, node3}
	node1 := &ListNode{4, node2}
	node11 := &ListNode{5, nil}
	node10 := &ListNode{4, node11}
	node9 := &ListNode{8, node10}
	node8 := &ListNode{1, node9}
	node7 := &ListNode{6, node8}
	node6 := &ListNode{5, node7}

	node := getIntersectionNode(node1, node6)

	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB

	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = headB
		}

		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = headA
		}
	}

	return p1
}
