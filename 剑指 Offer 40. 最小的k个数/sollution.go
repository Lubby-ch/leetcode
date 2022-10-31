package main

import "fmt"

func getLeastNumbers(arr []int, k int) []int {
	heap := NewHeap(k)
	for _, v := range arr {
		heap.Push(v)
	}
	return heap.data
}

type Heap struct {
	data   []int
	length int
	size   int
}

func NewHeap(length int) *Heap {
	return &Heap{
		data:   make([]int, 0, length),
		length: length,
	}
}

func (heap *Heap) Replace(val int) {
	if heap.data[0] <= val {
		return
	}
	heap.data[0] = val
	heap.Down(0)
}

func (heap *Heap) Push(val int) {
	if heap.size > heap.length {
		return
	}

	if heap.size == heap.length {
		heap.Replace(val)
	} else {
		heap.data = append(heap.data, val)
		heap.size++
		heap.Up(heap.size - 1)
	}
}

func (heap *Heap) Pop() int {
	if heap.size == 0 {
		return -1
	}
	var ret = heap.data[0]
	heap.data[0] = heap.data[heap.size-1]
	heap.size--
	heap.Down(0)

	return ret
}

func (heap *Heap) Down(index int) { // 下沉
	tmp := index
	left := 2*index + 1
	right := 2*index + 2
	if left < heap.size && heap.data[tmp] < heap.data[left] {
		tmp = left
	}
	if right < heap.size && heap.data[tmp] < heap.data[right] {
		tmp = right
	}
	if tmp == index {
		return
	}
	heap.data[index], heap.data[tmp] = heap.data[tmp], heap.data[index]
	heap.Down(tmp)
}

func (heap *Heap) Up(index int) { // 上浮
	father := (index - 1) / 2
	if heap.data[index] > heap.data[father] {
		heap.data[index], heap.data[father] = heap.data[father], heap.data[index]
		heap.Up(father)
	}
}

func main() {
	fmt.Println(getLeastNumbers([]int{5, 3, 2, 1, 6}, 3))
}
