package main

import "log"

// Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
// Note: For the purpose of this problem, we define empty string as valid palindrome.
func main() {
	log.Println(isPalindrome("   "))                            // true
	log.Println(isPalindrome(""))                               // true
	log.Println(isPalindrome(" "))                              // true
	log.Println(isPalindrome("rotor"))                          // true
	log.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	log.Println(isPalindrome("race a car"))                     // false
	log.Println(isPalindrome("0P"))                             // false
	log.Println(isPalindrome("0P0"))                            // true
}

func isPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		if !inBounds(s[start]) {
			start++
			continue
		}

		if !inBounds(s[end]) {
			end--
			continue
		}

		if toLower(s[start]) != toLower(s[end]) {
			return false
		}

		start++
		end--
	}

	return true
}

func toLower(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return char + ('a' - 'A')
	}

	return char
}

func inBounds(char byte) bool {
	return (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')
}
