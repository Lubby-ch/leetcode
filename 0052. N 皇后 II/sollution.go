package main

func totalNQueens(n int) int {
	var ans int
	var direct1 = make(map[int]bool)
	var direct2 = make(map[int]bool)
	var direct3 = make(map[int]bool)

	var dfs func(int)
	dfs = func(idx int) {
		if idx == n {
			ans++
			return
		}
		for i := 0; i < n; i++ {
			if direct1[i] {
				continue
			}
			if direct2[idx-i] {
				continue
			}
			if direct3[idx+i] {
				continue
			}
			direct1[i] = true
			direct2[idx-i] = true
			direct3[idx+i] = true

			dfs(idx + 1)

			delete(direct1, i)
			delete(direct2, idx-i)
			delete(direct3, idx+i)
		}
	}
	dfs(0)
	return ans
}

