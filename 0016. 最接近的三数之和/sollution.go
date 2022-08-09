package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var ans = nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		var third = len(nums) - 1
		var second = i + 1
		for second < third {
			sum := nums[third] + nums[i] + nums[second]
			if math.Abs(float64(target-ans)) < math.Abs(float64(target-ans)) {
				ans = sum
			}
			if target < ans {
				third--
			} else if target > ans {
				second++
			} else {
				return ans
			}
		}
	}
	return ans
}

func main() {
	testCases := []struct {
		nums   []int
		target int
	}{
		{
			nums:   []int{4, 0, 5, -5, 3, 3, 0, -4, -5},
			target: -2,
		},
	}
	for i := range testCases {
		testCase := testCases[i]
		fmt.Println(threeSumClosest(testCase.nums, testCase.target))
	}
}
