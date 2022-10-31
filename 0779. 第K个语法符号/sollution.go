package main

import "fmt"

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	pre := kthGrammar(n-1, (k+1)/2)
	if (pre == 0 && k%2 == 1) || (pre == 1 && k%2 ==0) {
		return 0
	}
	return 1
}

func main () {
	fmt.Println(kthGrammar(3,3))
}