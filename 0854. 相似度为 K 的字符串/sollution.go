package main

import "fmt"

func kSimilarity(s1 string, s2 string) int {
	if s1 == s2 {
		return 0
	}
	target := []byte(s2)
	input := []byte(s1)
	n := len(s1)
	var dfs func(int, int)
	var minCost = n
	dfs = func(index int, cost int) {
		if index == n-1 {
			minCost = min(cost, minCost)
			return
		}
		if input[index] == target[index] {
			dfs(index+1, cost)
		} else {
			for j := index + 1; j < n; j++ {
				if input[j] == input[index] {
					continue
				}
				if input[j] == target[j] {
					continue
				}
				if input[j] == target[index] {
					input[index], input[j] = input[j], input[index]
					dfs(index+1, cost+1)
					input[index], input[j] = input[j], input[index]
				}
			}
		}
	}
	dfs(0, 0)
	return minCost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(kSimilarity("abac", "baca"))
}
