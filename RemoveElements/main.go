package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func removeElements(head *ListNode, val int) *ListNode {
	current := head
	temp := head
	for current != nil {

		if (current.Val == val) && (current == head) {
			head = current.Next
			temp = head
		} else if current.Val == val {

			temp.Next = current.Next
		} else {
			temp = current
		}
		current = current.Next
	}

	return head

}
