package main

import (
	"sort"
)

func orderlyQueue(s string, k int) string {
	if k == 1 {
		var ans = s
		for i := 0; i < len(s); i++ {
			s = s[1:] + s[:1]
			if s < ans {
				ans = s
			}
		}
		return ans
	}
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
