package main

import "fmt"

func minCostToMoveChips(position []int) int {
	var oddNum, evenNum int
	for i := range position {
		chipPos := position[i]
		if chipPos%2 == 0 {
			evenNum++
		} else {
			oddNum++
		}
	}
	if oddNum > evenNum {
		return evenNum
	}
	return oddNum
}

func main() {
	testCases := []struct {
		position []int
	}{
		{position: []int{1, 2, 3}},
		{position: []int{2, 2, 2, 3, 3}},
		{position: []int{1, 1000000000}},
	}
	for _, testCase := range testCases {
		fmt.Println(minCostToMoveChips(testCase.position))
	}
}
