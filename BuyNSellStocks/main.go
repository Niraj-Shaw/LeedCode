package main

import "fmt"

func main() {
	prices := []int{3}
	fmt.Print(maxProfit(prices))
}

func maxProfit(prices []int) int {
	sellingPrice, BuyingPrice := 0, prices[0]
	max := 0
	for _, val := range prices {
		if val <= BuyingPrice {
			BuyingPrice = val
			sellingPrice = 0
		}
		if val >= sellingPrice {
			sellingPrice = val
		}
		if sellingPrice-BuyingPrice > max {
			max = sellingPrice - BuyingPrice
		}
	}

	return max
}
