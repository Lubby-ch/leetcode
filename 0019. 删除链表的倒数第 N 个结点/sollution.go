package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 栈
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var stack []*ListNode
	ptr := head
	for ptr != nil {
		stack = append(stack, ptr)
		ptr = ptr.Next
	}
	if len(stack)-n-1 < 0 {
		return stack[len(stack)-n].Next
	}
	stack[len(stack)-n-1].Next = stack[len(stack)-n].Next
	return head
}

// 双指针
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	var first, second = head, head
	for {
		if first == nil {
			break
		}
		first = first.Next
		if n > -1 {
			n--
			continue
		}
		second = second.Next
	}
	if n == 0 {
		return head.Next
	}
	second.Next = second.Next.Next
	return head
}
