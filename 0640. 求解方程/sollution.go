package main

import "fmt"

func solveEquation(equation string) string {
	var coeff int
	var value int
	var isreverse bool
	var tmp int
	var sign = 1
	for i, char := range equation {
		if char == 'x' {
			if tmp == 0 {
				if i>0 && equation[i-1] =='0' {
					tmp = 0
				} else {
					tmp = 1
				}
			}
			if isreverse {
				coeff -= sign * tmp
			} else {
				coeff += sign * tmp
			}
			tmp = 0
		} else if char >= '0' && char <= '9' {
			tmp = tmp*10 + int(char-'0')
		} else {
			if tmp != 0 {
				if isreverse {
					value -= sign * tmp
				} else {
					value += sign * tmp
				}
				tmp = 0
			}
			if char == '=' {
				isreverse = true
				sign = 1
			}
			if char == '-' {
				sign = -1
			} else if char == '+' {
				sign = 1
			}
		}
		if i == len(equation)-1 && tmp != 0 {
			if isreverse {
				value -= sign * tmp
			} else {
				value += sign * tmp
			}
			tmp = 0
		}
	}
	if coeff == 0 && value == 0 {
		return "Infinite solutions"
	} else if coeff == 0 && value != 0 {
		return "No solution"
	}
	return fmt.Sprintf("x=%d", -1*value/coeff)
}
