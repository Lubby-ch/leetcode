package main

func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	var ret = make([][]int, m)
	for i:=0;i<m;i++ {
		ret[i] = make([]int, n)
	}
	for i:=0;i<m;i++ {
		for j:=0;j<n;j++ {
			index := (i*m+j+k) % (n*m)
			ret[index/n][index%n]=grid[i][j]
		}
	}
	return ret
}