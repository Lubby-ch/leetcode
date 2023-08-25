package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, columns            = len(matrix), len(matrix[0])
		ans                      = make([]int, rows*columns)
		index                    = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)

	for left <= right && top <= bottom {
		fmt.Println(1)
		for i := left; i <= right; i++ {
			ans[index] = matrix[top][i]
			index++
		}
		for i := top+1; i <= bottom; i++ {
			ans[index] = matrix[i][right]
			index++
		}
		if  left < right && top < bottom {
			for i := right-1; i > left; i-- {
				ans[index] = matrix[bottom][i]
				index++
			}
			for i := bottom; i > top; i-- {
				ans[index] = matrix[i][left]
				index++
			}
		}

		left++
		right--
		top++
		bottom--
	}

	return ans
}
