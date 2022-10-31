package main

import (
	"fmt"
	"sort"
)

func canBeEqual(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)
	if len(target) != len(target) {
		return false
	}
	n := len(target)
	for i:=0;i<n;i++{
		if target[i] != arr[i] {
			return false
		}
	}
	return true
}


