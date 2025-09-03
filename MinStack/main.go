package main

import "fmt"

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	print(minStack.stack)
	minStack.Push(0)
	print(minStack.stack)
	minStack.Push(-3)
	print(minStack.stack)
	fmt.Println(minStack.getMin())
	minStack.Pop()
	print(minStack.stack)
	fmt.Println(minStack.Top())
	fmt.Println(minStack.getMin())
}

func print(stack [][]int) {
	for _, val := range stack {
		fmt.Print(val[0])
		fmt.Print(" ")
	}
	fmt.Println()
}
