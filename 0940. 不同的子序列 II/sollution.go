package main

// 动态规划
func distinctSubseqII(s string) int {
	const mod int = 1e9 + 7
	n := len(s)
	last := make([]int, 26)
	for i := 0; i < 26; i++ {
		last[i] = -1
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i, c := range s {
		for j := range last {
			if last[j] != -1 {
				dp[i] = (dp[last[j]] + dp[i]) % mod
			}
		}
		last[c-'a'] = i
	}
	var ans int
	for i := range last {
		if last[i] != -1 {
			ans = (ans + dp[last[i]]) % mod
		}

	}
	return ans
}

// 动态规划-- 优化
func distinctSubseqII1(s string) int {
	const mod int = 1e9 + 7
	last := make([]int, 26)
	var cur int
	for _, c := range s {
		var lastRecord = cur + 1
		cur = ((cur%mod + (cur+1)%mod - last[c-'a']%mod) + mod) % mod
		last[c-'a'] = lastRecord
	}
	return cur
}
