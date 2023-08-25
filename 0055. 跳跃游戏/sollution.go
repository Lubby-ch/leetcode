package main

func canJump(nums []int) bool {
	var end = nums[0]
	for i := 0; i < len(nums); i++ {
		end = max(end, i+nums[i])
		if end >= len(nums)-1 {
			return true
		}
		if i == end {
			return false
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
