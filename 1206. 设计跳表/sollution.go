package main

import (
	"math"
	"math/rand"
	"time"
)

type Node struct {
	val  int
	Next *Node
	Down *Node
}

type Skiplist struct {
	MaxLevel   int
	LevelHeads []*Node
}

func Constructor() Skiplist {
	return Skiplist{
		MaxLevel:   0,
		LevelHeads: make([]*Node, 1),
	}
}

func (this *Skiplist) Search(target int) bool {
	head := this.LevelHeads[this.MaxLevel]
	for {
		if head.val == target {
			return true
		}
		if head.Next != nil && head.Next.val <= target {
			head = head.Next
			continue
		}
		if head.Down != nil {
			head = head.Down
			continue
		}
		return false
	}
}

func (this *Skiplist) Add(num int) {
	level := GetLevel(10)
	if level > this.MaxLevel {
		this.LevelHeads = append(this.LevelHeads, make([]*Node, level-this.MaxLevel)...)
		this.MaxLevel = level
	}
	var upNode *Node
	for i := level; i >= 0; i-- {
		head := this.LevelHeads[i]
		if head == nil {
			head = &Node{
				val: math.MinInt,
			}
			this.LevelHeads[i] = head
		}
		if i+1 <= this.MaxLevel && this.LevelHeads[i+1].Down == nil {
			this.LevelHeads[i+1].Down = head
		}
		for {
			/* 官方允许跳表中的值重复，如果不需要值重复则可以取消下面注释
			if head.val == num {
				return
			}
			 */
			if head.Next == nil || head.Next.val > num {
				// 新增一个节点
				node := &Node{
					val:  num,
					Next: head.Next,
				}
				head.Next = node
				if upNode != nil {
					upNode.Down = node
				}
				upNode = node
				break
			}
			head = head.Next
		}
	}
	return
}

func (this *Skiplist) Erase(num int) bool {
	head := this.LevelHeads[this.MaxLevel]
	var isFind bool
	for {
		if head.Next != nil && head.Next.val == num {
			isFind = true
			// 删除逻辑
			head.Next = head.Next.Next
			if head.Down == nil {
				return isFind
			}
			head = head.Down
			continue
		}
		if head.Next != nil && head.Next.val < num {
			head = head.Next
			continue
		}
		if head.Down != nil {
			head = head.Down
			continue
		}
		return isFind
	}
}

func GetLevel(maxLevel int) int {
	var level int
	rand.Seed(time.Now().UnixMilli())
	for {
		if rand.Float64() < 0.5 && level < maxLevel {
			level++
			continue
		}
		break
	}
	return level
}