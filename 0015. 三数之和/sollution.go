package main

import (
	"fmt"
	"sort"
)

/*func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	var ans [][]int
	var numMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = i
	}
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if 0-nums[i]-nums[j] < nums[i] {
				break
			}
			v, ok := numMap[0-nums[i]-nums[j]]
			if !ok {
				continue
			}
			if v <= j {
				break
			}
			if len(ans) > 0 && ans[len(ans)-1][0]==nums[i] && ans[len(ans)-1][1]==nums[j] {
				continue
			}
			ans = append(ans, []int{nums[i], nums[j], 0-nums[i]-nums[j]})
		}
	}
	var index = 1
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0] || (ans[i][0] == ans[j][0] && ans[i][1] < ans[j][1])
	})
	for {
		if index >= len(ans) {
			break
		}
		if ans[index][0] == ans[index-1][0] && ans[index][1] == ans[index-1][1] {
			ans = append(ans[:index], ans[index+1:]...)
			continue
		}
		index++
	}
	return ans
}
*/
// 排序 + 双指针
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		var third = n-1
		for second := first +1;second<n;second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for nums[third] > -nums[second]-nums[first] && third > second {
				third --
			}
			if third <= second {
				break
			}
			if nums[third] < -nums[second]-nums[first] {
				continue
			}
			ans = append(ans, []int{nums[first], nums[second], -nums[first]-nums[second]})
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		arr []int
	}{
		{
			arr: []int{-1,0,1,2,-1,-4},
		},
	}
	for i := range testCases {
		testCase := testCases[i]
		fmt.Println(threeSum(testCase.arr))
	}
}