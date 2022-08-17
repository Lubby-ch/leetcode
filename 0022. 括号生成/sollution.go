package main

import "fmt"

func generateParenthesis(n int) []string {
	return generate(n, n)
}

func generate(right, left int) []string {
	var ans []string
	if left == 0 && right == 0 {
		return []string{""}
	}
	if left == right {
		strs := generate(right-1, left)
		for _, str := range strs {
			ans = append(ans, "("+str)
		}
	} else if left > right {
		if right > 0 {
			strs := generate(right-1, left)
			for _, str := range strs {
				ans = append(ans, "("+str)
			}
		}
		if left > 0 && right >= 0 {
			strs := generate(right, left-1)
			for _, str := range strs {
				ans = append(ans, ")"+str)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(generateParenthesis(3))
}
