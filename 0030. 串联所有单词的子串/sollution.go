package main

// 超出时间限制
func findSubstring1(s string, words []string) []int {
	var (
		record = make(map[string]struct{})
		Func   func(arr []string, tmp string)
		lw     = len(words) * len(words[0])
		n      = len(s)
		ans    []int
	)
	Func = func(arr []string, tmp string) {
		if len(arr) == 0 {
			record[tmp] = struct{}{}
		}
		for i := 0; i < len(arr); i++ {
			cp := append([]string{}, arr...)
			cp[0], cp[i] = cp[i], cp[0]
			Func(cp[1:], tmp+cp[0])
		}
	}
	Func(words, "")
	for i := 0; i < n; i++ {
		if i+lw > n {
			break
		}
		if _, ok := record[(s[i : i+lw])]; ok {
			ans = append(ans, i)
		}
	}
	return ans
}

// 超出时间限制
func findSubstring(s string, words []string) []int {
	var (
		wordMap = make(map[string]int)
		wordLen = len(words[0])
		strLen  = len(s)
		subLen  = len(words)
		ans     []int
		sum     int
	)
	for i, word := range words {
		if _, ok := wordMap[word]; !ok {
			wordMap[word] = i + 1
		}
		sum += wordMap[word] * wordMap[word]
	}

	for i := 0; i < wordLen; i++ {
		var j, k = i, i
		var tmp int
		for ; k < strLen-wordLen+1; k += wordLen {
			if v, ok := wordMap[s[k:k+wordLen]]; ok {
				tmp += v*v
			}
			if k == (subLen-1)*wordLen+j {
				if tmp == sum {
					ans = append(ans, j)
				}
				if v, ok := wordMap[s[j:j+wordLen]]; ok {
					tmp -= v*v
				}
				j += wordLen
			}
		}
	}
	return ans
}

