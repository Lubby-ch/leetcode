package main

import "fmt"

func maxEqualFreq(nums []int) int {
	numFreq := make(map[int]int)
	Freq := make(map[int]int)
	var ans int
	var maxFreq int
	for i, num := range nums {
		if numFreq[num] > 0 {
			Freq[numFreq[num]]--
			if Freq[numFreq[num]] == 0 {
				delete(Freq, numFreq[num])
			}
		}
		numFreq[num]++
		Freq[numFreq[num]]++
		maxFreq = max(numFreq[num], maxFreq)
		if len(numFreq) == 1 || maxFreq == 1 || (len(Freq) == 2 && (Freq[1] == 1 || (Freq[maxFreq] == 1 && Freq[maxFreq-1] > 0))) {

			ans = max(ans, i+1)
		}
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	fmt.Println(maxEqualFreq([]int{1,2,3,1,2,3,4,4,4,4,1,2,3,5,6}))
}
