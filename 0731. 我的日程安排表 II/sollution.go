package main

import (
	"fmt"
	"math"
)

type MyCalendarTwo struct {
	node *Node
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		node: &Node{
			start: math.MinInt,
			end:   math.MinInt,
		},
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	if !this.node.Query(start, end) {
		return false
	}
	this.node.Insert(start, end, 1)
	return true
}

type Node struct {
	left  *Node
	right *Node
	start int
	end   int
	count int
}

func NewNode(start int, end int, count int) *Node {
	return &Node{
		start: start,
		end:   end,
		count: count,
	}
}

func (node *Node) Insert(start int, end int, count int) {
	if start >= node.end {
		if node.right != nil {
			node.right.Insert(start, end, count)
		} else {
			node.right = NewNode(start, end, count)
		}
		return
	} else if end <= node.start {
		if node.left != nil {
			node.left.Insert(start, end, count)
		} else {
			node.left = NewNode(start, end, count)
		}
		return
	}
	if start < node.start {
		if node.left != nil {
			node.left.Insert(start, node.start, count)
		} else {
			node.left = NewNode(start, node.start, count)
		}
		if end <= node.end {
			node.end = end
			if node.right != nil {
				node.right.Insert(end, node.end, node.count)
			} else {
				node.left = NewNode(end, node.end, node.count)
			}
		} else {
			if node.right != nil {
				node.right.Insert(node.end, end, count)
			} else {
				node.left = NewNode(node.end, end, count)
			}
		}
	} else {
		if node.start != start {
			node.start = start
			if node.left != nil {
				node.left.Insert(node.start, start, node.count)
			} else {
				node.left = NewNode(node.start, start, node.count)
			}
		}
		if end <= node.end {
			node.end = end
			if node.right != nil {
				node.right.Insert(end, node.end, node.count)
			} else {
				node.left = NewNode(node.start, start, node.count)
			}
		} else {
			if node.right != nil {
				node.right.Insert(node.end, end, count)
			} else {
				node.left = NewNode(node.end, end, count)
			}
		}
	}
	node.count++
}

//type SegmentNode struct {
//	lt, rt *SegmentNode
//	num, lazy int
//}
//func (node *SegmentNode) update(L int, R int, start int, end int) {
//	if start <= L && R <= end {
//		node.num++; node.lazy++
//		return
//	}
//	node.pushdown()
//	mid := (L + R) >> 1
//	if start <= mid {node.lt.update(L, mid, start, end)}
//	if end > mid {node.rt.update(mid + 1, R, start, end)}
//	node.pushup()
//}
//func (node *SegmentNode) pushup() {node.num = max(node.lt.num, node.rt.num)}
//func (node *SegmentNode) pushdown() {
//	if node.lt == nil {node.lt = new(SegmentNode)}
//	if node.rt == nil {node.rt = new(SegmentNode)}
//	if node.lazy > 0 {
//		node.lt.num += node.lazy
//		node.lt.lazy += node.lazy
//		node.rt.num += node.lazy
//		node.rt.lazy += node.lazy
//		node.lazy = 0
//	}
//}
//func (node *SegmentNode) query(L, R, start, end int) (ans int) {
//	if start <= L && R <= end {
//		return node.num
//	}
//	node.pushdown()
//	mid := (L + R) >> 1
//	if start <= mid {ans = node.lt.query(L, mid, start, end)}
//	if end > mid {ans = max(ans, node.rt.query(mid + 1, R, start, end))}
//	return
//}
//func max(a, b int) int {
//	if a > b {return a}
//	return b
//}
//type MyCalendarTwo struct {
//	Root SegmentNode
//}
//func Constructor() MyCalendarTwo {return MyCalendarTwo{}}
//
//func (this *MyCalendarTwo) Book(start int, end int) bool {
//	if this.Root.query(0, 1e9, start, end -1) >=2 {return false}
//	this.Root.update(0, 1e9, start, end - 1)
//	return true
//}

func (node *Node) Query(start int, end int) bool {
	if start >= node.end {
		if node.right != nil {
			return node.right.Query(start, end)
		}
		return true
	}
	if end <= node.start {
		if node.left != nil {
			return node.left.Query(start, end)
		}
		return true
	}
	if node.count == 2 {
		return false
	}
	var ok = true
	if start < node.start {
		if node.left != nil {
			ok = ok && node.left.Query(start, node.start)
			if !ok {
				return false
			}
		}
		if end > node.end && node.right != nil {
			ok = ok && node.right.Query(node.end, end)
		}
	} else {
		if end > node.end && node.right != nil {
			ok = ok && node.right.Query(node.end, end)
		}
	}
	return ok
}

//["MyCalendarTwo","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book"]
//[[],[47,50],[1,10],[27,36],[40,47],[20,27],[15,23],[10,18],[27,36],[17,25],[8,17],[24,33],[23,28],[21,27],[47,50],[14,21],[26,32],[16,21],[2,7],[24,33],[6,13],[44,50],[33,39],[30,36],[6,15],[21,27],[49,50],[38,45],[4,12],[46,50],[13,21]]
func main() {
	//testCases := []struct {
	//	start int
	//	end   int
	//}{
	//	{
	//		47, 50,
	//	},
	//	{
	//		1, 10,
	//	},
	//	{
	//		27, 36,
	//	},
	//	{
	//		40, 47,
	//	},
	//	{
	//		20, 27,
	//	},
	//	{
	//		15, 23,
	//	},
	//	{
	//		10, 18,
	//	},
	//	{
	//		27, 36,
	//	},
	//	{
	//		17, 25,
	//	},
	//	{
	//		8, 17,
	//	},
	//}
	//c := Constructor()
	//for _, testcase := range testCases {
	//	fmt.Println(c.Book(testcase.start, testcase.end))
	//}
	var state uint32
	for i := 1;i <= 10;i++ {
		if state == 5 {
			continue
		}
		state |= 1<<i
	}
	fmt.Println(state&(1<<5))
}
