package main

import "fmt"

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 3))
}

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	result := ""

	matrix := make([][]byte, numRows)

	colCount := 0

	res := len(s)
	for res > 0 {
		res = res - numRows
		colCount++

		for k := 1; res > 0 && k <= numRows-2; k, res = k+1, res-1 {
			colCount++
		}
	}

	for i := range numRows {
		matrix[i] = make([]byte, colCount)
	}

	l := 0
	j := 0
	i := 0

	for l < len(s) {
		for i < numRows && l < len(s) {
			matrix[i][j] = s[l]
			l++
			i++
		}

		for i, j = i-2, j+1; l < len(s) && i > 0 && j < colCount; i, j, l = i-1, j+1, l+1 {
			matrix[i][j] = s[l]

		}

	}

	for _, row := range matrix {
		for _, v := range row {
			if v != 0 {
				result += string(v)
			}
		}
	}

	return result

}
