package main

import (
	"fmt"
	"sort"
)

// 单调栈
func maxChunksToSorted1(arr []int) int {
	var stack []int
	for _, v := range arr {
		if len(stack) > 0 && stack[len(stack)-1] > v {
			// 出栈
			tmp := stack[len(stack)-1]
			for len(stack) > 0 && stack[len(stack)-1] > v {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, tmp)
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack)
}

// 动态规划
func maxChunksToSorted2(arr []int) int {
	var n = len(arr)
	var maxLeft = make([]int, n)
	var minRight = make([]int, n)
	var ans = 1
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			minRight[i] = arr[n-1]
			continue
		}
		minRight[i] = min(arr[i], minRight[i+1])
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			maxLeft[i] = arr[0]
			continue
		}
		maxLeft[i] = max(arr[i], maxLeft[i-1])
		if minRight[i] >= maxLeft[i-1] {
			ans++
		}
	}
	return ans
}

// 贪心
func maxChunksToSorted(arr []int) int {
	var n = len(arr)
	var sum1, sum2, ans int
	var cp = make([]int, n)
	copy(cp, arr)
	sort.Ints(cp)
	for i := 0; i < n; i++ {
		sum1 += arr[i]
		sum2 += cp[i]
		if sum1 == sum2 {
			ans++
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(maxChunksToSorted([]int{1, 3, 1, 4}))
	fmt.Println(-1%10)
}
