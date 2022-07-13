package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var iseven bool
	var mid float64
	tag := (len(nums1) + len(nums2)) / 2
	if (len(nums1)+len(nums2))%2 == 0 {
		iseven = true
		tag -= 1
	}
	for tag > 0 {
		var count1, num1, count2, num2 int
		if len(nums1) > 0 {
			count1, num1 = GetPosNum(nums1, tag/2)
		}
		if len(nums2) > 0 {
			count2, num2 = GetPosNum(nums2, tag/2)
		}
		if count1 > 0 && count2 > 0 {
			if num1 < num2 {
				nums1 = nums1[count1:]
				tag -= count1
			} else {
				nums2 = nums2[count2:]
				tag -= count2
			}
		} else if count1 > 0 {
			nums1 = nums1[count1:]
			tag -= count1
		} else {
			nums2 = nums2[count2:]
			tag -= count2
		}
	}
	if len(nums1) > 0 && len(nums2) > 0 {
		if nums1[0] < nums2[0] {
			mid = float64(nums1[0])
			nums1 = nums1[1:]
			if iseven {
				if len(nums1) == 0 || nums1[0] > nums2[0] {
					mid += float64(nums2[0])
				} else {
					mid += float64(nums1[0])
				}
				mid /= 2
			}
		} else {
			mid = float64(nums2[0])
			nums2 = nums2[1:]
			if iseven {
				if len(nums2) == 0 || nums1[0] < nums2[0] {
					mid += float64(nums1[0])
				} else {
					mid += float64(nums2[0])
				}
				mid /= 2
			}
		}
	} else if len(nums1) > 0 {
		mid += float64(nums1[0])
		if iseven {
			mid += float64(nums1[1])
			mid /= 2
		}
	} else {
		mid += float64(nums2[0])
		if iseven {
			mid += float64(nums2[1])
			mid /= 2
		}
	}
	return mid
}

func GetPosNum(nums []int, pos int) (n int, num int) {
	if pos == 0 {
		pos = 1
	}
	if pos < len(nums) {
		return pos, nums[pos-1]
	}
	return len(nums), nums[len(nums)-1]
}

func main() {
	testCases := []struct {
		nums1 []int
		nums2 []int
	}{
		{
			nums1: []int{1, 2,3},
			nums2: []int{1,2},
		},
	}
	for i := range testCases {
		testCase := testCases[i]
		fmt.Println(findMedianSortedArrays(testCase.nums1, testCase.nums2))
	}

}
