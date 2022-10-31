package main

func countStudents(students []int, sandwiches []int) int {
	var (
		sum1, sum2, ans int
		n               = len(students)
	)

	for i := range students {
		sum1 += students[i]
		sum2 += sandwiches[i]
	}
	if sum1 == sum2 {
		return ans
	}

	var (
		remain int
		count  int
	)
	if sum1 < sum2 {
		remain = 1
		count = sum2 - sum1
	} else {
		count = sum1 - sum2
	}
	for i := n - 1; i >= 0; i-- {
		ans++
		if sandwiches[i] == remain {
			count--
			if count == 0 {
				break
			}
		}
	}
	return ans
}
