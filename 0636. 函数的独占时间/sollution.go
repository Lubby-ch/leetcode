package main

import (
	"strconv"
	"strings"
)

type exeTime struct {
	id        int
	starttime int
}

func exclusiveTime(n int, logs []string) []int {
	ans := make([]int, n)
	stack := make([]*exeTime, 0, n)
	for _, log := range logs {
		strs := strings.Split(log, ":")
		if len(strs) != 3 {
			return nil
		}
		if strs[1] == "start" { // 入栈
			tmp := &exeTime{}
			tmp.id, _ = strconv.Atoi(strs[0])
			tmp.starttime, _ = strconv.Atoi(strs[2])
			if len(stack) > 0 {
				ans[stack[len(stack)-1].id] += (tmp.starttime - stack[len(stack)-1].starttime)
			}
			stack = append(stack, tmp)
		} else { // 出栈
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			endtime, _ := strconv.Atoi(strs[2])
			ans[tmp.id] += (endtime - tmp.starttime)
			if len(stack) > 0 {
				stack[len(stack)-1].starttime, _ = strconv.Atoi(strs[2])
			}
		}
	}
	return ans
}
