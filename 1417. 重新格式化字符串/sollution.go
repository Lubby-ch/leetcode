package main

import (
	"fmt"
	"math"
)

func reformat1(s string) string {
	if len(s) == 1 {
		return s
	}
	var cnum, nnum int
	n := len(s)
	for i := 0; i < n; i++ {
		if IsNum(s[i]) {
			nnum++
		} else {
			cnum++
		}
	}
	if math.Abs(float64(cnum-nnum)) > 1 {
		return ""
	}
	var first = 0
	var second = 0
	var ans string
	var sign int
	if cnum > nnum {
		sign = 1
	}
	for len(ans) < n {
		if sign == 0 {
			for first < n && !IsNum(s[first]) {
				first++
			}
			if first < n {
				ans = ans + string(s[first])
				first++
			}
		} else {
			for second < n && IsNum(s[second]) {
				second++
			}
			if second < n {
				ans = ans + string(s[second])
				second++
			}
		}
		sign ^= 1
	}
	return ans
}

func reformat(s string) string {
	var i = -1
	var j = 0
	var count int
	n := len(s)
	var ans = make([]byte, n+1)
	for _, c := range []byte(s) {
		if IsNum(c) {
			i = (i + 2) % (n+1)
			ans[i] = c
			count++
		} else {
			j = (j + 2) % (n+1)
			ans[j] = c
			count--
		}
	}
	if math.Abs(float64(count)) > 1 {
		return ""
	}
	fmt.Println(string(ans), "  ",count)
	if count >= 0 {
		return string(ans[1:])
	}
	return string(ans[:len(s)])
}

func IsNum(c byte) bool {
	return c >= '0' && c <= '9'
}
