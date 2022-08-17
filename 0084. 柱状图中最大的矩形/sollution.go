package main

import "fmt"

// 动态规划
func largestRectangleArea(heights []int) int {
	var n = len(heights)
	var minLeft = make([]int, n)
	var minRight = make([]int, n)
	minLeft[0] = -1
	minRight[n-1] = n
	for i := 1; i < n; i++ {
		j := i - 1
		for j >= 0 {
			if heights[j] >= heights[i] && j != -1 {
				j = minLeft[j]
				continue
			}
			break
		}
		minLeft[i] = j
	}
	for i := n - 2; i >= 0; i-- {
		j := i + 1
		for j < n {
			if heights[j] >= heights[i] {
				j = minRight[j]
				continue
			}
			break
		}
		minRight[i] = j
	}
	var ans int
	for i := 0; i < n; i++ {
		ans = max(ans, (minRight[i]-minLeft[i]-1)*heights[i])
	}
	return ans
}

// 单调栈
func largestRectangleArea1(heights []int) int {
	var n = len(heights)
	var minLeft = make([]int, n)
	var minRight = make([]int, n)
	var stack = make([]int, 0)
	for i := range heights {
		minRight[i] = n
		for len(stack) != 0 && heights[i] <= heights[stack[len(stack)-1]] {
			// 出栈
			minRight[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			minLeft[i] = stack[len(stack)-1]
		} else {
			minLeft[i] = -1
		}
		// 进栈
		stack = append(stack, i)
	}
	var ans int
	for i := 0; i < n; i++ {
		ans = max(ans, (minRight[i]-minLeft[i]-1)*heights[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(largestRectangleArea1([]int{2, 4}))
}
