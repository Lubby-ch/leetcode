package main

import (
	"fmt"
	"unsafe"
)

func areAlmostEqual(s1 string, s2 string) bool {
	n := len(s1)
	var record []int
	for i := 0; i < n; i++ {
		if s1[i] != s2[i] {
			record = append(record, i)
			if len(record) > 2 {
				return false
			}
		}
	}
	return len(record) == 0 || (len(record) == 2 && s1[record[0]] == s2[record[1]] && s1[record[1]] == s2[record[0]])
}

func areAlmostEqual2(s1, s2 string) bool {
	var i, j = -1, -1
	for idx := range s1 {
		if s1[idx] != s2[idx] {
			if i < 0 {
				i = idx
			} else if j < 0 {
				j = idx
			} else {
				return false
			}
		}
	}
	return i < 0 || j >= 0 && s1[i] == s2[j] && s1[j] == s2[i]
}


func main() {
	s := make([]int, 9, 20)
	var Len = unsafe.Pointer(&s)
	fmt.Println(Len, len(s)) // 9 9
}
