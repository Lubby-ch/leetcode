package main

func scoreOfParentheses(s string) int {
	var (
		ans   int
		stack []int
		n     = len(s)
	)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, 0)
		} else {
			var cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if cur == 0 {
				cur = 1
			} else {
				cur *= 2
			}
			if len(stack) == 0 {
				ans += cur
			} else {
				stack[len(stack)-1] += cur
			}
		}
	}
	return ans
}
