package main

import "log"

// Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.
func main() {
	log.Println(validPalindrome("aba"))   // true
	log.Println(validPalindrome("abca"))  // true
	log.Println(validPalindrome("abcaa")) // false
	log.Println(validPalindrome("allea")) // true
	log.Println(validPalindrome("abc"))   // true
}

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l <= r {
		if s[l] == s[r] {
			l++
			r--
			continue
		}

		return isPalindrome(s, l+1, r) || isPalindrome(s, l, r-1)
	}

	return true
}

func isPalindrome(s string, l, r int) bool {
	for l <= r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			return false
		}
	}

	return true
}
