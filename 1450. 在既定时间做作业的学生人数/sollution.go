package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var n = len(startTime)
	var ans int
	for i := 0; i < n; i++ {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			ans++
		}
	}
	return ans
}