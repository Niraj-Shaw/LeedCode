package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	prevNode := head
	for current != nil {
		if (current != head) && (current.Val == prevNode.Val) {
			prevNode.Next = current.Next
		} else {
			prevNode = current
		}
		current = current.Next

	}
	return head
}
