package main

func insert(intervals [][]int, newInterval []int) [][]int {
	var ans [][]int
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] > newInterval[0] && intervals[i][0] < newInterval[1] { // 有交集更新newInterval
			newInterval[1] = max(newInterval[1], intervals[i][1])
		} else if intervals[i][0] < newInterval[0] && intervals[i][1] > newInterval[0] {
			newInterval[0] = intervals[i][0]
			newInterval[1] = max(newInterval[1], intervals[i][1])
		} else {
			if newInterval[0] < intervals[i][0] {
				ans = append(ans, newInterval)
			}
			ans = append(ans, intervals[i])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
