package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(multply("2", "3"))
	fmt.Println(multply("123", "456"))
	fmt.Println(multply("9", "99"))
	fmt.Println(multply("498828660196", "840477629533"))
}

func multply(num1 string, num2 string) string {
	// len1, len2 := len(num1)-1, len(num2)-1
	// carry := 0
	// iterationSum := 0
	// product := 0
	// z := 0
	// for j, k := len2, 0; j >= 0 && k <= len2; j, k = j-1, k+1 {
	// 	sum := 0
	// 	z = k
	// 	for i := len1; i >= 0; i-- {
	// 		a, _ := strconv.Atoi(string(num2[j]))
	// 		b, _ := strconv.Atoi(string(num1[i]))
	// 		sum = a*b + carry
	// 		carry = sum / 10
	// 		product += int(math.Pow10(z)) * (sum % 10)
	// 		z++
	// 	}
	// 	if carry > 0 {
	// 		product += int(math.Pow10(z)) * (carry % 10)
	// 	}
	// 	iterationSum += product
	// 	product = 0
	// 	carry = 0
	// }

	// return strconv.Itoa(iterationSum)
	var sb strings.Builder

	m, n := len(num1), len(num2)
	products := make([]int, m+n)

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			a, b := int(num2[i]-'0'), int(num1[j]-'0')
			product := a*b + products[i+j+1]
			products[i+j+1] = product % 10
			products[i+j] += product / 10
		}
	}

	j := 0
	for j = range products {
		if products[j] != 0 {
			break
		}
	}

	products = products[j:]

	for _, val := range products {
		sb.WriteString(strconv.Itoa(val))
	}

	result := sb.String()

	return result
}
