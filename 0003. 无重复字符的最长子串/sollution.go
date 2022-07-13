package main

func lengthOfLongestSubstring(s string) int {
	var maxLength, pre int
	asisstMap := make(map[byte]int)
	for i:=0; i<len(s); i++ {
		if value, ok :=asisstMap[s[i]]; ok && value+1>pre {
			pre = value + 1
		}
		if maxLength < i-pre+1 {
			maxLength = i-pre+1
		}
		asisstMap[s[i]] = i
	}
	return maxLength
}
