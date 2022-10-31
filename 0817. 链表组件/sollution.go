package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	var numMap = make(map[int]struct{})
	for _, v := range nums {
		numMap[int(v)] = struct{}{}
	}
	var ans int
	var count int16
	for {
		if head == nil {
			if count > 0 {
				ans++
			}
			break
		}
		if _, ok := numMap[head.Val]; ok {
			count++
		} else {
			if count > 0 {
				ans++
				count = 0
			}
		}
		head = head.Next
	}
	return ans
}
