package main

import "fmt"

func oddCells(m int, n int, indices [][]int) int {
	var rmap = make(map[int]int)
	var cmap = make(map[int]int)
	for i := range indices {
		rmap[indices[i][0]] += 1
		cmap[indices[i][1]] += 1
	}
	var rcount = 0
	for _, v := range rmap {
		rcount += v % 2
	}
	var ccount = 0
	for _, v := range cmap {
		ccount += v % 2
	}
	return rcount*(n-ccount) + ccount*(m-rcount) // 同奇则为偶
}

func main() {
	testCases := []struct {
		m       int
		n       int
		indices [][]int
	}{
		{
			m:       2,
			n:       3,
			indices: [][]int{{0, 1}, {1, 1}},
		},
	}
	for _, testcase := range testCases {
		fmt.Println(oddCells(testcase.m, testcase.n, testcase.indices))
	}
}
