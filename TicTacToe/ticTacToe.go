package main

type TicTacToe struct {
	size        int
	rows        [][2]int
	cols        [][2]int
	diagonal    []int
	antidiagnol []int
	board       [][]string
}

func Constructor(n int) TicTacToe {
	bd := make([][]string, n)
	for i, _ := range bd {
		bd[i] = make([]string, n)
	}
	return TicTacToe{
		size:        n,
		board:       bd,
		rows:        getIniatializedArray(n),
		cols:        getIniatializedArray(n),
		diagonal:    make([]int, n),
		antidiagnol: make([]int, n),
	}
}

func (this *TicTacToe) Move(row int, col int, player int) int {
	v := ""
	if player == 1 {
		v = "x"
	} else {
		v = "o"
	}
	if this.board[row][col] != "" {
		return 0
	}
	this.board[row][col] = v

	this.rows[row][player-1]++
	if this.rows[row][player-1] == this.size {
		return player
	}

	this.cols[col][player-1]++
	if this.cols[col][player-1] == this.size {
		return player
	}

	if row == col {
		this.diagonal[player-1]++
	}
	if this.diagonal[player-1] == this.size {
		return player
	}

	if row == this.size-col-1 {
		this.antidiagnol[player-1]++
	}
	if this.antidiagnol[player-1] == this.size {
		return player
	}

	return 0

}

func getIniatializedArray(size int) [][2]int {
	array := [][2]int{{}}
	for i := 0; i < size; i++ {
		array = append(array, [2]int{})
	}
	return array
}
