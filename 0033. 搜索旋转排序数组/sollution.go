package main

import "fmt"

func search1(nums []int, target int) int {
	var searchFunc func(start int, end int) int
	searchFunc = func(start int, end int) int {
		mid := (end + start) / 2
		if end < start || end == start && nums[mid] != target {
			return -1
		}
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			if nums[end] < nums[start] && nums[mid] >= nums[start] && target < nums[start] {
				return searchFunc(mid+1, end)
			} else {
				return searchFunc(start, mid-1)
			}
		} else {
			if nums[end] < nums[start] && nums[mid] <= nums[end] && target > nums[end] {
				return searchFunc(start, mid-1)
			} else {
				return searchFunc(mid+1, end)
			}
		}
	}
	return searchFunc(0, len(nums)-1)
}

func search(nums []int, target int) int {
	var (
		start, end = 0, len(nums) - 1
	)
	for {
		mid := (end + start) / 2
		if end < start || end == start && nums[mid] != target {
			return -1
		}
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			if nums[end] < nums[start] && nums[mid] >= nums[start] && target < nums[start] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			if nums[end] < nums[start] && nums[mid] <= nums[end] && target > nums[end] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
}

func main() {
	fmt.Println(search([]int{3, 1}, 1))
}
