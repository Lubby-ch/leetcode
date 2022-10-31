package main

import (
	"fmt"
	"math"
)

// 两层遍历
func decrypt1(code []int, k int) []int {
	n := len(code)
	if k == 0 {
		return make([]int, n)
	}
	var ans = make([]int, 0, len(code))
	for i := 0; i < n; i++ {
		var (
			index = i
			dec   int
		)
		if k > 0 {
			index = i + k + 1
		}
		for j := 1; j <= int(math.Abs(float64(k))); j++ {
			index--
			if index < 0 {
				index += n
			}
			index = index % n
			dec += code[index]
		}
		ans = append(ans, dec)
	}
	return ans
}

func decrypt(code []int, k int) []int {
	n := len(code)
	var (
		ans = make([]int, n)
		sum = 0
	)
	if k == 0 {
		return ans
	}
	var left, right = 1, k
	if k < 0 {
		right = n - 1
		left = n + k
	}
	for i := left; i <= right; i++ {
		sum += code[(i+n)%n]
		fmt.Println(sum)
	}

	for i := 0; i < n; i++ {
		ans[i] = sum
		sum -= code[left%n]
		sum += code[(right+1)%n]
		left++
		right++
	}
	return ans
}

func main() {
	fmt.Println(decrypt([]int{1, 2, 3, 4}, -3))
}
