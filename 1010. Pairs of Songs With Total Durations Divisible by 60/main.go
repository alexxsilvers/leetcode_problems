package main

import "log"

// In a list of songs, the i-th song has a duration of time[i] seconds.
//
// Return the number of pairs of songs for which their total duration in seconds
// is divisible by 60.  Formally, we want the number of indices i < j with (time[i] + time[j]) % 60 == 0.
func main() {
	log.Println(numPairsDivisibleBy60([]int{30, 20, 150, 100, 40})) // 3
	log.Println(numPairsDivisibleBy60([]int{60, 60, 60}))           //3
	log.Println(numPairsDivisibleBy60([]int{20, 40}))               //1
}

// Runtime - O(n)
// Space O(1)
func numPairsDivisibleBy60(time []int) int {
	div := make([]int, 60)

	for i := 0; i < len(time); i++ {
		div[time[i]%60]++
	}

	cnt := 0
	for i := 1; i < 30; i++ {
		cnt += div[i] * div[60-i]
	}

	cnt += div[0] * (div[0] - 1) / 2
	cnt += div[30] * (div[30] - 1) / 2

	return cnt
}
