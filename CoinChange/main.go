package main

import (
	"fmt"
)

func main() {
	fmt.Print(coinChange([]int{2}, 3))
	//fmt.Print(coinChange([]int{1}, 0))
}

func coinChange(coins []int, amount int) int {
	//BruteForce
	// have to find 1 + min(f(n-1), ...) where n is the amount
	// add results of f(n) in map if not present
	//for
	return dfs(coins, amount)

}

func dfs(coins []int, amount int) int {
	// minCount := 0
	// if amount == 0 {
	// 	return 0
	// }
	// if amount < 0 {
	// 	return -1
	// }
	// result, minCount := 0, math.MaxInt
	// for i := 0; i < len(coins) && amount > 0; i++ {
	// 	if val, ok := dp[amount-coins[i]]; ok {
	// 		result = val
	// 	} else {
	// 		result = dfs(coins, amount-coins[i])
	// 	}
	// 	if result != -1 && (result+1 < minCount) {
	// 		minCount = result + 1
	// 	}
	// }
	// if minCount == math.MaxInt {
	// 	dp[amount] = -1
	// 	return -1
	// } else {
	// 	dp[amount] = minCount
	// 	return minCount
	// }
	// dp bottom up approach
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	for i := 1; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
