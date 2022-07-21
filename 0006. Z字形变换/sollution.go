package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	t := 2*numRows - 2
	n := len(s)
	ret := make([]byte, 0, n)
	for i := 0; i < numRows; i++ {
		for j := i; j < n; j++ {
			ret = append(ret, s[j])
			if i > 0 && i < numRows-1 && j+t-2*i < n {
				ret = append(ret, s[j])
			}
		}
	}
	return string(ret)
}

func convert1(s string, numRows int) string {
	mat := make([][]byte, numRows)
	for i:=0;i<numRows;i++ {
		mat[i] = make([]byte,0)
	}
	t:=2*numRows-2
	for i:=0;i<len(s);i++ {
		for j:=0;j<numRows;j++{
			if i%t == j || i%t == t-j {
				mat[j] = append(mat[j], s[i])
				break
			}
		}
	}
	var ret = make([]byte,0,len(s))
	for i:=0;i<numRows;i++ {
		ret = append(ret, mat[i]...)
	}
	return string(ret)
}

func main() {
	testCases := []struct {
		s       string
		numRows int
	}{
		{
			s:  "PAYPALISHIRING",
			numRows:  3,
		},
	}
	for _, testCase := range testCases {
		fmt.Println(convert(testCase.s, testCase.numRows))
	}
}
