package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Print(reverse(120))
}

func reverse(x int) int {
	str := strconv.Itoa(x)
	isNegative := false
	if x < 0 {
		str = str[1:]
		isNegative = true
	}
	runeArray := []rune(str)
	for i, j := 0, len(runeArray)-1; i < j; i, j = i+1, j-1 {
		runeArray[i], runeArray[j] = runeArray[j], runeArray[i]
	}
	str = string(runeArray)
	val, _ := strconv.Atoi(str)
	if val < math.MinInt32 || val > math.MaxInt32 {
		return 0
	}
	if isNegative {
		return -val
	}
	return val
}
