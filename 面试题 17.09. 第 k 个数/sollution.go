package main

import "fmt"

// 暴力
func getKthMagicNumber(k int) int {
	var ans int
	var count int
	var record = make(map[int]struct{})
	for count < k {
		ans++
		var tmp = ans
		if tmp%3 == 0 {
			tmp /= 3
		} else if tmp%5 == 0 {
			tmp /= 5
		} else if tmp%7 == 0 {
			tmp /= 7
		}
		if _, ok := record[tmp]; ok || tmp == 1 {
			count++
			record[ans] = struct{}{}
		}
	}
	return ans
}

// 动态规划
func getKthMagicNumber1(k int) int {
	var dp = make([]int, k)
	dp[0] = 1
	var p3, p5, p7 = 0, 0, 0
	var min = func() int {
		var ans int
		if dp[p3]*3 < dp[p5]*5 {
			if dp[p3]*3 < dp[p7]*7 {
				ans = dp[p3] * 3
				p3++
			} else if dp[p3]*3 > dp[p7]*7 {
				ans = dp[p7] * 7
				p7++
			} else {
				ans = dp[p3] * 3
				p3++
				p7++
			}
		} else if dp[p3]*3 > dp[p5]*5 {
			if dp[p5]*5 < dp[p7]*7 {
				ans = dp[p5] * 5
				p5++
			} else if dp[p5]*5 > dp[p7]*7 {
				ans = dp[p7] * 7
				p7++
			} else {
				ans = dp[p5] * 5
				p5++
				p7++
			}
		} else {
			if dp[p5]*5 < dp[p7]*7 {
				ans = dp[p5] * 5
				p3++
				p5++
			} else if dp[p5]*5 > dp[p7]*7 {
				ans = dp[p7] * 7
				p7++
			} else {
				ans = dp[p5] * 5
				p3++
				p5++
				p7++
			}
		}
		return ans
	}
	for i := 1; i < k; i++ {
		dp[i] = min()
	}
	return dp[k-1]
}

type node struct {
	val   int
	left  *node
	right *node
}

type Heap struct {
	vals   []int
	length int
}

func (heap *Heap) Pop() int {
	if len(heap.vals) == 0 {
		return -1
	}
	var val = heap.vals[0]
	heap.vals[0] = heap.vals[heap.length-1]
	heap.vals = heap.vals[:heap.length-1]
	heap.length--
	heap.down()
	return val
}

func (heap *Heap) Push(val int) {
	// 在末端添加
	heap.vals = append(heap.vals, val)
	heap.length++
	heap.up()
	// 上浮
}

func (heap *Heap) up() {
	var cur = heap.length - 1
	for cur > 0 {
		father := (cur - 1) / 2
		if heap.vals[father] > heap.vals[cur] {
			heap.vals[father], heap.vals[cur] = heap.vals[cur], heap.vals[father]
			cur = father
			continue
		}
		break
	}
}

func (heap *Heap) down() {
	var cur = 0
	for cur < heap.length-1 {
		left := 2*cur + 1
		right := 2*cur + 2
		var min = cur
		if left < heap.length && heap.vals[min] > heap.vals[left] {
			min = left
		}
		if right < heap.length && heap.vals[min] > heap.vals[right] {
			min = right
		}
		if min == cur {
			return
		}
		heap.vals[cur], heap.vals[min] = heap.vals[min], heap.vals[cur]
		cur = min
	}
}

// 堆排序 小顶堆
// 动态规划
func getKthMagicNumber2(k int) int {
	var heap = Heap{}
	var ans int
	var record = make(map[int]struct{})
	heap.Push(1)
	var factors = []int{3, 5, 7}
	for i := 0; i < k; i++ {
		ans = heap.Pop()
		for _, v := range factors {
			if _, ok := record[v*ans]; !ok {
				record[v*ans] = struct{}{}
				heap.Push(v * ans)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(getKthMagicNumber2(11))
}
