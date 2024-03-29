package main

import (
	"fmt"
	"reflect"
)

type MyCircularQueue struct {
	head, tail, length int
	elems              []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		elems:  make([]int, k+1),
		length: k + 1,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.tail = (this.tail + 1) % this.length
	this.elems[this.tail] = value
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % this.length
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.elems[(this.head+1)%this.length]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.elems[this.tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.tail == this.head
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.tail+1)%this.length == this.head
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	var a int64 = 100
	// reflectSetValue1(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&a)
	fmt.Println(a)
}