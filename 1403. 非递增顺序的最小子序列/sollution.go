package main

import (
	"fmt"
	"sort"
)

func minSubsequence(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	var head = 0
	var tail = len(nums) - 1
	var sum1 = nums[head]
	var sum2 = nums[tail]
	for {
		if head == tail {
			break
		}
		if sum1 <= sum2 {
			head++
			sum1 += nums[head]
		} else if sum1 > sum2 {
			tail--
			sum2 += nums[tail]
		}
	}
	return nums[:head+1]
}

func main() {
	testCases := []struct {
		arr []int
	}{
		{
			arr: []int{10,2,5},
		},
	}
	for i := range testCases {
		testCase := testCases[i]
		fmt.Println(minSubsequence(testCase.arr))
	}
}
