package main

import (
	"log"
	"math"
)

// Say you have an array for which the ith element is the price of a given stock on day i.
//
// If you were only permitted to complete at most one transaction
// (i.e., buy one and sell one share of the stock), design an algorithm to
// find the maximum profit.
//
// Note that you cannot sell a stock before you buy one.
func main() {
	log.Println(maxProfit([]int{2, 6, 7, 10, 1, 8})) // 8.
	// Explanation: buy at day 0 with price = 2 and sell at day 3 with price = 10.
	// Profit = 8

	log.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 5.
	// Explanation: buy at day 4 with price = 6 and sell at day 1 with price = 1.
	// Profit = 5

	log.Println(maxProfit([]int{7, 6, 5})) // 0.

	log.Println(maxProfit([]int{7, 6, 4, 3, 1})) // 0.
}

// Runtime O(n2)
// Space O(1)
//func maxProfit(prices []int) int {
//	max := 0
//	for i := 0; i < len(prices); i++ {
//		for j := len(prices) - 1; j > i; j-- {
//			profit := prices[j] - prices[i]
//			if profit > max {
//				max = profit
//			}
//		}
//	}
//	return max
//}

// Runtime O(n)
// Space O(1)
func maxProfit(prices []int) int {
	min := math.MaxInt64
	profit := 0
	for _, p := range prices {
		if p < min {
			min = p
		}

		if p-min > profit {
			profit = p - min
		}
	}
	return profit
}
