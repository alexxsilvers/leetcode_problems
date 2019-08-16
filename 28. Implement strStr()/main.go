package main

import "log"

// Implement strStr().
//
// Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
func main() {
	log.Println(strStr("hello", "ll"))
	log.Println(strStr("aaaaaaa", "baa"))
	log.Println(strStr("asdfg", "g"))
}

func strStr(haystack string, needle string) int {
	hL, nL := len(haystack), len(needle)

	if nL == 0 {
		return 0
	}

	if nL > hL {
		return -1
	}

	for i := 0; i < hL; i++ {
		if haystack[i] == needle[0] {
			if i+nL > hL {
				return -1
			}

			finded := true
			for j := 0; j < nL; j++ {
				if haystack[i+j] != needle[j] {
					finded = false
					break
				}
			}

			if finded {
				return i
			}
		}
	}

	return -1
}
