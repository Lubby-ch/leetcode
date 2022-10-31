package main

var next = [4][2]int{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1},
}

func shortestBridge(grid [][]int) int {
	n := len(grid)
	var container [][2]int
	var dfs func(x, y int)
	dfs = func(x, y int) {
		grid[x][y] = 2
		var save bool
		for _, v := range next {
			newX, newY := x+v[0], y+v[1]
			if newX < 0 || newX >= n || newY < 0 || newY >= n {
				continue
			}
			if grid[newX][newY] == 1 {
				dfs(newX, newY)
			} else if grid[newX][newY] == 0 {
				save = true
			}
		}
		if save {
			container = append(container, [2]int{x, y})
		}
	}
	for i := 0; i < n; i++ {
		var j int
		for j = 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				break
			}
		}
		if j < n {
			break
		}
	}
	var bfs func(ans int) int
	bfs = func(ans int) int {
		length := len(container)
		for i := 0; i < length; i++ {
			node := container[i]
			for _, v := range next {
				newX, newY := node[0]+v[0], node[1]+v[1]
				if newX < 0 || newX >= n || newY < 0 || newY >= n {
					continue
				}
				if grid[newX][newY] == 1 {
					return ans
				} else if grid[newX][newY] == 0 {
					grid[newX][newY] = 2
					container = append(container, [2]int{newX, newY})
				}
			}
		}
		return bfs(ans + 1)
	}
	return bfs(0)
}
