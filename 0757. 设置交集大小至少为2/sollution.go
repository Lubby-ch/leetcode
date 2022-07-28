package main

import (
	"fmt"
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i int, j int) bool {
		a, b := intervals[i], intervals[j]
		return a[1] < b[1] || (a[1] == b[1] && a[0] > b[0])
	})
	length := len(intervals)
	var vals []int
	vals = append(vals, intervals[0][1]-1, intervals[0][1])
	for i := 1; i < length; i++ {
		if vals[len(vals)-2] >= intervals[i][0] {
			continue
		}
		if vals[len(vals)-1] < intervals[i][0] {
			vals = append(vals, intervals[i][1]-1)
		}
		vals = append(vals, intervals[i][1])
	}
	return len(vals)
}

func main() {
	testCases := []struct {
		intervals [][]int
	}{
		{
			intervals: [][]int{{2, 5}, {3, 5}, {1, 3}, {1, 4}},
		},
		{
			intervals: [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}},
		},
	}
	for _, testcase := range testCases {
		fmt.Println(intersectionSizeTwo(testcase.intervals))
	}
}
