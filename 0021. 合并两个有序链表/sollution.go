package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var ptr = new(ListNode)
	var head = ptr
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			ptr.Next = list1
			list1 = list1.Next
		} else {
			ptr.Next = list2
			list2 = list2.Next
		}
		ptr = ptr.Next
	}
	if list1 != nil {
		ptr.Next = list1
	}
	if list2 != nil {
		ptr.Next = list2
	}
	return head.Next
}
