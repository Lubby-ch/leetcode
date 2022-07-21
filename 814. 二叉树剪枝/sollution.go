package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	val := LRD(root)
	if val == 0 && root.Val == 0 {
		return nil
	}
	return root
}

func LRD(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := LRD(root.Left)
	r := LRD(root.Right)
	if l == 0 {
		root.Left = nil
	}
	if r == 0 {
		root.Right = nil
	}
	if l == 1 || r == 1 || root.Val == 1 {
		return 1
	}
	return 0
}
