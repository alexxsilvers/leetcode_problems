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
	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}
	}
	return max
}
