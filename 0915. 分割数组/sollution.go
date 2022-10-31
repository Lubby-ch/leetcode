package main

import (
	"fmt"
	"sort"
)

func partitionDisjoint1(nums []int) int {
	var (
		n     = len(nums)
		asist = make([]int, n)
	)
	for i := 1; i < n; i++ {
		asist[i] = i
	}
	sort.Slice(asist, func(i, j int) bool {
		return nums[asist[i]] < nums[asist[j]] || (nums[asist[i]] == nums[asist[j]] && i <= j)
	})
	var max = asist[0]
	if max == 0 {
		return 1
	}
	for i := 1; i < n; i++ {
		if max < asist[i] {
			max = asist[i]
		}
		if max == i {
			return i + 1
		}
	}
	return -1
}



func partitionDisjoint(nums []int) int {
	var (
		n       = len(nums)
		leftmax = 0
		ans = -1
	)
	for i := 1; i < n; i++ {
		if nums[i] >= nums[leftmax] {
			if ans == -1 {
				ans = i
			}
			leftmax = i
		}
		if ans != -1 && nums[i] < nums[ans] {
			ans = -1
		}
	}
	return ans
}

func main() {
	fmt.Println(partitionDisjoint([]int{5,0,3,8,6}))
}
