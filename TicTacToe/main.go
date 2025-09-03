package main

import "fmt"

func main() {
	// TicTacToe := Constructor(3)
	// fmt.Print(TicTacToe.Move(0, 0, 1))
	// fmt.Print(TicTacToe.Move(0, 2, 2))
	// fmt.Print(TicTacToe.Move(2, 2, 1))
	// fmt.Print(TicTacToe.Move(1, 1, 2))
	// fmt.Print(TicTacToe.Move(2, 0, 1))
	// fmt.Print(TicTacToe.Move(1, 0, 2))
	// fmt.Print(TicTacToe.Move(2, 1, 1))

	TicTacToe := Constructor(2)
	fmt.Print(TicTacToe.Move(0, 1, 1))
	fmt.Print(TicTacToe.Move(1, 1, 2))
	fmt.Print(TicTacToe.Move(1, 0, 1))

}
