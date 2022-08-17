package main

import "fmt"

func maxScore(s string) int {
	var max int
	var sum int
	for _, c := range s {
		if c == '1' {
			sum++
		}
	}
	if sum == len(s) {
		return sum
	}
	var count int
	for i, c := range s {
		if c == '1' {
			count++
		} else {
			if sum - count + (i+1-count)  > max {
				max = sum - count + (i+1-count)
			}
		}
	}
	return max
}

func main() {
	fmt.Println(maxScore("01"))
}

