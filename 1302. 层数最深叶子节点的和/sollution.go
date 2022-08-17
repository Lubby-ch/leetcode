package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1.栈 2.深度优先遍历
func deepestLeavesSum(root *TreeNode) int {
	maxlevel := -1
	var dfs func(*TreeNode, int)
	var ans int
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level > maxlevel {
			maxlevel = level
			ans = node.Val
		} else if level == maxlevel {
			ans += node.Val
		}
		dfs(node.Right, level+1)
		dfs(node.Left, level+1)
	}
	dfs(root, 0)
	return ans
}
