package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]int
	var grid [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			v := (board[i][j] - '0') - 1
			fmt.Println(v)
			if rows[i][v] != 0 {
				return false
			} else {
				rows[i][v] = 1
			}

			if columns[j][v] != 0 {
				return false
			} else {
				columns[j][v] = 1
			}
			if grid[getGridNum(i, j)][v] != 0 {
				return false
			} else {
				grid[getGridNum(i, j)][v] = 1
			}
		}
	}

	return true
}

func getGridNum(i, j int) int {
	return (j/3)*3 + i/3
}
