package main

func movesToChessboard(board [][]int) int {
	n := len(board)

	var r1,c1 int

	var rowSwap int
	var colSwap int
	var sign = 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[0][0]^board[i][0]^board[0][j]^board[i][j] != 0 {
				return -1
			}
		}

		if board[0][i]^sign == 1 {
			rowSwap++
		}
		if board[i][0]^sign == 1 {
			colSwap++
		}
		sign ^= 1

		if board[0][i] == 1 {
			r1++
		}
		if board[i][0] == 1 {
			c1++
		}
	}

	if r1 != n / 2 && r1 != (n+1)/2 {
		return -1
	}

	if c1 != n / 2 && c1 != (n+1)/2 {
		return -1
	}

	if n%2 == 1 {
		if rowSwap%2 == 1 {
			rowSwap = n - rowSwap
		}
		if colSwap%2 == 1 {
			colSwap = n - colSwap
		}
	} else {
		rowSwap = min(rowSwap, n-rowSwap)
		colSwap = min(colSwap, n-colSwap)
	}
	return (rowSwap+colSwap) / 2
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}