package main

import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}

func main(){
	node7 := &ListNode{5, nil}
	node6 := &ListNode{5, node7}
	node5 := &ListNode{5, node6}

	node4 := &ListNode{3, nil}
	node3 := &ListNode{4, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{7, node2}

	result := addTowNumbers(node1, node5)
	for result != nil{
		fmt.Println(result.Val)
		result = result.Next
	}
}

func addTowNumbers(l1, l2 *ListNode) *ListNode{
	//reverse
	//sum
	//return the result
}

