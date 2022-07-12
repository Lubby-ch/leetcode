package main

import "fmt"

func lenLongestFibSubseq(arr []int) int {
	dp := make([][]int, len(arr))
	for i := range dp {
		dp[i] = make([]int, len(arr))
	}
	arrMap := make(map[int]int)
	var ret = 0
	for i := 0; i < len(arr); i++ {
		arrMap[arr[i]] = i
		for j := i - 1; j > 0; j-- {
			if arr[j] < arr[i]/2 {
				break
			}
			tmp := arr[i] - arr[j]
			if value, ok := arrMap[tmp]; ok {
				if value >= j {
					break
				}
				dp[i][j] = dp[j][value] + 1
				fmt.Println(arr[i],",",arr[j])
				if dp[i][j] < 3 {
					dp[i][j] = 3
				}
				if ret < dp[i][j] {
					ret = dp[i][j]
				}
			}
		}
	}
	return ret
}

func main() {
	testCases := []struct {
		arr []int
	}{
		{
			arr: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			arr: []int{2392,2545,2666,5043,5090,5869,6978,7293,7795},
		},
	}
	for i := range testCases {
		testCase := testCases[i]
		fmt.Println(lenLongestFibSubseq(testCase.arr))
	}

}
