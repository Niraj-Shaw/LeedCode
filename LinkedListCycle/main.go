package main

type ListNode struct {
	Val  int
	next *ListNode
}

func main() {

}
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
		return false

	}
	return false

}
