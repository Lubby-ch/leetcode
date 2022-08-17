package main

import "fmt"

// 单调栈
func trap1(height []int) int {
	var stack = make([]int, 0)
	var ans int
	for i, v := range height {
		if len(stack) != 0 && height[stack[len(stack)-1]] <= v {
			var sub int
			var tmp int
			for {
				tmp = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				//
				if len(stack) > 0 {
					sub += (tmp - stack[len(stack)-1]) * height[tmp]
				}
				if len(stack) == 0 || height[stack[len(stack)-1]] >= v {
					// 进栈
					break
				}
			}
			if len(stack) != 0 {
				tmp = stack[len(stack)-1]
			}
			ans += (i-tmp-1)*min(height[tmp], v) - sub
		}
		stack = append(stack, i)
	}
	return ans
}

// 动态规划
func trap(height []int) int {
	n := len(height)
	var maxLeft = make([]int, n)
	var maxRight = make([]int, n)
	for i := 1; i < n; i++ {
		if height[i-1] >= maxRight[i-1] {
			maxRight[i] = height[i-1]
		} else {
			maxRight[i] = maxRight[i-1]
		}
	}
	for i := n - 2; i >= 0; i-- {
		if height[i+1] >= maxLeft[i+1] {
			maxLeft[i] = height[i+1]
		} else {
			maxLeft[i] = maxLeft[i+1]
		}
	}
	var ans int
	for i := 1; i < n-1; i++ {
		if height[i] >= min(maxLeft[i], maxRight[i]) {
			continue
		}
		fmt.Println(max(min(maxLeft[i],maxRight[i]) - height[i],0))
		ans += (max(min(maxLeft[i],maxRight[i]) - height[i], 0))
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(trap([]int{4, 2, 0, 3}))
}
