package main

/*type StockSpanner struct {
	history []int
	record []int
}


func Constructor() StockSpanner {
	return StockSpanner{}
}


func (this *StockSpanner) Next(price int) int {
	var ans int
	var n = len(this.record)
	var index = n - 1

	if len(this.history) == 0 {
		ans =  1
		index = 0
	} else {
		for {
			if index == -1 || price < this.history[index] {
				break
			} else if price >= this.history[index] && index == 0 {
				index = -1
				break
			}
			index = this.record[index]
		}
		ans = n - index
	}
	this.history = append(this.history, price)
	this.record = append(this.record, index)

	return ans
}
*/

type StockSpanner struct {
	stack [][2]int
	index int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	var ans int
	this.index++
	for {
		n := len(this.stack)
		if n == 0 || price < this.stack[len(this.stack)-1][0] {
			// 入栈 推出
			if n == 0 {
				ans = this.index
			} else if n != 0 {
				ans = this.index - this.stack[n-1][1]
			}
			this.stack = append(this.stack, [2]int{price, this.index})
			break
		} else if price >= this.stack[len(this.stack)-1][0] {
			this.stack = this.stack[:n-1]
		}
	}
	return ans
}
