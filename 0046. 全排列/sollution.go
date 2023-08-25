package main

func permute(nums []int) [][]int {
	var ans [][]int
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(nums)-1 {
			return
		}
		for i := idx; i < len(nums); i++ {
			if nums[i] == nums[idx] {
				dfs(idx + 1)
				continue
			}
			nums[idx], nums[i] = nums[i], nums[idx]
			ans = append(ans, append([]int{}, nums...))
			dfs(idx + 1)
			nums[idx], nums[i] = nums[i], nums[idx]
		}
	}
	ans = append(ans, append([]int{}, nums...))
	dfs(0)
	return ans
}
