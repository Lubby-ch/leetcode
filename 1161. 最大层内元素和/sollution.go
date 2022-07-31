package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广度优先搜索
func maxLevelSum(root *TreeNode) int {
	var ans int
	var queue = []*TreeNode{root}
	var childNum = 1
	var curLevel = 0
	var max = math.MinInt

	for queue != nil {
		var sum int
		var addNum int
		for i := 0; i < childNum; i++ {
			sum += queue[0].Val
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
				addNum++
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
				addNum++
			}
			queue = queue[1:]
			childNum = addNum
		}
		curLevel++
		if sum > max {
			max = sum
			ans = curLevel
		}
	}
	return ans
}

// 深度优先搜索
func maxLevelSum1(root *TreeNode) int {
	var sum []int
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if len(sum) == level {
			sum = append(sum, node.Val)
		} else {
			sum[level] += node.Val
		}
		if node.Left != nil {
			dfs(node.Left, level+1)
		}
		if node.Right != nil {
			dfs(node.Right, level+1)
		}
	}
	dfs(root, 0)
	var ans = 0
	for i := 1; i < len(sum); i++ {
		if sum[i] > sum[ans] {
			ans = i
		}
	}
	return ans + 1
}
