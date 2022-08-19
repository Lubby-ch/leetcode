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

// 分支
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		var tmp []*ListNode
		var n = len(lists)
		for i := 0; i < n; i += 2 {
			if i+1 < n {
				tmp = append(tmp, mergeTwoLists(lists[i], lists[i+1]))
			} else {
				tmp = append(tmp, lists[i])
			}
		}
		lists = tmp
	}
	return lists[0]
}


