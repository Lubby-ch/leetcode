package main

func strStr(haystack string, needle string) int {
	n := len(haystack)
	l := len(needle)
	for i := 0; i < n; i++ {
		var j = 0
		var tmp = i
		for ;j<l && tmp < n;j++ {
			if haystack[tmp] != needle[j] {
				break
			}
			tmp++
		}
		if j == l {
			return i
		}
	}
	return -1
}
