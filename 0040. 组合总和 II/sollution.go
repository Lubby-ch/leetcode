package main

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var ans [][]int
	var dfs func(target, idx int)
	var tmp []int
	var used = make(map[int]bool)
	dfs = func(target, idx int) {
		if target == 0 {
			ans = append(ans, append([]int{}, tmp...))
			return
		}
		if idx >= len(candidates) || candidates[idx] > target {
			return
		}
		dfs(target, idx+1)
		if idx > 0 && candidates[idx] == candidates[idx-1] && used[idx-1] == false {
			return
		}

		if target >= candidates[idx] {
			used[idx] = true
			tmp = append(tmp, candidates[idx])
			dfs(target-candidates[idx], idx+1)
			used[idx] = false
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(target, 0)
	return ans
}

func main() {
	combinationSum2([]int{2, 5, 2, 1, 2}, 5)
}
