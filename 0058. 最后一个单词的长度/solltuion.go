package main

import (
	"fmt"
	"sort"
)

func lengthOfLastWord(s string) int {
	var ans int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if ans == 0 {
				break
			}
			continue
		}
		ans++
	}
	return ans
}

func main() {
	testt := []int64{5, 4,2, 3, 1}
	sort.Slice(testt, func(i, j int) bool {
		return testt[i] < testt[j]
	})
	fmt.Println(testt)
}
