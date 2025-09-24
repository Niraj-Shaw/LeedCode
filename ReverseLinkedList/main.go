package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node5 := &ListNode{Val: 5, Next: nil}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	head := reverseList(node1)

	for head != nil {
		println(head.Val)
		head = head.Next
	}

}

func reverseList(head *ListNode) *ListNode {
	// var prevHead *ListNode

	// for head != nil{
	// 	nextNode := head.Next
	// 	head.Next = prevHead
	// 	prevHead = prevHead.Next
	// 	head = nextNode
	// }

	// return  prevHead

	// if head == nil || head.Next == nil {
	// 	return head
	// }

	// p := reverseList(head.Next)
	// head.Next.Next = head
	// head.Next = nil

	// return p

	if head == nil || head.Next == nil{
		return head
	}

	p := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return p

}
