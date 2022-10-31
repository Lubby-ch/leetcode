package main

import (
	"fmt"
	"reflect"
)

func constructArray(n int, k int) []int {
	var ans []int
	for i := 1; i < n-k; i++ {
		ans = append(ans, i)
	}
	for i, j := n-k, n; i <= j; i++ {
		ans = append(ans, i)
		if j > i {
			ans = append(ans, j)
		}
		j--
	}
	return ans
}

func main () {
	str := "go语言9"
	for i, s := range str {
		fmt.Println(i,",",s,",",reflect.TypeOf(s))
	}
}