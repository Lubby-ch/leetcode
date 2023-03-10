package main

import "fmt"

func searchInsert1(nums []int, target int) int {
	var n = len(nums)
	var start, end = 0, n - 1
	var binarySearch func() int
	binarySearch = func() int {
		mid := (start + end) / 2
		if start > end || start >= n {
			return start
		}
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
			return binarySearch()
		} else {
			start = mid + 1
			return binarySearch()
		}
	}
	return binarySearch()
}

func searchInsert(nums []int, target int) int {
	var n = len(nums)
	var start, end = 0, n - 1
	for {
		mid := (start + end) / 2
		if start > end || start >= n {
			return start
		}
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
}

func main() {
	a := []int{1}
	fmt.Println(a[3:])
}
