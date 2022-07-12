package main

import "fmt"

// 线段树
type MyCalendar struct {
	root *Node
}

type Node struct {
	start int
	end   int
	left  *Node
	right *Node
}

func Constructor() MyCalendar {
	return MyCalendar{
		root: &Node{
			start: -1,
			end:   0,
		},
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if start > end {
		return false
	}
	return insertChild(this.root, start, end)
}

func insertChild(root *Node, start, end int) bool {
	if start >= root.end {
		if root.right == nil {
			root.right = &Node{
				start: start,
				end:   end,
			}
			return true
		}
		return insertChild(root.right, start, end)
	}
	if end <= root.start {
		if root.left == nil {
			root.left = &Node{
				start: start,
				end:   end,
			}
			return true
		}
		return insertChild(root.left, start, end)
	}
	return false
}

func main() {
	lists := []struct {
		start, end int
	}{
		{10, 20},
		{15, 25},
		{20, 30},
	}
	Calendar := Constructor()
	for i := range lists {
		list := lists[i]
		fmt.Println(Calendar.Book(list.start, list.end))
	}
}
