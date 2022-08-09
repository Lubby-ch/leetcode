package main

func minStartValue(nums []int) int {
	var sum int
	var min int
	for i := 0; i<len(nums); i++ {
		sum += nums[i]
		if sum < min {
			min = sum
		}
	}
	if min >= 0 {
		return 1
	} else {
		return 1-min
	}
}
