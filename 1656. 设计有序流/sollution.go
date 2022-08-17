package main

import "fmt"

type OrderedStream struct {
	vals   []string
	length int
	ptr    int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		vals:   make([]string, n),
		length: n,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	if idKey > this.length {
		return nil
	}
	this.vals[idKey-this.ptr-1] = value
	var ans []string
	for 0 < len(this.vals) {
		if len(this.vals[0]) == 0 {
			break
		}
		ans = append(ans, this.vals[0])
		this.vals = this.vals[1:]
		this.ptr++
	}
	return ans
}

func main() {
	stream := Constructor(5)
	fmt.Println(stream.Insert(3, "c"))
	fmt.Println(stream.Insert(1, "a"))
	fmt.Println(stream.Insert(2, "b"))
	fmt.Println(stream.Insert(5, "e"))
	fmt.Println(stream.Insert(4, "d"))
}
