package main

import (
	"runtime"
	"sort"
	"time"
)

// 暴力 超出时间限制
func advantageCount1(nums1 []int, nums2 []int) []int {
	var n = len(nums1)
	var ans = make([]int, n)
	var maxCount int
	var dfs func(index int, count int)
	dfs = func(index int, count int) {
		if index == n && count > maxCount {
			maxCount = count
			copy(ans, nums1)
			return
		}
		for i := index; i < n; i++ {
			nums1[index], nums1[i] = nums1[i], nums1[index]
			if nums1[index] > nums2[index] {
				dfs(index+1, count+1)
			} else {
				dfs(index+1, count)
			}
			nums1[index], nums1[i] = nums1[i], nums1[index]
		}
	}
	dfs(0, 0)
	return ans
}

// 田忌赛马
func advantageCount(nums1 []int, nums2 []int) []int {
	var n = len(nums1)

	var arr1 = make([]int, n)
	var arr2 = make([]int, n)
	for i := 0; i < n; i++ {
		arr1[i] = i
		arr2[i] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		return nums1[arr1[i]] < nums1[arr1[j]]
	})
	sort.Slice(arr2, func(i, j int) bool {
		return nums2[arr2[i]] < nums2[arr2[j]]
	})
	var ans = make([]int, n)
	var left, right = 0, n - 1
	for i := 0; i < n; i++ {
		if nums1[arr1[i]] > nums2[arr2[left]] {
			ans[arr2[left]] = nums1[arr1[i]]
			left++
		} else {
			ans[arr2[right]] = nums1[arr1[i]]
			right--
		}
	}
	return ans
}

// 田忌赛马
func advantageCount2(nums1 []int, nums2 []int) []int {
	var n = len(nums1)
	var arr2 = make([]int, n)
	for i := 0; i < n; i++ {
		arr2[i] = i
	}
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(arr2, func(i, j int) bool {
		return nums2[arr2[i]] < nums2[arr2[j]]
	})
	var ans = make([]int, n)
	var left, right = 0, n - 1
	for i := 0; i < n; i++ {
		if nums1[i] > nums2[arr2[left]] {
			ans[arr2[left]] = nums1[i]
			left++
		} else {
			ans[arr2[right]] = nums1[i]
			right--
		}
	}
	return ans
}

func main() {
	go func() {
		for {
		}
	}()

	time.Sleep(time.Millisecond)
	runtime.GC()
	println("OK")
}
