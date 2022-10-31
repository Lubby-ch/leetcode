package main

func shortestSubarray(nums []int, k int) int {
	var stack [][2]int
	stack = append(stack, [2]int{0,0})
	var sum int
	var ans = len(nums)+1
	for i := range nums {
		sum += nums[i]
		for len(stack) > 0 && sum <= stack[len(stack)-1][0] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, [2]int{sum, i+1})
	}
	for i := 0;i<len(stack);i++{
		var find bool
		for j:=i+1;j<len(stack);j++{
			if stack[j][0] - stack[i][0] >= k {
				ans = min(ans, stack[j][1] - stack[i][1])
				find = true
			}
		}
		if !find {
			break
		}
	}
	if ans > len(nums) {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
func main() {
	shortestSubarray([]int{17,85,93,-45,-21}, 150)
}