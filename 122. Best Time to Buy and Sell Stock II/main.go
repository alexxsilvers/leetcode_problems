package main

import "log"

// Say you have an array for which the ith element is the price of a given stock on day i.
//
// Design an algorithm to find the maximum profit. You may complete as many transactions as
// you like (i.e., buy one and sell one share of the stock multiple times).
//
// Note: You may not engage in multiple transactions at the same time (i.e., you must sell
// the stock before you buy again).
func main() {
	log.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 7
	log.Println(maxProfit([]int{}))                 // 0
	log.Println(maxProfit([]int{1}))                // 0
}

// Use peak-valley strategy
// Runtime - O(n)
// Space - O(1)
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	total := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			total += prices[i+1] - prices[i]
		}
	}

	return total
}
