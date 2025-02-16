package main

type MyQueue struct {
	list []int
}

func main() {
	/**
	 * Your MyQueue object will be instantiated and called as such:
	 * obj := Constructor();
	 * obj.Push(x);
	 * param_2 := obj.Pop();
	 * param_3 := obj.Peek();
	 * param_4 := obj.Empty();
	 */

}

func Constructor() MyQueue {
	var result MyQueue
	result.list = []int{}
	return result
}

func (this *MyQueue) Push(x int) {
	this.list = append(this.list, x)

}

func (this *MyQueue) Pop() int {
	result := this.list[0]
	this.list = this.list[1:]
	return result

}

func (this *MyQueue) Peek() int {
	return this.list[0]
}

func (this *MyQueue) Empty() bool {
	if len(this.list) == 0 {
		return true
	}
	return false

}
