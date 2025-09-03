package main

import "math"

type MinStack struct {
	stack [][]int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([][]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	length := len(this.stack)
	if length == 0 {
		this.stack = append(this.stack, []int{val, val})
	} else {
		temp := this.stack[length-1]
		min := temp[1]
		if val < min {
			min = val
		}
		this.stack = append(this.stack, []int{val, min})
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1][0]
	}
	return math.MinInt32
}

func (this *MinStack) getMin() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1][1]
	}
	return math.MinInt32
}
