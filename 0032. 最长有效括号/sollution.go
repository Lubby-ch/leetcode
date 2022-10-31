package main

import (
	"context"
	"fmt"
	"time"
)

// 动态规划
func longestValidParentheses(s string) int {
	n := len(s)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		if s[i] == '(' {
			continue
		}
		if s[i-1] == '(' {
			if i-2 >= 0 {
				dp[i] = dp[i-2] + 2
			} else {
				dp[i] = 2
			}
		} else {
			if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				dp[i] = dp[i-1] + 2
				if i-dp[i-1]-2 >= 0 {
					dp[i] += dp[i-dp[i-1]-2]
				}
			}
		}
	}
	var ans int
	for i := 0; i < n; i++ {
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

// 栈
func longestValidParentheses1(s string) int {
	n := len(s)
	var stack = make([]int, 0, n)
	var ans int
	for i := 0; i < n; i++ {
		if len(stack) == 0 || s[i] == '(' || s[stack[len(stack)-1]] == ')' {
			stack = append(stack, i)
			continue
		}
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			ans = i+1
		} else {
			ans = max(ans, i-stack[len(stack)-1])
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type MyContext struct {
	// 这里的 Context 是我 copy 出来的，所以前面不用加 context.
	context.Context
}

func main() {
	childCancel := true

	parentCtx, parentFunc := context.WithCancel(context.Background())
	mctx := MyContext{parentCtx}

	childCtx, childFunc := context.WithCancel(mctx)

	if childCancel {
		childFunc()
	} else {
		parentFunc()
	}

	fmt.Println(parentCtx)
	fmt.Println(mctx)
	fmt.Println(childCtx)

	// 防止主协程退出太快，子协程来不及打印
	time.Sleep(10 * time.Second)
}