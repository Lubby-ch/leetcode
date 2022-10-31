package main

import "fmt"

// 基于状态转变的动态规划算法
func minSwap1(nums1 []int, nums2 []int) int {
	n := len(nums1)
	var dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = 1
	for i := 1; i < n; i++ {
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
				dp[i][0] = min(dp[i-1][0], dp[i-1][1])
				dp[i][1] = min(dp[i-1][0]+1, dp[i-1][1]+1)
			} else {
				dp[i][0] = dp[i-1][0]
				dp[i][1] = dp[i-1][1] + 1
			}
		} else {
			dp[i][0] = dp[i-1][1]
			dp[i][1] = dp[i-1][0] + 1
		}
	}

	return min(dp[n-1][0], dp[n-1][1])
}

// 深度优先算法 超出时间限制
func minSwap2(nums1 []int, nums2 []int) int {
	n := len(nums1)
	var dfs func(index, times, flag int) int
	dfs = func(index, times, flag int) int {
		if index == n {
			return times
		}
		if nums1[index] > nums1[index-1] && nums2[index] > nums2[index-1] {
			if nums1[index] > nums2[index-1] && nums2[index] > nums1[index-1] {
				return min(dfs(index+1, times+1, 1), dfs(index+1, times, 0))
			} else {
				return dfs(index+1, times+flag, flag)
			}
		} else {
			return dfs(index+1, times+(flag^1), flag^1)
		}
	}
	return min(dfs(1, 1, 1), dfs(1, 0, 0))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var tmp = 0
	for i := 0; i <= 10; i++ {
		tmp ^= 1
		fmt.Println(tmp)
	}
}
