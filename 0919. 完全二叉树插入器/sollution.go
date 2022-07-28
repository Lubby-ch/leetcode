package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	root      *TreeNode
	candidate []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	queue := []*TreeNode{root}
	var candidate []*TreeNode
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left == nil || node.Right == nil {
			candidate = append(candidate, node)
		}
	}
	return CBTInserter {
		root: root,
		candidate: candidate,
	}
}

func (this *CBTInserter) Insert(val int) int {
	child := &TreeNode{
		Val: val,
	}
	node := this.candidate[0]
	if node.Left == nil {
		node.Left = child
	} else {
		node.Right = child
		this.candidate = this.candidate[1:]
	}
	this.candidate = append(this.candidate, child)
	return node.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
