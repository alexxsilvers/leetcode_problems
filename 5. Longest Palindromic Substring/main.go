package main

import "log"

// Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
func main() {
	//log.Println(longestPalindrome("babad")) // "bab" or "aba"
	log.Println(longestPalindrome("ababa")) // "ababa"
	log.Println(longestPalindrome(""))      // ""
	log.Println(longestPalindrome("1"))     // "1"
	log.Println(longestPalindrome("22"))    // "22"
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	l := len(s)

	min, max := 0, 0
	for i := 0; i < l; i++ {
		low := i
		high := i
		// check odd
		for low >= 0 && high < l && s[low] == s[high] {
			low--
			high++
		}
		if high-low-1 > max-min {
			min = low + 1
			max = high
		}

		low = i
		high = i + 1
		// check even
		for low >= 0 && high < l && s[low] == s[high] {
			low--
			high++
		}
		if high-low-1 > max-min {
			min = low + 1
			max = high
		}
	}

	return s[min:max]
}
