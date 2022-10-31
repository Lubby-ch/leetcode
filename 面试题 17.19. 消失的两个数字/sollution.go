package main

func missingTwo(nums []int) []int {
	var nor = 0
	var n = len(nums)
	for i := 0; i < n; i++ {
		nor ^= nums[i]
		nor ^= i + 1
	}
	for i := n + 1; i <= n+2; i++ {
		nor ^= i
	}

	diff := nor & -nor
	var x1, x2 int
	for i := 0; i < n; i++ {
		if nums[i]&diff == 0 {
			x1 ^= nums[i]
		} else {
			x2 ^= nums[i]
		}
		if (i+1)&diff == 0 {
			x1 ^= i + 1
		} else {
			x2 ^= i + 1
		}
	}
	for i := n + 1; i <= n+2; i++ {
		if i&diff == 0 {
			x1 ^= i
		} else {
			x2 ^= i
		}
	}
	return []int{x1, x2}
}



