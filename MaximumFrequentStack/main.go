package main

import "fmt"

func main() {
	fS := Constructor()
	fS.Push(4)
	fS.Push(0)
	fS.Push(9)
	fS.Push(3)
	fS.Push(4)
	fS.Push(2)
	fmt.Print(fS.Pop())
	fS.Push(6)
	fmt.Print(fS.Pop())
	fS.Push(1)
	fmt.Print(fS.Pop())
	fS.Push(1)
	fmt.Print(fS.Pop())
	fS.Push(4)
	fmt.Print(fS.Pop())
	fmt.Print(fS.Pop())
	fmt.Print(fS.Pop())
	fmt.Print(fS.Pop())
	fmt.Print(fS.Pop())
	fmt.Print(fS.Pop())
}
