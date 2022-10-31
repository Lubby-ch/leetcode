package main

import (
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	height := getHeight(root)
	weight := int(math.Pow(2, float64(height))) - 1
	var ans = make([][]string, height)
	ans[0] = make([]string, weight)
	ans[0][uint((weight-1)/2)] = strconv.Itoa(root.Val)
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, row, column int) {
		if node.Right != nil {
			if len(ans[row+1]) == 0 {
				ans[row+1] = make([]string, weight)
			}
			ans[row+1][column+int(math.Pow(2, float64(height-2-row)))] = strconv.Itoa(node.Right.Val)
			dfs(node.Left, row+1, column+1)
		}
		if node.Left != nil {
			if len(ans[row+1]) == 0 {
				ans[row+1] = make([]string, weight)
			}
			ans[row+1][column+int(math.Pow(2, float64(height-2-row)))] = strconv.Itoa(node.Right.Val)
			dfs(node.Left, row+1, column-1)
		}
	}
	dfs(root, 0, (weight-1)/2)
	return ans
}

func getHeight(root *TreeNode) int {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		var lh, rh int
		if node.Right != nil {
			rh = dfs(node.Right)
		}

		if node.Left != nil {
			lh = dfs(node.Left)
		}

		return max(rh, lh) + 1
	}
	return dfs(root)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}