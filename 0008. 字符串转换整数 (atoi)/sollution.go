package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	var ans int
	var sign int
	var i int
	for i = 0; i < len(s) && s[i] == ' '; i++ {
	}
	if i < len(s) {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
		for ; i< len(s); i++ {
			if s[i]>= '0' && s[i] <= '9' {
				ans = ans*10 + int(s[i]-'0')
			} else {
				break
			}
			if ans >  math.MaxInt32 {
				break
			}
		}
	}

	if sign == -1 {
		ans *= sign
	}
	if ans < math.MinInt32 {
		ans = math.MinInt32
	}
	if ans > math.MaxInt32 {
		ans = math.MaxInt32
	}
	return ans
}

func main() {
	testCases := []struct {
		s string
	}{
		{
			s: " -42",
		},
	}
	for i := range testCases {
		testCase := testCases[i]
		fmt.Println(myAtoi(testCase.s))
	}
}