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
	log.Println(maxProfitPeaksAndValleys([]int{2, 6, 7, 10, 1, 8})) // 8
	// Explanation: buy at day 0 with price = 2 and sell at day 3 with price = 10.
	// Profit = 8

	log.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 5.
	log.Println(maxProfitPeaksAndValleys([]int{7, 1, 5, 3, 6, 4})) // 5.
	// Explanation: buy at day 4 with price = 6 and sell at day 1 with price = 1.
	// Profit = 5

	log.Println(maxProfit([]int{7, 6, 5})) // 0.
	log.Println(maxProfitPeaksAndValleys([]int{7, 6, 5})) // 0.

	log.Println(maxProfit([]int{7, 6, 4, 3, 1})) // 0.
	log.Println(maxProfitPeaksAndValleys([]int{7, 6, 4, 3, 1})) // 0.
}

func maxProfit(prices []int) int {
	return maxProfitBF(prices)
}

// Runtime O(n), Space O(1)
func maxProfitBF(prices []int) int {
	bestProfit := 0

	for i := 0; i < len(prices); i++ {
		for j := i+1; j < len(prices); j++ {
			if prices[j] - prices[i] > bestProfit {
				bestProfit = prices[j] - prices[i]
			}
		}
	}

	return bestProfit
}

func maxProfitPeaksAndValleys(prices []int) int {
	maxProfit, minPrice := 0, math.MaxInt64

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		if prices[i] - minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}