package main

type MyCircularDeque struct {
	front, rear int
	vals []int
	length int
}


func Constructor(k int) MyCircularDeque {
	return MyCircularDeque {
		vals : make([]int, k+1),
		length: k+1,
	}
}


func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.vals[this.front] = value
	this.front = (this.front-1) % this.length
	if this.front < 0 {
		this.front += this.length
	}
	return true
}


func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.rear = (this.rear+1) % this.length
	this.vals[this.rear] = value
	return true
}


func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front+1) % this.length
	return true
}


func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.rear = (this.rear-1) % this.length
	if this.rear < 0 {
		this.rear += this.length
	}
	return true
}


func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.vals[(this.front+1)%this.length]
}


func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.vals[this.rear]
}


func (this *MyCircularDeque) IsEmpty() bool {
	if this.front == this.rear {
		return true
	}
	return false
}


func (this *MyCircularDeque) IsFull() bool {
	if (this.rear + 1) % this.length == this.front {
		return true
	}
	return false
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
