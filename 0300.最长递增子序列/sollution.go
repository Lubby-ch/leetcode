package main

import "fmt"

func lengthOfLIS(nums []int) int {
	var dp = make([]int, len(nums))
	var ret int
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := i-1; j >=0; j-- {
			if nums[i] > nums[j]  {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
				continue
			}
		}
		if ret < dp[i] {
			ret = dp[i]
		}
	}
	return ret
}

func main() {
	testCases := []struct {
		arr []int
	}{
		{
			arr: []int{4, 10, 4, 3, 8, 9},
		},
	}
	for i := range testCases {
		testCase := testCases[i]
		fmt.Println(lengthOfLIS(testCase.arr))
	}

}
