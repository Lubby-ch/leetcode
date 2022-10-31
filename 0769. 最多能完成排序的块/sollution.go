package main

import (
	"fmt"
	rand2 "math/rand"
	"time"
)

func maxChunksToSorted(arr []int) int {
	n := len(arr)
	var ans int
	var record = -1
	for i := 0; i < n; i++ {
		if arr[i] > i {
			record = max(record, arr[i])
		} else if arr[i] == i {
			if record == -1 {
				ans++
			}
		} else {
			if record == i {
				ans++
				record = -1
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rand(min, max float64) float64 {
	return min + rand2.Float64()*(max-min)
}

func main() {
	rand2.Seed(time.Now().Unix())
	var a, b int
	for i := 0; i < 10000000; i++ {
		if rand(-1, 1) > 0 {
			a++
		}else{
			b++
		}
	}
	fmt.Println(a)
	fmt.Println(b)
}
