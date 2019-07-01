package main

import "log"

// Given a string S, return the number of substrings of length K with no repeated characters.
func main() {
	log.Println(numKLenSubstrNoRepeats("havefunonleetcode", 5)) // 6
	log.Println(numKLenSubstrNoRepeats("home", 5))              // 0
	log.Println(numKLenSubstrNoRepeats("home", 4))              // 1
}

func numKLenSubstrNoRepeats(S string, K int) int {
	if K > len(S) {
		return 0
	}

	cnt, start, l := 0, 0, len(S)
	for start+K <= l {
		if isNoRepeated(S[start : start+K]) {
			cnt++
		}
		start++
	}

	return cnt
}

func isNoRepeated(s string) bool {
	t := make(map[rune]struct{})
	for _, r := range s {
		t[r] = struct{}{}
	}

	return len(t) == len(s)
}
