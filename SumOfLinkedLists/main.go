package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n1 := &ListNode{}
	n1.Val = 2
	n2 := &ListNode{}
	n2.Val = 4
	n3 := &ListNode{}
	n3.Val = 3
	n1.Next = n2
	n2.Next = n3

	l1 := &ListNode{}
	l1.Val = 5
	l2 := &ListNode{}
	l2.Val = 6
	l3 := &ListNode{}
	l3.Val = 4
	l1.Next = l2
	l2.Next = l3
	list := addTwoNumbers(n1, l1)
	for list != nil {
		fmt.Print(list.Val)
		list = list.Next
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// sum, carry := 0, 0
	// result := &ListNode{}
	// for node := result; l1 != nil || l2 != nil || carry > 0; node = node.Next {
	// 	L1, L2 := 0, 0
	// 	if l1 != nil {
	// 		L1 = l1.Val
	// 		l1 = l1.Next
	// 	}
	// 	if l2 != nil {
	// 		L2 = l2.Val
	// 		l2 = l2.Next
	// 	}
	// 	sum = L1 + L2 + carry
	// 	carry = sum / 10
	// 	sum = sum % 10

	// 	node.Next = &ListNode{sum, nil}

	// }
	// return result.Next

	Carry := 0
	result := new(ListNode)

	for node := result; l1 != nil || l2 != nil || Carry > 0; node = node.Next {
		if l1 != nil {
			Carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			Carry += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode{Carry % 10, nil}
		Carry /= 10

	}
	return result.Next

}
