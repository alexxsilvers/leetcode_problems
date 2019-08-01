package main

import "log"

// We have two special characters. The first character can be represented by one bit 0.
// The second character can be represented by two bits (10 or 11).
//
// Now given a string represented by several bits. Return whether the last character must
// be a one-bit character or not. The given string will always end with a zero.
func main() {
	log.Println(isOneBitCharacter([]int{1, 1, 1, 0})) // false
	log.Println(isOneBitCharacter([]int{1, 1, 0}))    // true
}

// Runtime - O(n)
// Space - O(1)

// if there is only one symbol in array the answer is always true (as last element is 0)
// if there are two 0s at the end again the answer is true no matter what the rest symbols are( ...1100, ...1000,)
// if there is 1 right before the last element(...10), the outcome depends on the count of sequential 1, i.e.
//   a) if there is odd amount of 1(10, ...01110, etc) the answer is false as there is a single 1 without pair
//   b) if it's even (110, ...011110, etc) the answer is true, as 0 at the end doesn't have anything to pair with
func isOneBitCharacter(bits []int) bool {
	ones := 0
	for i := len(bits) - 2; i >= 0 && bits[i] != 0; i-- {
		ones++
	}

	if ones&1 == 0 {
		return true
	}

	return false
}
