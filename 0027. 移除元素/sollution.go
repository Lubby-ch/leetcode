package main

func removeElement(nums []int, val int) int {
	var i int
	n := len(nums)
	for j := 0; j < n; j++ {
		if nums[j] != val {
			if i != j {
				nums[i] = nums[j]
			}
			i++
		}
	}
	return i
}
