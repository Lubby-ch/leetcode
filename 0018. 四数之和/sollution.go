package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			var fourth = n - 1
			for k := j + 1; k < n; k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				for nums[fourth]+nums[i]+nums[j]+nums[k] > target && fourth > k {
					fourth--
				}
				if fourth <= k {
					break
				}
				if nums[fourth]+nums[i]+nums[j]+nums[k] < target {
					continue
				}
				ans = append(ans, []int{nums[i], nums[j], nums[k], nums[fourth]})
			}
		}
	}
	return ans
}
