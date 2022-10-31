package main

func totalFruit(fruits []int) int {
	var first, second = 0, -1
	var ans = 0
	var cur = 1
	for i := 1; i < len(fruits); i++ {
		if fruits[i] == fruits[i-1] {
			cur++
		} else if second == -1 {
			second = i
			cur++
		} else if fruits[i] == fruits[first] || fruits[i] == fruits[second] {
			first = second
			second = i
			cur++
		} else {
			first = second
			second = i
			cur = i - first + 1
		}
		ans = max(ans, cur)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
