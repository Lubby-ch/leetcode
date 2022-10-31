package main

import (
	"math"
	"strconv"
)

// 暴力--超时
func atMostNGivenDigitSet1(digits []string, n int) int {
	length := len(digits)
	var (
		dfs func(value int)
		ans int
	)
	dfs = func(value int) {
		value *= 10
		if value > n {
			return
		}
		for i := 0; i < length; i++ {
			if value+int(digits[i][0]-'0') <= n {
				ans++
				dfs(value + int(digits[i][0]-'0'))
				continue
			}
			return
		}
	}
	dfs(0)
	return ans
}

func atMostNGivenDigitSet(digits []string, n int) int {
	var (
		stack  []int
		ans    int
		length = len(digits)
	)
	for n != 0 {
		stack = append(stack, n%10)
		n /= 10
	}
	var base int
	if length == 1 {
		base = 1
	} else {
		base = length * int(math.Pow(float64(length), float64(len(stack)-1))-1) / (length - 1)
	}
	var i = 0
	ans += base
	for ; i < length; i++ {
		if int(digits[i][0]-'0') < stack[len(stack)-1] {
			continue
		}
		if int(digits[i][0]-'0') == stack[len(stack)-1] {
			ans += atMostNGivenDigitSet(digits, n-stack[len(stack)-1]*int(math.Pow(10, float64(len(stack)-1))))
		}
		break
	}
	ans += i * base
	return ans
}

// 动态规划
func atMostNGivenDigitSet2(digits []string, n int) int {
	var (
		stack  = strconv.Itoa(n)
		length = len(stack)
		dp     = make([][2]int, length+1)
	)
	dp[0][1] = 1
	for i := 1; i <= length; i++ {
		for _, v := range digits {
			if stack[i-1] > v[0] {
				dp[i][0] += dp[i-1][1]
			} else if stack[i-1] == v[0] {
				dp[i][1] = dp[i-1][1]
			} else {
				break
			}
		}
		if i > 1 {
			dp[i][0] += len(digits) + dp[i-1][0] * len(digits)
		}
	}
	return dp[length][0] + dp[length][1]
}
