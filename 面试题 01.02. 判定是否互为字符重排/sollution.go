package main

import (
	"fmt"
	"sort"
	"time"
)

func CheckPermutation(s1 string, s2 string) bool {
	b1 := []byte(s1)
	b2 := []byte(s2)
	sort.Slice(b1, func(i, j int) bool {
		return b1[i] > b1[j]
	})
	sort.Slice(b2, func(i, j int) bool {
		return b2[i] > b2[j]
	})
	return string(b1) == string(b2)
}

func CheckPermutation2(s1 string, s2 string) bool {
	var c1 = make(map[rune]int)
	for _, ch := range s1 {
		c1[ch]++
	}
	for _, ch := range s2 {
		if v, ok := c1[ch]; !ok || v == 0 {
			return false
		}
		c1[ch]--
	}
	return true
}

type user struct {
	name string
	age int8
}

var u = user{name: "Ankur", age: 25}
var g = &u

func modifyUser(pu *user) {
	fmt.Printf("modifyUser Received Vaule %p \n", pu)
	pu.name = "Anand"
}

func printUser(u <-chan *user) {
	time.Sleep(2 * time.Second)
	fmt.Printf("printUser goRoutine called %p \n", <-u)
}

func main() {
	c := make(chan *user, 5)
	c <- g
	fmt.Printf("%p \n", g)
	// modify g
	g = &user{name: "Ankur Anand", age: 100}
	fmt.Printf("%p \n", g)
	go printUser(c)
	go modifyUser(g)
	time.Sleep(5 * time.Second)
	fmt.Println(g)
}