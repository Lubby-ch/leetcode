package main

import (
	"fmt"
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	var sum = 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	partSum := sum / k
	sort.Ints(nums)
	if nums[len(nums)-1] > partSum {
		return false
	}
	var dfs func(buckets []int, index int, target int) bool
	dfs = func(buckets []int, index int, target int) bool {
		if index < 0 {
			return true
		}
		for i := 0; i < k; i++ {
			if buckets[i]+nums[index] <= target {
				buckets[i] += nums[index]
				if dfs(buckets, index-1, target) {
					return true
				}
				buckets[i] -= nums[index]
			}
			if buckets[i] == 0 {
				break
			}
		}
		return false
	}
	return dfs(make([]int, k), len(nums)-1, partSum)
}

func main() {
	fmt.Println(canPartitionKSubsets([]int{1, 1, 1, 1, 2, 2, 2, 2}, 2))
}
