package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	var ans [][]int
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(nums)-1 {
			return
		}
		for i := idx; i < len(nums); i++ {
			if i == idx {
				dfs(idx + 1)
				continue
			}
			if nums[i] == nums[i-1] || nums[idx]== nums[i] {
				continue
			}
			nums[idx], nums[i] = nums[i], nums[idx]
			fmt.Println(nums, " ", idx, " ",i)
			ans = append(ans, append([]int{}, nums...))
			dfs(idx + 1)
			nums[idx], nums[i] = nums[i], nums[idx]
		}
	}
	ans = append(ans, append([]int{}, nums...))
	dfs(0)
	return ans
}

func main() {
	permuteUnique([]int{-1,-1,-1,0,1})
}