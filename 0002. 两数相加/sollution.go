package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carryBit int
	var head = &ListNode{}
	var ptr = head
	for {
		var val int
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		val += carryBit
		carryBit = val / 10
		val = val % 10
		node := &ListNode{
			Val: val,
		}
		ptr.Next = node
		ptr = ptr.Next
		if l1 == nil && l2 == nil && carryBit == 0 {
			break
		}
	}
	return head.Next
}