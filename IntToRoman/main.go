package main

import "fmt"

func main() {
	fmt.Print(intToRoman(58))

}

func intToRoman(num int) string {
	var ans string
	var digit int

	if num < 1 && num > 3999 {
		return ""
	}
	for num > 0 {

		if num >= 1000 {
			ans += "M"
			num = num - 1000

		} else if num >= 500 {
			digit = num / 100
			if digit == 9 {
				ans += "CM"
				num = num - 900

			} else {
				ans += "D"
				num = num - 500

			}

		} else if num >= 100 {
			digit = num / 100
			if digit == 4 {
				ans += "CD"
				num = num - 400

			} else {
				ans += "C"
				num = num - 100

			}

		} else if num >= 50 {
			digit = num / 10
			if digit == 9 {
				ans += "XC"
				num = num - 90

			} else {
				ans += "L"
				num = num - 50

			}

		} else if num >= 10 {
			digit = num / 10
			if digit == 4 {
				ans += "XL"
				num = num - 40

			} else {
				ans += "X"
				num = num - 10

			}

		} else {
			ans += MapDigit(num)
			num = 0

		}

	}

	return ans

}
func MapDigit(digit int) string {
	switch digit {
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 6:
		return "VI"
	case 7:
		return "VII"
	case 8:
		return "VIII"
	case 9:
		return "IX"
	default:
		return ""
	}

}
