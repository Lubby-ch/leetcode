package main

import (
	"fmt"
	"math"
	"sort"
)
// 纯数学问题， 本方法需要排序， 通过边判断则不需要排序
// 正方形对角线垂直平分且相等
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	list := [][]int{p1, p2, p3, p4}
	sort.Slice(list, func(i, j int) bool {
		return list[i][0] < list[j][0] || (list[i][0] == list[j][0] && list[i][1]< list[j][1])
	})
	// 判断对角线是否垂直
	if !DiagIsVertical([]int{list[0][0]-list[3][0],list[0][1]-list[3][1]},
		[]int{list[1][0]-list[2][0],list[1][1]-list[2][1]}) {
		return false
	}

	// 判断点到直线的距离是否相等
	if DistPointTOLine(list[1], list[2], list[0]) != DistPointTOLine(list[1], list[2], list[3]) {
		return false
	}
	if DistPointTOLine(list[0], list[3], list[1]) != DistPointTOLine(list[0], list[3], list[2]) {
		return false
	}
	if DistPointTOLine(list[0], list[3], list[1]) != DistPointTOLine(list[1], list[2], list[0]) {
		return false
	}
	return true
}

func DiagIsVertical(a []int, b []int) bool {
	return a[0]*b[0] + a[1]*b[1] == 0
}

func DistPointTOLine(a []int, b []int, c []int) float64 {
	A := b[1] - a[1]
	B := a[0] - b[0]
	C := b[0]*a[1] - a[0]*b[1]
	return math.Abs(float64(A*c[0]+B*c[1]+C)) / math.Pow(float64(A*A+B*B), 0.5)
}

func main () {
	fmt.Println(1<<8)
	fmt.Println(258>>8)
	fmt.Println(258&0xff)
}
