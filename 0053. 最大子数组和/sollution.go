package main

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	var ans = nums[0]
	var cur int

	for i := 0; i < n; i++ {
		cur += nums[i]
		if ans < cur {
			ans = cur
		}
		if cur < 0 {
			cur = 0
		}
	}
	return ans
}
