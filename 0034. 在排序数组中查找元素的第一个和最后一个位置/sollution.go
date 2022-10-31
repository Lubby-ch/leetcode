package main

func searchRange(nums []int, target int) []int {
	var (
		n            = len(nums)
		binarySearch func(start, end int, flag bool) int
		ans          = make([]int, 0, 2)
	)

	binarySearch = func(start, end int, flag bool) int {
		mid := (start + end) / 2

		if end < start {
			return -1
		}

		if nums[mid] < target {
			return binarySearch(mid+1, end, flag)
		} else if nums[mid] > target {
			return binarySearch(start, mid-1, flag)
		}

		if flag {
			if mid == start || nums[mid-1] != target {
				return mid
			}
			return binarySearch(start, mid-1, flag)
		} else {
			if mid == end || nums[mid+1] != target {
				return mid
			}
			return binarySearch(mid+1, end, flag)
		}
	}

	ans = append(ans, binarySearch(0, n, true))
	ans = append(ans, binarySearch(0, n, false))

	return ans
}
