package main

import "fmt"

func isMatch1(s string, p string) bool { // 超出时间 可以借用slice存储状态 优化
	var dp [][]int
	m := max(len(s), len(p))
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, m+1))
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if len(s) == i || len(p) == j {
			if len(s) == i && len(p) == j {
				if dp[i][j] == 0 {
					dp[i][j] = 1
				}
				return dp[i][j]
			} else if len(p) != j {
				if p[j] == '*' {
					if dp[i][j+1] == 0 {
						dp[i][j+1] = dfs(i, j+1)
					}
					return dp[i][j+1]
				}
			}
		} else {
			if s[i] == p[j] || p[j] == '?' {
				if dp[i+1][j+1] == 0 {
					dp[i+1][j+1] = dfs(i+1, j+1)
				}
				return dp[i+1][j+1]
			}
			if p[j] == '*' {
				if dp[i][j+1] == 0 {
					dp[i][j+1] = dfs(i, j+1)
				}
				if dp[i][j+1] == 1 {
					return dp[i][j+1]
				}
				if dp[i+1][j] == 0 {
					dp[i+1][j] = dfs(i+1, j)
				}
				return dp[i+1][j]
			}
		}
		dp[i][j] = -1
		return dp[i][j]
	}
	return dfs(0, 0) == 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isMatch(s string, p string) bool { // 超出时间 可以借用slice存储状态 优化
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}

func main() {
	fmt.Println(isMatch("aaab", "a*"))
}
