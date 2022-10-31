package main

func removeDuplicates(nums []int) int {
	var (
		i = 0
		j = 1
		tmp = nums[0]
	)
	n := len(nums)
	for j<n {
		if nums[j] != tmp {
			i++
			tmp = nums[j]
			nums[i] = tmp
		}
		j++
	}
	return len(nums)
}
