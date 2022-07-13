package main

import (
	"fmt"
)

func asteroidCollision(asteroids []int) []int {
	var result []int
	for _, asteroid := range asteroids {
		var ispush = true
		// 判断进出stack
		for len(result) > 0 {
			if result[len(result)-1] > 0 && asteroid < 0 {
				if 0-asteroid > result[len(result)-1] {
					result = result[:len(result)-1]
					continue
				} else if 0-asteroid == result[len(result)-1] {
					result = result[:len(result)-1]
				}
				ispush = false
				break
			}
			break
		}
		if ispush {
			result = append(result, asteroid)
		}
	}
	return result
}

func main() {
	testCases := []struct {
		asteroids []int
	}{
		{
			asteroids: []int{5, 10, -5},
		},
	}
	for _, testcase := range testCases {
		fmt.Println(asteroidCollision(testcase.asteroids))
	}
}
