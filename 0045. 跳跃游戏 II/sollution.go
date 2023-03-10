package main

func jump(nums []int) int {
	n := len(nums)
	var dp = make([]int, n)
	for i := n - 2; i >= 0; i-- {
		if nums[i] == 0 {
			continue
		}
		if nums[i]+i >= n-1 {
			dp[i] = 1
			continue
		}
		for j := 1; j <= nums[i]; j++ {
			if i+j >= n {
				break
			}
			if dp[i+j] == 0 {
				continue
			}
			if dp[i] == 0 {
				dp[i] = dp[i+j] + 1
			} else if dp[i+j]+1 < dp[i] {
				dp[i] = dp[i+j] + 1
			}
		}
	}
	return dp[0]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func jump1(nums []int) int {
	length := len(nums)
	maxPos := 0     // 下一步能跳的最远边界
	step := 0       // 步数
	lastMaxPos := 0 // 当前这一步到达的最远边界
	for i := 0; i < length; i++ {
		maxPos = max(maxPos, i+nums[i])
		if i == lastMaxPos {
			lastMaxPos = maxPos
			step++
		}
	}
	return step
}
