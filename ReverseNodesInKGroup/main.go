package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	str := " "
	//Node5 := &ListNode{5, nil}
	Node4 := &ListNode{4, nil}
	Node3 := &ListNode{3, Node4}
	Node2 := &ListNode{2, Node3}
	Node1 := &ListNode{1, Node2}
	Node1.Next = Node2
	Node2.Next = Node3
	Node3.Next = Node4
	//Node4.Next = Node5
	revHead := reverseKGroup(Node1, 2)
	for revHead != nil {
		//fmt.Println(revHead.Val)
		str = str + " " + strconv.Itoa(revHead.Val)
		revHead = revHead.Next
	}
	fmt.Print(str)

}
func reverseKGroup(head *ListNode, k int) *ListNode {
	// solution for different question where only first and kth are reversed
	// start := 1
	// temNode := &ListNode{0, nil}
	// First := head
	// var Prev *ListNode
	// for head != nil {
	// 	if start == 1 {
	// 		temNode = head

	// 	}

	// 	if start == k-1 && head.Next != nil {
	// 		nextNode := head.Next
	// 		head.Next = &ListNode{temNode.Val, temNode.Next}
	// 		head.Next.Next = nextNode.Next
	// 		nextNode.Next = temNode.Next
	// 		if Prev != nil {
	// 			Prev.Next = nextNode
	// 		}
	// 		if temNode.Val == First.Val && temNode.Next == First.Next {
	// 			First.Val = nextNode.Val
	// 			First.Next = nextNode.Next
	// 		}

	// 	}
	// 	if start == k {
	// 		Prev = head
	// 	}
	// 	head = head.Next
	// 	if start == k {
	// 		start = 1
	// 	} else {
	// 		start++
	// 	}
	// }
	// return First
	// solution for different question where only first and kth are reversed <--

	// recursive approach

	if head == nil {
		return nil
	}
	if k == 1 {
		return head
	}
	index := k
	curr := head
	var prev *ListNode
	for ; curr != nil && index > 0; index-- {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next

	}
	if index > 0 {
		return reverseKGroup(prev, k-index)
	}
	head.Next = reverseKGroup(curr, k)

	return prev

}
