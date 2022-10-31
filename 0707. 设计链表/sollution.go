package main

type Node struct {
	val  int
	next *Node
}
type MyLinkedList struct {
	head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head: &Node{},
	}
}

func (this *MyLinkedList) Get(index int) int {
	node := this.head
	for i := 0; i <= index; i++ {
		if node.next == nil {
			return -1
		}
		node = node.next
	}
	return node.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.head.next = &Node{
		val: val,
		next: this.head.next,
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := this.head
	for node.next != nil {
		node = node.next
	}
	node.next = &Node{
		val: val,
		next: node.next,
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	node := this.head
	for i := 0; i <= index-1; i++ {
		if node.next == nil {
			break
		}
		node = node.next
	}
	node.next = &Node{
		val:  val,
		next: node.next,
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	node := this.head
	for i := 0; i <= index-1; i++ {
		if node.next == nil {
			break
		}
		node = node.next
	}
	if node.next != nil {
		node.next = node.next.next
	}
}

func main (){
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1,2) //链表变为1-> 2-> 3
	linkedList.Get(1)        //返回2
	linkedList.DeleteAtIndex(1) //现在链表是1-> 3
	linkedList.Get(1)
}