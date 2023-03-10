package main

func solveSudoku(board [][]byte) {
	var rows, columns [9][9]bool
	var grid [9][9]bool
	var space [][2]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				space = append(space, [2]int{i, j})
				continue
			}
			v := board[i][j] - '1'
			rows[i][v] = true
			columns[j][v] = true
			grid[getGridNum(i, j)][v] = true
		}
	}
	var dfs func(index int) bool
	dfs = func(index int) bool {
		if index == len(space) {
			return true
		}
		i := space[index][0]
		j := space[index][1]
		for v := byte(0); v < 9; v++ {
			if rows[i][v] || columns[j][v] || grid[getGridNum(i, j)][v] {
				continue
			}
			rows[i][v] = true
			columns[j][v] = true
			grid[getGridNum(i, j)][v] = true
			board[i][j] = v
			if dfs(index + 1) {
				return true
			}
			rows[i][v] = false
			columns[j][v] = false
			grid[getGridNum(i, j)][v] = false
			board[i][j] = '.'
		}
		return false
	}
	dfs(0)
}

func getGridNum(i, j int) int {
	return (j/3)*3 + i/3
}
