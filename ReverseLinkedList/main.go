package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reverseList(head *ListNode) *ListNode {
	var revHead *ListNode

	for head != nil {

		nextNode := head.Next
		head.Next = revHead
		revHead = head
		head = nextNode

	}
	return revHead

}
