package main

import (
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var ans = "0"
	for i := len(num1); i >= 0; i-- {
		var cur = "0"
		x := int(num1[i] - '0')
		for j := len(num2); j >= 0; j-- {
			y := int(num2[j] - '0')
			product := strconv.Itoa(x * y)
			cur = addStrings(cur, product+strings.Repeat("0", len(num2)-j-1))
		}
		ans = addStrings(ans, cur+strings.Repeat("0", len(num1)-i-1))
	}

	return ans
}

func addStrings(num1, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	add := 0
	ans := ""
	for ; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		ans = strconv.Itoa(result%10) + ans
		add = result / 10
	}
	return ans
}
