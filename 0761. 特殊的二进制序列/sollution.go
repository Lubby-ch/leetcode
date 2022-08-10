package main

import (
	"sort"
	"strings"
)

func makeLargestSpecial(s string) string {
	if len(s) <= 2 {
		return s
	}
	var subs sort.StringSlice
	var count int
	var ptr int
	for i, c := range s {
		if c == '1' {
			count++
		} else {
			count--
			if count == 0 {
				subs = append(subs, "1"+makeLargestSpecial(s[ptr+1:i])+"0")
				ptr = i + 1
			}
		}
	}
	sort.Sort(sort.Reverse(subs))
	return strings.Join(subs, "")
}
