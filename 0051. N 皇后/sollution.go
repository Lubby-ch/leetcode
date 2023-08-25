package main

func solveNQueens(n int) [][]string {
	var ret [][]string
	var direct1 = make(map[int]bool)
	var direct2 = make(map[int]bool)
	var direct3 = make(map[int]bool)

	var ans = make([]int, 0, n)
	var dfs func(int)
	dfs = func(idx int) {
		if idx == n {
			ret = append(ret, generateAns(ans, n))
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
			ans = append(ans, i)

			dfs(idx + 1)

			ans = ans[:idx]
			delete(direct1, i)
			delete(direct2, idx-i)
			delete(direct3, idx+i)
		}
	}
	dfs(0)
	return ret
}

func generateAns(ans []int, n int) []string {
	var ret []string
	for i := 0; i < n; i++ {
		var str = make([]byte, n)
		for j := 0; j < n; j++ {
			str[j] = '.'
		}
		str[ans[i]] = 'q'
		ret = append(ret, string(str))
	}
	return ret
}

func main() {
	solveNQueens(10)
}
