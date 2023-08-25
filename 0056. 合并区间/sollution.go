package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var ans = make([][]int, 0)
	ans = append(ans, intervals[0])
	for i := 0; i < len(intervals); i++ {
		if len(ans) == 0 {
			ans = append(ans, intervals[i])
			continue
		}
		if intervals[i][0] <= ans[len(ans)-1][1] { // 合并
			if ans[len(ans)-1][1] >= intervals[i][1] {
				continue
			}
			ans[len(ans)-1][1] = intervals[i][1]
		} else {
			ans = append(ans, intervals[i])
		}
	}
	return ans
}
