package main

func isValid(s string) bool {
	var stack []byte
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, byte(c))
			continue
		} else if c == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
		} else if c == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
		} else if c == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
		}
		stack = stack[:len(stack)-1]
	}
	if len(stack) > 0 {
		return false
	}
	return true
}