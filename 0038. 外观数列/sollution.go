package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	var ans string
	str := countAndSay(n - 1)
	var i, j int
	for i < len(str) {
		if j == len(str) || str[j] != str[i] {
			ans += strconv.Itoa(j-i) + string(str[i])
			if j == len(str) {
				return ans
			}
			i = j
		}
		j++
	}
	return ans
}

func countAndSay1(n int) string {
	var str = "1"
	var ans = "1"
	for count := 0; count < n; count++ {
		ans = ""
		var i, j int
		for i < len(str) {
			if j == len(str) || str[j] != str[i] {
				ans += strconv.Itoa(j-i) + string(str[i])
				if j == len(str) {
					break
				}
				i = j
			}
			j++
		}
		str = ans
	}
	return ans
}

func main() {
	fmt.Println(countAndSay1(4))
}
