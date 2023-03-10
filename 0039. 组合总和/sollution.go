package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) (ret [][]int) {
	var dfs func(target, idx int)
	var ans []int
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ret = append(ret, append([]int{}, ans...))
			return
		}
		dfs(target, idx+1)
		if target-candidates[idx] >= 0 {
			ans = append(ans, candidates[idx])
			dfs(target-candidates[idx], idx+1)
			ans = ans[:len(ans)-1]
		}
	}
	dfs(target, 0)
	return ret
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
