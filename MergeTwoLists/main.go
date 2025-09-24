package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	// dummy := &ListNode{}
    // node := dummy
    // for list1 != nil && list2 != nil{
    //     if list1.Val >= list2.Val{
    //         node.Next = list1
    //         list1 = list1.Next
    //     }else{
    //         node.Next = list2
    //         list2 = list2.Next         
    //     }
    //     node = node.Next
    // }
    // if list1 != nil{
    //     node.Next = list1
    //     list1 = list1.Next
        
    // }
    // if list2 != nil{
    //     node.Next = list2
    //     list2 = list2.Next
    // }
    // return dummy.Next
//----
	// dummy := new(ListNode)
	// current := dummy

	// for list1 != nil && list2 != nil {

	// 	if list1.Val <= list2.Val {
	// 		current.Next = list1
	// 		list1 = list1.Next
	// 	} else {
	// 		current.Next = list2
	// 		list2 = list2.Next
	// 	}
	// 	current = current.Next
	// }

	// if list1 == nil {
	// 	current.Next = list2
	// }

	// if list2 == nil {
	// 	current.Next = list1
	// }
	// return dummy.Next

	if list1 == nil{
		return list2
	}else if list2 == nil{
		return list1
	}else{
		if list1.Val < list2.Val{
			list1.Next = mergeTwoLists(list1.Next, list2)
			return list1
		}else{
			list2.Next = mergeTwoLists(list1, list2.Next)
			return list2
		}
	}

}
