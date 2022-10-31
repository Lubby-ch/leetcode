package main

import "fmt"

func uniqueLetterString(s string) int {
	var charMap = make(map[rune][]int)
	for i, c := range s {
		arr, ok := charMap[c]
		if !ok {
			arr = []int{-1}
		}
		arr = append(arr, i)
		charMap[c] = arr
	}

	var ans int
	for _, arr := range charMap {
		arr = append(arr, len(s))
		for i := 1; i < len(arr)-1; i++ {
			ans += (arr[i] - arr[i-1]) * (arr[i+1]-arr[i])
		}
	}
	return ans
}


func main() {
	fmt.Println(uniqueLetterString("ABC"))
}