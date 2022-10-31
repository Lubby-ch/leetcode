package main

import (
	"fmt"
	"math"
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	if n <= k {
		return arr
	}
	// top-k 算法
	sort.Ints(arr)
	var i, j int
	for i = k; i < n; i++ {
		if math.Abs(float64(arr[i]-x)) < math.Abs(float64(arr[j+k]-x)) || math.Abs(float64(arr[i]-x)) < math.Abs(float64(arr[j]-x)) {
			j = i - k + 1
		}
	}
	return arr[j : j+k]
}

func main() {
	fmt.Println(findClosestElements([]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5))
}
