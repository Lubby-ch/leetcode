package main

import "fmt"

func twoSum(nums []int, target int) []int {
	container := make(map[int]int)
	var (
		ok    bool
		value int
	)
	for i := 0; i < len(nums); i++ {
		if value, ok = container[target-nums[i]]; ok && value != i {
			return []int{i, value}
		}
		container[nums[i]] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{-1, -2, -3, -4, -5}, -8))
}
