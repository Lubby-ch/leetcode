package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var pre, cur *ListNode
	pre = &ListNode{}
	pre.Next = head
	cur = head
	head = pre
	for pre != nil && cur != nil && cur.Next != nil {
		pre.Next = cur.Next
		cur.Next = pre.Next.Next
		pre.Next.Next = cur

		pre = cur
		cur = cur.Next
	}
	return head.Next
}
