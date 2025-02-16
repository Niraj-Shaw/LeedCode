package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Print(myAtoi("42"))

}

func myAtoi(s string) int {

	index := 0
	sign := 1
	res := 0
	MAXINT := math.MaxInt32
	MININT := math.MinInt32
	overLimit := false
	if len(s) == 0{
		return 0
	}
	//check for whitespaces in begining
	s = strings.Trim(s, " ")

	//store the sign

	if s[0] == '-' {
		sign = -1
		index++
	}else if (s[0] == '+'){
		sign = 1
		index++
	}
	//proceed and calculate using prev result * 10 + current number
	// stop when finding non integer oe end of the string
	//return MAXINT or MININT in case of overflow or underflow
	for index < len(s) && isDigit(s[index]) {
		curr := int(s[index] - '0')
		if res > MAXINT/10 || (res == MAXINT/10 && s[index] > '7') {
			overLimit = true
			break
		}
		res = 10*res + curr
		index++

	}
	if overLimit {
		if sign == -1 {
			return MININT
		} else {
			return MAXINT
		}
	}
	if sign == -1{
		return -(res)
	}
	return res

}

func isDigit(digit byte) bool {
	return digit >= '0' && digit <= '9'

}
